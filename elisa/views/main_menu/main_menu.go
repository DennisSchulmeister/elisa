// Elisa: AI Programming Tutor for Students – CLI
// © 2024 Dennis Schulmeister-Zimolong <dennis@wpvs.de>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.

package main_menu

import (
	tea "github.com/charmbracelet/bubbletea"
)

// Main menu of the application that will be presented on when the program starts.
// Here the use can choose to see the current exercise description, to ask questions
// on the exercise, to ask questions on the own source code or to have the AI LLM
// rate the solution. Of course the application can be quit, too.
type MainMenuModel struct {
	quitting bool
}

// Create new instance of the model to display the main menu on screen
func New() MainMenuModel {
	return MainMenuModel{}
}

// Run initial command
func (m MainMenuModel) Init() tea.Cmd {
	return nil
}

// Process event messages
func (m MainMenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit
		}
	}

	return m, nil
}

// Render view
func (m MainMenuModel) View() string {
	if m.quitting {return ""}
	return "Hauptmenü\nDrücken Sie q zum Beenden"
}