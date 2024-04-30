// Elisa: AI Programming Tutor for Students – CLI
// © 2024 Dennis Schulmeister-Zimolong <dennis@wpvs.de>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.

package ui

import (
	"fmt"
	"strings"

	"github.com/DennisSchulmeister/elisa/elisa/global"
	"github.com/charmbracelet/bubbles/key"
)

type Page struct {
	Title       string        // Page title visible on the top
	Footer      []string      // Optional footer text visible on the bottom
	Message     string        // Optional message visible below the footer
	KeyBindings []key.Binding // Optional key bindings visible below the message
}

// Get the reserved height occupied by the header, which is one line. Note that
// line breaks in the title are not handled and may disrupt the layout.
func (page Page) HeaderHeight() int {
	return 1
}

// Get the reserved height occupied by the footer. Normally this is one line for footer
// text and message, plus another line for the key bindings. Note that line-breaks in
// any of the footer content are not considered and may disrupt the layout.
func (page Page) FooterHeight() int {
	keyBindingsHeight := 0
	if len(page.KeyBindings) > 0 {keyBindingsHeight = 1}

	return 1 + keyBindingsHeight
}

// Get the reserved height occupied by header and footer.
func (page Page) ReservedHeight() int {
	return page.HeaderHeight() + page.FooterHeight()
}

// Render page header
func (page Page) ViewHeader() string {
	title   := fmt.Sprintf(" %s ", page.Title)
	program := fmt.Sprintf(" %s %s ", global.PROGRAM, global.VERSION)
	padding := strings.Repeat(" ", max(Width - len(title) - len(program), 0))

	return PageTitleLabelStyle.Render(title) + PageTitleBgStyle.Render(padding + program)
}

// Render blank lines to push the footer to the bottom of the screen. contentHeight is the number
// of lines that have already been used by the content between header and footer. header and footer
// declare, if the height of header and footer must be taken into account, because ViewHeader() and
// ViewFooter() are also called.
func (page Page) ViewMain(main string, header bool, footer bool) string {
	lineCount := Height - strings.Count(main, "\n") - 1;
	if header {lineCount -= page.HeaderHeight()}
	if footer {lineCount -= page.FooterHeight()}

	return main + strings.Repeat("\n", max(lineCount, 0))
}

// Render page footer
func (page Page) ViewFooter() string {
	message := page.Message
	if len(message) > 0 {message = " ❰ " + message + " ❱ "}

	fields  := fmt.Sprintf(" %s ", strings.Join(page.Footer, " | "))
	padding := strings.Repeat(" ", max(Width - len(message) - len(fields), 0))
	line1   := PageFooterMsgStyle.Render(message + padding) + PageFooterFieldStyle.Render(fields)
	line2   := ""

	return line1 + line2
}
