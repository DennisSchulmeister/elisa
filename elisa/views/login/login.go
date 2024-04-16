// Elisa: AI Programming Tutor for Students – CLI
// © 2024 Dennis Schulmeister-Zimolong <dennis@wpvs.de>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.

package login

import (
	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

// TODO: Welcome message
// TODO: Input field for server address (pre-filled)
// TODO: Input field for course password
// TODO: Login button
// TODO: Exit button

// TODO: Read elisa.yml file at startup, show warning if not found
// TODO: Send qRPC request to server to check credentials

// Initial screen after program start-up. This reads the exercise description from
// the file `elisa.yaml` and presents an input field where the course password can
// be entered to login on the server.
type LoginModel struct {
	welcomeMessage string
	ServerAddress  textinput.Model
	CoursePassword textinput.Model
	cursorMode     cursor.Mode
}

// Create new instance of the model to read the exercise and login on the server
func NewLoginModel() LoginModel {
	return LoginModel{
		welcomeMessage: "Willkommen bei Elisa, der KI-Programmiertutorin für Studierende",
		ServerAddress: textinput.Model{
			Placeholder: "Server-Adresse",
			CharLimit:   255,
			// Cursor.Style:
			// PromptStype:
			// TextStyle:
		},
		CoursePassword: textinput.Model{
			Placeholder:   "Kurs-Passwort",
			CharLimit:     255,
			EchoMode:      textinput.EchoPassword,
			EchoCharacter: '•',
			// Cursor.Style:
			// PromptStype:
			// TextStyle:
		},
	}
}

// Run initial command
func (m LoginModel) Init() tea.Cmd {
	return textinput.Blink
}

// Process event messages
func (m LoginModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

// Render view
func (m LoginModel) View() string {
	return "Login"
}
