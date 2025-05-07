package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	C_DOTCHAR = " • "
	C_MAINPAGE = "main"
	C_CREATEPAGE = "create"
	C_VIEWPAGE = "retrieve"
	C_HELPPAGE = "help"
)

var (
	titleStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("21"))
	keywordStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("203"))
	subtleStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
	dotStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("236")).Render(C_DOTCHAR)
)

type model struct {
	currentPage string
	choices []string
	cursor int
	redirectTo string
	redirect bool
	version string
}


func InitialModel(version string) model {
	return model{
		choices: []string{"Store a new Password", "View stored Passwords", "Help"},
		currentPage: C_MAINPAGE,
		version: version,
		redirect: false,
	}
}

func (m model) Init() tea.Cmd {
	return tea.HideCursor
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg: // key Press
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", " ":
			m.redirect = true
			switch m.cursor {
				case 0:
					m.redirectTo = "create" 
				case 1:
					m.redirectTo = "retrieve" 
				case 2:
					m.redirectTo = "help" 
			}
		case "b":
			if (m.currentPage != C_MAINPAGE) {
				m.redirectTo = "main"
				fmt.Print("Test");
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	s := titleStyle.Render("PasswordManager v" + m.version) + "\n\n"
	for i, choice := range m.choices {

        cursor := " "
        if m.cursor == i {
            cursor = ">" 
			s += fmt.Sprintf("%s %s\n", keywordStyle.Render(cursor), keywordStyle.Render(choice))
        } else {
			s += fmt.Sprintf("%s %s\n", cursor, choice)
		}
    }

    s += "\n" + dotStyle + subtleStyle.Render("q, esc: quit") + dotStyle
	if (m.currentPage != C_MAINPAGE) {
		s += dotStyle + subtleStyle.Render("b: back") + dotStyle
	}

    return s
}


func updateMainMenu(msg tea.Msg, m model) (tea.Model, tea.Cmd) {
	return m, nil
}

func updateSecondMenu(msg tea.Msg, m model) (tea.Model, tea.Cmd) {
	return m, nil
}