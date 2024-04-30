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
	"strconv"

	"github.com/DennisSchulmeister/elisa/elisa/view/login"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

const PROGRAM = "elisa"
const VERSION = "0.0.1"

// Structure for the command line options
type options struct {
	debug  bool   // Log to debug.out
	server string // Adress of the Elisa server
}

var Options options = options{}

// Root command of the Corba library, which is used to parse the command line options.
// This library supports the program command -args argument style. But we are using
// an anonymous root command only, as currently we don't need sub-commands.
var rootCommand = &cobra.Command{
	Use:   PROGRAM,
	Short: "AI Programmiertutor*in für Studierende",
	Long: fmt.Sprintf(`%s %s – AI Programmiertutor*in für Studierende

Elisa ist eine intelligente Lernhilfe für Studierende im Informatik-Studium. Sie beantwortet Fragen
zu den Übungsaufgaben, zu deinem Quellcode und hilft dir bei den Übungsaufgaben. Das Studium hat sie
schon seit ein paar Jahren abgeschlossen. Jetzt hilft sie dir dabei, dein Studium zu meistern.

Elisa ist die Urenkel*in von Joseph Weizenbaums Eliza. Aktuell sind noch nicht alle Funktionen ausgereift.
Es handelt sich um ein Experiment für Studierende an der DHBW Karlsruhe, das im Rahmen des Forschungsprojekts
KoLLI durchgeführt wird. Fragen, Wünsche, Anregungen, Fehlermeldungen sind stets willkommen. Schicke hierzu
einfach eine E-Mail an: dennis.schulmeister-zimolong@dhbw-karlsruhe.de`, PROGRAM, VERSION),

	// Run the application
	Run: func(cmd *cobra.Command, args []string) {
		// Read .env file
		godotenv.Load()

		// Enable debug logs
		if !Options.debug {
			debug_, err := strconv.ParseBool(os.Getenv("DEBUG"))
			if err == nil {
				Options.debug = debug_
			}
		}

		if Options.debug {
			f, err := tea.LogToFile("debug.log", "debug")
			if err != nil {
				panic(err)
			}
			defer f.Close()
		}

		// Start TUI
		m := login.NewLoginModel()
		p := tea.NewProgram(m)

		if _, err := p.Run(); err != nil {
			panic(err)
		}
	},
}

// Init function called by GoLang
func init() {
	rootCommand.Version = VERSION

	rootCommand.PersistentFlags().BoolVarP(&Options.debug, "debug", "d", false, "Logausgaben in die Datei debug.log schreiben")
	rootCommand.PersistentFlags().StringVarP(&Options.server, "server", "s", "https://elisa.zimolong.eu", "Adresse des Elisa-Servers")
}

// Main function called by GoLang
func main() {
	rootCommand.Execute()
}
