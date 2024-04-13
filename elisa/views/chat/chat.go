// Elisa: AI Programming Tutor for Students – CLI
// © 2024 Dennis Schulmeister-Zimolong <dennis@wpvs.de>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.

package chat

import (
	tea "github.com/charmbracelet/bubbletea"
)

// Chat view where the user can ask questions about the exercise or the own source code
// similar to the web UI of ChatGPT and similar AI chat bots.
type ChatModel struct {}

// Create new instance of the model to start a new chat and display it on screen
func New() ChatModel {
	return ChatModel{}
}

// Run initial command
func (m ChatModel) Init() tea.Cmd {
	return nil
}

// Process event messages
func (m ChatModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

// Render view
func (m ChatModel) View() string {
	return "Chat"
}