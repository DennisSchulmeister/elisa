// Elisa: AI Programming Tutor for Students – CLI
// © 2024 Dennis Schulmeister-Zimolong <dennis@wpvs.de>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.

package menu

import (
	"github.com/DennisSchulmeister/elisa/elisa/ui"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Selection menu
type Menu struct {
	Title				string				// Menu title
	items               []MenuItem			// Menu items
	list				list.Model			// BubbleTea list widget
	SelectionHand       string				// Prefix for selected items
	SelectedIndex		int					// Value of currently selected item
	prevSelectedIndex	int					// Value of previously selected item
	ChosenIndex			int					// Value of actually chosen item
	ExitModel           tea.Model           // View model to show after exiting the menu
	ExitCommands        []tea.Cmd			// Command to execute to exit the menu
	ItemStyleLabel		lipgloss.Style		// Style to render the items label
	ItemStyleDisabled	lipgloss.Style		// Style to render disabled items
	ItemStyleSelected   lipgloss.Style		// Style to render selected items
}

type MenuOption func(m *Menu)

// Create new menu
func NewMenu(options ...MenuOption) *Menu {
	menu := &Menu{
		Title:             "Menü",
		SelectionHand:     "→",
		SelectedIndex:     -1,
		prevSelectedIndex: -1,
		ChosenIndex:       -1,
		ExitCommands: 	   []tea.Cmd{tea.Quit},

		ItemStyleLabel:    ui.LabelStyle,
		ItemStyleSelected: ui.SelectedStyle,
		ItemStyleDisabled: ui.DisabledStyle,
	}

	for _, option := range options {
		option(menu)
	}
	
	var listItems []list.Item

	for i, item := range menu.items {
		item.menu = menu

		if !item.Disabled && !item.Separator && menu.SelectedIndex < 0 {
			menu.SelectedIndex = i
		}

		listItems = append(listItems, item)
	}

	menu.list = list.New(listItems, MenuItemDelegate{}, 0, 0)
	menu.list.SetShowTitle(false)
	menu.list.SetShowStatusBar(false)
	menu.list.SetSize(ui.Width, ui.Height)

	return menu
}

// Set title of new menu
func WithTitle(title string) MenuOption {
	return func(m *Menu) { m.Title = title }
}

// Set items of new menu
func WithItems(items []MenuItem) MenuOption {
	return func(m *Menu) { m.items = items }
}

// Set selection hand of new menu
func WithSelectionHand(selectionHand string) MenuOption {
	return func(m *Menu) { m.SelectionHand = selectionHand }
}

// Set view model to show after exiting the menu
func WithExitModel(exitModel tea.Model) MenuOption {
	return func(m *Menu) { m.ExitModel = exitModel }
}

// Set command to run when exiting the menu
func WithExitCommands(exitCommands []tea.Cmd) MenuOption {
	return func(m *Menu) { m.ExitCommands = exitCommands }
}

// Set normal item style of new menu
func WithItemStyleLabel(style lipgloss.Style) MenuOption {
	return func(m *Menu) { m.ItemStyleLabel = style }
}

// Set selected item style of new menu
func WithItemStyleSelected(style lipgloss.Style) MenuOption {
	return func(m *Menu) { m.ItemStyleSelected = style }
}

// Set disabled item style of new menu
func WithItemStyleDisabled(style lipgloss.Style) MenuOption {
	return func(m *Menu) { m.ItemStyleDisabled = style }
}

// Initial command not supprted
func (menu Menu) Init() tea.Cmd {
	return nil
}

// Process event messages
func (menu Menu) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	model, command := ui.Update(msg)
	if model != nil || command != nil {return model, command}

	menu.list.SetSize(ui.Width, ui.Height)

	switch msg := msg.(type) {
		case tea.KeyMsg:
			// Handle custom keys
			keypress := msg.String()

			if keypress == "q" || keypress == "esc" {
				// Exit menu
				return menu.ExitModel, tea.Sequence(menu.ExitCommands...)
			} else if keypress == "enter" {
				// Choose and execute item
				menu.ChosenIndex = menu.SelectedIndex
				item := menu.items[menu.SelectedIndex]

				var _model tea.Model = menu
				var _command tea.Cmd = nil

				if item.Model   != nil { _model = item.Model }
				if item.Command != nil { _command = item.Command }

				return _model, _command
			}
	}

	// Delegate messages to list widget
	var cmd tea.Cmd
	menu.list, cmd = menu.list.Update(msg)

	// Skip separators and disabled items
	menu.SelectedIndex = menu.list.Index()
	direction := 0

	if menu.SelectedIndex > menu.prevSelectedIndex {
		direction = 1
	} else if menu.SelectedIndex < menu.prevSelectedIndex {
		direction = -1
	}

	for direction != 0 {
		if menu.SelectedIndex < 0 {
			menu.SelectedIndex = 0
			break
		} else if menu.SelectedIndex > len(menu.items) - 1 {
			menu.SelectedIndex = len(menu.items) - 1
			break
		}

		selectedItem := menu.items[menu.SelectedIndex]

		if selectedItem.Separator || selectedItem.Disabled {
			menu.SelectedIndex += direction
		} else {
			break
		}
	}

	menu.list.Select(menu.SelectedIndex)
	menu.prevSelectedIndex = menu.SelectedIndex

	// Return new model and command
	return menu, cmd
}

// Render menu
func (menu Menu) View() string {
	result := ui.ScreenTitleStyle.Render(menu.Title) + "\n"
	result += menu.list.View()
	return result
}