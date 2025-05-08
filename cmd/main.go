package main

import (
	"log"

	"github.com/JonasLindermayr/PasswordManager/lib"
	"github.com/JonasLindermayr/PasswordManager/ui"
	tea "github.com/charmbracelet/bubbletea"
)

var version = "0.0.1"

func main() {

	store := new(lib.Store)
	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	m := ui.NewModel(store, version)
	if _, err := tea.NewProgram(m).Run(); err != nil {
		log.Fatal(err)
	}


}

