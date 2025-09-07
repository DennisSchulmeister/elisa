# Elisa: AI Learning Assistant
# © 2025 Dennis Schulmeister-Zimolong <dennis@wpvs.de>
#
# This program is free software: you can redistribute it and/or modify
# it under the terms of the GNU Affero General Public License as
# published by the Free Software Foundation, either version 3 of the
# License, or (at your option) any later version.

from __future__ import annotations
from typing     import TYPE_CHECKING, override
from pydantic   import BaseModel, Field
from json       import dumps as json_dumps

from .._agent   import AgentBase, CreateActivity, Stateless
from ..shared   import default_role_description, default_summary_message, short_role_description

if TYPE_CHECKING:
    from ...auth.user import User
    from .._agent     import ProcessChatMessageResult
    from ..models     import AssistantChatMessage, UserChatMessage, SpeakMessageContent

class Choice(BaseModel):
    """
    A single menu choice.
    """
    agent:       str
    activity:    str
    description: str

class ChoiceActivity(BaseModel):
    """
    Activity where the user can choose other activities from a menu.
    """
    choices: list[Choice] = Field(
        default     = [],        
        description = "Activities from which the user may choose",
    )

class ChoiceActivityUpdate(ChoiceActivity):
    """
    Newly created or updated choices.
    """
    title: str = Field(
        description = "Menu title",
    )

    is_new: bool = Field(
        default     = True,
        description = "Create a new list of activities if there is none yet or the user wants to keep the previous list",
    )

class ChoiceAgent(AgentBase[Stateless]):
    """
    Suggest and offer a choice of interactive activities.
    """
    code = "choice-agent"

    def __init__(self, **kwargs):
        super().__init__(**kwargs)
    
    activities = {
        "choice": "A menu with interactive learning activities to choose from",
    }

    @override
    async def process_chat_message(self, msg: UserChatMessage, user: User) -> ProcessChatMessageResult:
        """
        Respond to the given user chat message. To increase the UX, two distinct LLM calls
        are made. First a synchronous call to update the choices then an asynchronous call
        to stream a reply message.
        """
        # Create a list of all available activities
        available_choices = []

        for agent_code, agent in self._assistant.agents.items():
            for activity_code, description in agent.activities.items():
                available_choices.append({
                    "agent":       agent_code,
                    "activity":    activity_code,
                    "description": description,
                })
            
        available_choices_json = json_dumps(available_choices)

        # Create or update activity data (choice activity)
        if self._assistant.current_activity \
        and self._assistant.match_current_activity(agent = self.code, activity = "choice"):
            current_activity = self._assistant.current_activity

            create_choices_message = """
                Task: Update the following menu of learning activities that you can perform with
                the user to support them in their learning:

                ```json
                {{ current_choices }}
                ```

                Available: Only choose activities from the following list and never invent others.

                Title: The current menu title is {{ current_title }}

                ```json
                {{ available_choices }}
                ```

                Language: Please respond in <language_code>{{ language }}</language_code>.
            """

            made_changes_message = """
                What happened: You have updated the menu of learning activities that you can perform
                with the user to support them in their learning. The previous menu was (given here as
                JSON, though the user cannot see the raw JSON data):

                ```json
                {{ current_choices }}
                ```

                The menu title was {{ current_title }}.

                {% if is_new %}
                You created a new menu with the following entries (again given as raw JSON that the
                user cannot see):
                {% else %}
                You made the following change (again given as raw JSON that the user cannot see):
                {% endif %}

                ```json
                {{ updated_choices }}
                ```

                For this you were able to choose from the following list (again as JSON raw data):

                ```json
                {{ available }}
                ```

                The new menu title is {{ updated_title }}.
            """

            current_title_json   = json_dumps(self._assistant.current_activity.title)
            current_choices_json = json_dumps(self._assistant.current_activity.data["choices"])
        else:
            current_activity = None

            create_choices_message = """
                Task: Create a choice of learning activities that you can perform with the user
                to support them in their learning. Only use activities from the following list
                and never invent others.

                ```json
                {{ available_choices }}
                ```

                Language: Please respond in <language_code>{{ language }}</language_code>.
            """

            made_changes_message = """
                What happened: You have created the following menu of learning activities that you can
                perform with the user to support them in their learning (given in JSON, though the user
                cannot see the raw JSON data):

                ```json
                {{ updated_choices }}
                ```

                For this you were able to choose from the following list (again as JSON raw data that
                the user cannot see):

                ```json
                {{ available_choices }}
                ```

                The new menu title is {{ updated_title }}.
            """

            current_title_json   = ""
            current_choices_json = ""
            
        update = await self._assistant.client.chat.completions.create(
            messages = [
                {
                    "role": "system", 
                    "content": short_role_description,
                }, {
                    "role": "system", 
                    "content": default_summary_message,
                }, {
                    "role": "system", 
                    "content": create_choices_message,
                }, {
                    "role": "system",
                    "content": """
                        Task: Please answer the user message and explain briefly what you did.

                        Language: Please respond in <language_code>{{ language }}</language_code>.
                    """
                },{
                    "role": "user",
                    "content": msg.content.speak,
                }
            ],
            context = {
                "memory":            self._assistant.state.memory,
                "language":          self._assistant.language,
                "available_choices": available_choices_json,
                "current_choices":   current_choices_json,
                "current_title":     current_title_json,
            },
            response_model = ChoiceActivityUpdate,
        )

        # Create or update activity
        if current_activity and not update.is_new:
            if current_activity.title != update.title:
                await self.update_activity("data.title", update.title, user)

            await self.update_activity("data.choices", update.choices, user)
        else:
            await self.create_activity(
                user   = user,
                create = CreateActivity(
                    activity = "choice",
                    title    = update.title,
                    data     = {
                        "choices": update.choices
                    },
                ),
            )

        # Stream response message to the user
        await self._assistant.stream_assistant_chat_message(
            user_message      = msg,
            assistant_message = AssistantChatMessage(),
            partials          = self._assistant.client.chat.completions.create_partial(
                messages = [
                    {
                        "role": "system",
                        "content": default_role_description,
                    }, {
                        "role": "system", 
                        "content": default_summary_message,
                    }, {
                        "role": "system", 
                        "content": made_changes_message,
                    }, {
                        "role": "user",
                        "content": msg.content.speak,
                    }
                ],
                context = {
                    "memory":            self._assistant.state.memory,
                    "language":          self._assistant.language,
                    "available_choices": available_choices_json,
                    "current_choices":   current_choices_json,
                    "current_title":     current_title_json,
                    "updated_choices":   json_dumps(update.choices),
                    "updated_title":     json_dumps(update.title),
                    "is_new":            update.is_new,
                },
                response_model = SpeakMessageContent,
            ),
        )

        return True