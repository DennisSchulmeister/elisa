// Elisa: AI Programming Tutor for Students – CLI
// © 2024 Dennis Schulmeister-Zimolong <dennis@wpvs.de>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.

package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/DennisSchulmeister/elisa/elisa/views/main_menu"
)

// Main function with BubbleTea boilerplate
func main() {
	m := main_menu.New()
	p := tea.NewProgram(&m)

    if _, err := p.Run(); err != nil {
        fmt.Printf("Fehler: %v", err)
        os.Exit(1)
    }
}