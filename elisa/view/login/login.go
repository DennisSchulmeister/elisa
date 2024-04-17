// Elisa: AI Programming Tutor for Students – CLI
// © 2024 Dennis Schulmeister-Zimolong <dennis@wpvs.de>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.

package login

import (
	"github.com/DennisSchulmeister/elisa/elisa/ui"
	"github.com/DennisSchulmeister/elisa/elisa/view/main_menu"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

// Initial screen after program start-up. This reads the exercise description from
// the file elisa.yaml and presents an input field where the course password can
// be entered to login on the server.
type LoginModel struct {
	welcomeMessage string
	ServerAddress  textinput.Model
	CoursePassword textinput.Model
}

// Create new instance of the model to read the exercise and login on the server
func NewLoginModel() LoginModel {
	model := LoginModel{
		welcomeMessage: "Willkommen bei Elisa",
		ServerAddress: textinput.Model{
			Placeholder: "Server-Adresse",
			CharLimit:   255,
			PromptStyle: ui.LabelStyle,
			TextStyle:   ui.SelectedStyle,
		},
		CoursePassword: textinput.Model{
			Placeholder:   "Kurs-Passwort",
			CharLimit:     255,
			EchoMode:      textinput.EchoPassword,
			EchoCharacter: '•',
			PromptStyle:   ui.LabelStyle,
			TextStyle:     ui.SelectedStyle,
		},
	}

	model.ServerAddress.Cursor.Style = ui.CursorStyle
	model.CoursePassword.Cursor.Style = ui.CursorStyle

	return model
}

// Run initial command
func (login LoginModel) Init() tea.Cmd {
	return tea.Sequence(tea.EnterAltScreen, textinput.Blink)
}

// Process event messages
func (login LoginModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	model, command := ui.Update(msg)
	if model != nil || command != nil {
		return model, command
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		// Handle custom keys
		keypress := msg.String()

		if keypress == "q" || keypress == "esc" {
			// Exit application
			return login, tea.Quit
		} else if keypress == "enter" {
			// Go to main menu
			mainMenu := main_menu.GetMainMenu()
			return mainMenu, nil
		}
	}

	return login, nil
}

// Render view
func (login LoginModel) View() string {
	result := ui.ScreenTitleStyle.Render(login.welcomeMessage) + "\n\n"
	result += "TEST TEST TEST"
	return result
}
