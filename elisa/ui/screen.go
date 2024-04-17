// Elisa: AI Programming Tutor for Students – CLI
// © 2024 Dennis Schulmeister-Zimolong <dennis@wpvs.de>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.

package ui

import "github.com/charmbracelet/bubbles/key"

type Screen struct {
    Title       string                     // Screen title visible on the top
    Footer      string                     // Optional footer text visible on the bottom
    Message     string                     // Optional message visible below the footer
    KeyBindings []key.Binding              // Optional key bindings visible below the message
}

// Get the reserved height that is occupied by the screen title, footer, etc.
func (screen Screen) ReservedHeight() int {
    return 0 // TODO
}

// Render screen header
func (screen Screen) ViewHeader() string {
    // TODO
    return ""
}

// Render blank lines to push the footer to the bottom of the screen. `contentHeight` is the number
// of lines that have already been used by the content between header and footer. `header` and `footer`
// declare, if the hight of header and footer are visible on screen and must therefor be taken into account.
func (screen Screen) ViewFillHeight(contentHeight int, header bool, footer bool) string {
    // TODO
    return ""
}

// Render screen footer
func (screen Screen) ViewFooter() string {
    // TODO
    return ""
}