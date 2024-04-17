// Elisa: AI Programming Tutor for Students – CLI
// © 2024 Dennis Schulmeister-Zimolong <dennis@wpvs.de>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.

package ui

import tea "github.com/charmbracelet/bubbletea"

// The terminal screen width
var Width int = 0

// The terminal screen height
var Height int = 0

// Default message handler for all view screens. Handles messages like terminal
// resizing or Ctrl+C to exit the application. Returns nil, nil in most cases,
// so that the caller may proceed as normal. But if anything else is returned,
// the caller should return these, instead.
func Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		// Handle window resize
		Width = msg.Width
		Height = msg.Height
	case tea.KeyMsg:
		// Handle custom keys
		keypress := msg.String()

		if keypress == "ctrl+c" {
			// Exit application immediately
			return nil, tea.Quit
		}
	}

	return nil, nil
}
