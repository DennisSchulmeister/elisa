// Elisa: AI Programming Tutor for Students – CLI
// © 2024 Dennis Schulmeister-Zimolong <dennis@wpvs.de>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.

package chat

import (
	"github.com/DennisSchulmeister/elisa/elisa/ui"
	tea "github.com/charmbracelet/bubbletea"
)

// Chat view where the user can ask questions about the exercise or the own source code
// similar to the web UI of ChatGPT and similar AI chat bots.
type ChatModel struct {
	started	bool		// Chat has already been started
}

var model ChatModel

// Create new instance of the model to start a new chat and display it on screen.
// Use this each time the user wants to start a new chat.
func NewChatModel() ChatModel {
	model = ChatModel{
		started: true,
	}

	return model
}

// Get previously used model to return to an already started chat. Use this, each time
// the user wants to return to the chat without restarting. Starts a new chat, if no
// chat was already existing.
func Get() ChatModel {
	if !model.started {
		return NewChatModel()
	}

	return model
}

// Run initial command
func (chat ChatModel) Init() tea.Cmd {
	return nil
}

// Process event messages
func (chat ChatModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	model, command := ui.Update(msg)
	if model != nil || command != nil {return model, command}
	
	return chat, nil
}

// Render view
func (chat ChatModel) View() string {
	return "Chat"
}