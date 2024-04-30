// Elisa: AI Programming Tutor for Students – CLI
// © 2024 Dennis Schulmeister-Zimolong <dennis@wpvs.de>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.

package ui

import (
	"strings"

	"github.com/charmbracelet/bubbles/key"
)

type Screen struct {
	Title       string        // Screen title visible on the top
	Footer      string        // Optional footer text visible on the bottom
	Message     string        // Optional message visible below the footer
	KeyBindings []key.Binding // Optional key bindings visible below the message
}

// Can be replaced with built-in max() function in Go 1.21
func _max(i, j int) int {
	if i > j {return i}
	return j
}

// Get the reserved height occupied by the header. Normally this is one line, except
// when the header contains line-breaks, followed by an empty-line.
func (screen Screen) HeaderHeight() int {
	return strings.Count(screen.Title, "\n") + 1
}

// Get the reserved height occupied by the footer. Normally this is one empty line
// plus one line for footer text and message, plus another line for the key bindings.
// Additional lines will be used, if the footer text or message contains line-breaks.
func (screen Screen) FooterHeight() int {
	footerHeight := strings.Count(screen.Footer, "\n")
	messageHeight := strings.Count(screen.Message, "\n")

	keyBindingsHeight := 0
	if len(screen.KeyBindings) > 0 {keyBindingsHeight = 1}

	return 1 + _max(footerHeight, messageHeight) + keyBindingsHeight
}

// Get the reserved height occupied by header and footer.
func (screen Screen) ReservedHeight() int {
	return screen.HeaderHeight() + screen.FooterHeight()
}

// Render screen header
func (screen Screen) ViewHeader() string {
	// TODO
	return ""
}

// Render blank lines to push the footer to the bottom of the screen. contentHeight is the number
// of lines that have already been used by the content between header and footer. header and footer
// declare, if the height of header and footer must be taken into account, because ViewHeader() and
// ViewFooter() are also called.
func (screen Screen) ViewFillHeight(contentHeight int, header bool, footer bool) string {
	lineCount := Height - contentHeight;
	if header {lineCount -= screen.HeaderHeight()}
	if footer {lineCount -= screen.FooterHeight()}

	return strings.Repeat("\n", lineCount)
}

// Render screen footer
func (screen Screen) ViewFooter() string {
	// TODO
	return ""
}
