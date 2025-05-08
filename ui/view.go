package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	appNameStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("21")).Padding(0, 1)
	faint = lipgloss.NewStyle().Foreground(lipgloss.Color("255")).Faint(true)
	listEnumeratorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("99")).MarginRight(1)
)

func (m model) View() string {
	s := appNameStyle.Render("PasswordManager v" + m.version) + "\n\n"

	if m.state == menuView {
		for i, choice := range m.choicesMainMenu {

			cursor := " "
			if m.choicesCursor == i {
				cursor = ">" 
				s += fmt.Sprintf("%s %s\n", keywordStyle.Render(cursor), keywordStyle.Render(choice))
			} else {
				s += fmt.Sprintf("%s %s\n", cursor, choice)
			}
		}
		 
		s += "\n\n"
		s += faint.Render("• q, ctrl+c - quit •")
	}

	if m.state == createView {
		s += "Create a new Password:\n\n"
		s += m.textInputName.View() + "\n\n"
		s += m.textInputPassword.View() + "\n\n"
		s += m.textInputPasswordRepeat.View() + "\n\n"
		s += faint.Render("ctrl+s - save • esc - discard")
	}

	if m.state == listView {
		for i, n := range m.passwords {
			prefix := " "
			if i == m.listIndex {
				prefix = ">"
			}
			shortBody := strings.ReplaceAll(n.Name, "\n", " ")
			if len(shortBody) > 30 {
				shortBody = shortBody[:30]
			}
			s += listEnumeratorStyle.Render(prefix) + n.Name + " | " + faint.Render(shortBody) + "\n\n"
		}
		s += faint.Render("n - new password • q - quit")
	}

	return s
}