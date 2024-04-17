// Elisa: AI Programming Tutor for Students – CLI
// © 2024 Dennis Schulmeister-Zimolong <dennis@wpvs.de>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.

package menu

import (
	tea "github.com/charmbracelet/bubbletea"
)

// Menu item
type MenuItem struct {
	Label		string		// Displayed label
	Separator	bool		// Item is a separator wihtout text
	Disabled	bool		// Item is inactive and cannot be selected
	Command     tea.Cmd		// Command function to execute when the item is chosen
	Model       tea.Model	// New tea.Model (view) to use when the item is chosen
	menu		*Menu		// Parent menu
}

// Allow searching and filtering by label
func (menuItem MenuItem) FilterValue() string {
	if !menuItem.Separator {
		return menuItem.Label
	} else {
		return ""
	}
}
