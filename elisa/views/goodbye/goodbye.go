// Elisa: AI Programming Tutor for Students – CLI
// © 2024 Dennis Schulmeister-Zimolong <dennis@wpvs.de>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.

package goodbye

import (
	"github.com/DennisSchulmeister/elisa/elisa/ui"
	tea "github.com/charmbracelet/bubbletea"
)

// Final goodbye screen when the program exits
type GoodbyeModel struct {}

// Create new instance of the goodbmodel to show a goodbye message
func NewGoodybeModel() GoodbyeModel {
	return GoodbyeModel{}
}

// Run initial command
func (goodbye GoodbyeModel) Init() tea.Cmd {
	return nil
}

// Process event messages
func (goodbye GoodbyeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	model, command := ui.Update(msg)
	if model != nil || command != nil {return model, command}

	return goodbye, nil
}

// Render view
func (goodbye GoodbyeModel) View() string {
	return "Auf Wiedersehen!"
}