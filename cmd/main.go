package main

import (
	"fmt"
	"os"

	"github.com/JonasLindermayr/PasswordManager/ui"
	tea "github.com/charmbracelet/bubbletea"
)

var version = "0.0.1"

func main() {
	// TODO: Check if db exists -> if not create one with master password
	// TODO: Check if db exists -> if check master password is correct

	p := tea.NewProgram(ui.InitialModel(version))
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

