// Elisa: AI Programming Tutor for Students – CLI
// © 2024 Dennis Schulmeister-Zimolong <dennis@wpvs.de>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.

package main_menu

import (
	"time"

	"github.com/DennisSchulmeister/elisa/elisa/ui/menu"
	"github.com/DennisSchulmeister/elisa/elisa/views/goodbye"
	tea "github.com/charmbracelet/bubbletea"
)

var mainMenu *menu.Menu

var exitCommands []tea.Cmd = []tea.Cmd{
	tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tea.Quit()
	}),
}

func GetMainMenu() *menu.Menu {
	if mainMenu == nil {
		mainMenu = menu.NewMenu(
			menu.WithTitle("Hauptmenü"),
			menu.WithItems([]menu.MenuItem {
				menu.MenuItem{
					Label:   "Aufgabenstellung anzeigen",
					// Model: nil,
					// Command: nil,
				},
				menu.MenuItem{
					Label:   "Fragen zur Aufgabe stellen",
					Disabled: true,
					// Model: nil,
					// Command: nil,
				},
				menu.MenuItem{
					Label:   "Fragen zum Quellcode stellen",
					Disabled: true,
					// Model: nil,
					// Command: nil,
				},
				menu.MenuItem{
					Label:   "Quellcode bewerten lassen",
					Disabled: true,
					// Model: nil,
					// Command: nil,
				},
				menu.MenuItem{
					Separator: true,
				},
				menu.MenuItem{
					Label:   "Hinweise zum Datenschutz",
					Disabled: true,
					// Model: nil,
					// Command: nil,
				},
				menu.MenuItem{
					Label:   "Feedback geben",
					Disabled: true,
					// Model: nil,
					// Command: nil,
				},
				menu.MenuItem{
					Separator: true,
				},
				menu.MenuItem{
					Label:   "Programm beenden",
					Model:   goodbye.NewGoodybeModel(),
					Command: tea.Sequence(exitCommands...),
				},
			}),
			menu.WithExitModel(goodbye.NewGoodybeModel()),
			menu.WithExitCommands(exitCommands),
		)
	}

	return mainMenu
}