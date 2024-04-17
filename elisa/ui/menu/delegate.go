// Elisa: AI Programming Tutor for Students – CLI
// © 2024 Dennis Schulmeister-Zimolong <dennis@wpvs.de>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.

package menu

import (
	"fmt"
	"io"
	"strings"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

// Delegate to support the rendering of menu items
type MenuItemDelegate struct {}

// Always reserve one line for menu items
func (d MenuItemDelegate) Height() int {
	return 1
}

// Don't add extra spacing aroung menu items
func (d MenuItemDelegate) Spacing() int {
	return 0
}

// Updating menu items is not supported (do nothing)
func (d MenuItemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd {
	return nil
}

// Render menu item
func (d MenuItemDelegate) Render(writer io.Writer, model list.Model, index int, listItem list.Item) {
	menuItem, ok := listItem.(MenuItem)
	if !ok {return}
	
	selectionHand := "  "
	label         := menuItem.Label
	styleLabel    := menuItem.menu.ItemStyleLabel

	if index == model.Index() {
		if len(menuItem.menu.SelectionHand) > 0 {
			selectionHand = " " + string([]rune(menuItem.menu.SelectionHand)[0])
		} else {
			selectionHand = " →"
		}

		styleLabel   = menuItem.menu.ItemStyleSelected
	}

	if menuItem.Disabled {
		selectionHand = "  "
		styleLabel    = menuItem.menu.ItemStyleDisabled
	}

	if menuItem.Separator {
		selectionHand = ""
		label   = ""
	}

	result := strings.Join([]string{
		styleLabel.Render(selectionHand),
		styleLabel.Render(label),
	}, " ")

	fmt.Fprint(writer, result)
}