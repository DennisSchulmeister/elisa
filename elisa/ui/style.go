// Elisa: AI Programming Tutor for Students – CLI
// © 2024 Dennis Schulmeister-Zimolong <dennis@wpvs.de>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.

package ui

import "github.com/charmbracelet/lipgloss"

// For color codes see: https://github.com/fidian/ansi
var PageTitleBgStyle    lipgloss.Style = lipgloss.NewStyle().Background(lipgloss.Color("18")).Foreground(lipgloss.Color("15"))
var PageTitleLabelStyle lipgloss.Style = lipgloss.NewStyle().Background(lipgloss.Color("25")).Foreground(lipgloss.Color("15"))
var PageFooterMsgStyle  lipgloss.Style = lipgloss.NewStyle().Background(lipgloss.Color("16")).Foreground(lipgloss.Color("15"))
var PageFooterFieldStyle lipgloss.Style = lipgloss.NewStyle().Background(lipgloss.Color("31")).Foreground(lipgloss.Color("7"))
var PageFooterKeyStyle = PageTitleBgStyle.Copy()

var LabelStyle          lipgloss.Style = lipgloss.NewStyle()
var DisabledStyle       lipgloss.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("253"))
var SelectedStyle       lipgloss.Style = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("170"))
var BlurredStyle        lipgloss.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("61"))
var CursorStyle         lipgloss.Style = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("170"))
var HelpStyle           lipgloss.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))