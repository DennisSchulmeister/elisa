// Elisa: AI Programming Tutor for Students – CLI
// © 2024 Dennis Schulmeister-Zimolong <dennis@wpvs.de>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.

package select_files

import (
	tea "github.com/charmbracelet/bubbletea"
)

// Temporary screen where the user can select one ore more files of the own source code
// to include them in their dialouge with the AI LLM.
type SelectFilesModel struct {}

// Create new instance of the model to show a file selection list on screen
func New() SelectFilesModel {
	return SelectFilesModel{}
}

// Run initial command
func (m SelectFilesModel) Init() tea.Cmd {
	return nil
}

// Process event messages
func (m SelectFilesModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

// Render view
func (m SelectFilesModel) View() string {
	return "Dateien auswählen"
}