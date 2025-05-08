package ui

import (
	"log"

	"github.com/JonasLindermayr/PasswordManager/lib"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	titleStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("21"))
	keywordStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("203"))
	subtleStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
)

const (
	menuView uint = iota 
	createView
	listView
	passwordView
)

type model struct {
	state uint
	currPassword lib.Password
	passwords []lib.Password
	store 	*lib.Store
	textInputName textinput.Model
	textInputPassword textinput.Model
	textInputPasswordRepeat textinput.Model
	listIndex	int
	version string
	choicesMainMenu []string
	choicesCursor int
}

func NewModel(store *lib.Store, version string) model {

	passwords, err := store.GetPasswords()
	if (err != nil) {
		log.Fatal(err)
	}

	return model{
		store: store,
		state: menuView,
		passwords: passwords,
		textInputName: textinput.New(),
		textInputPassword: textinput.New(),
		textInputPasswordRepeat: textinput.New(),
		version: version,
		choicesMainMenu: []string{"Store a new Password", "View stored Passwords"},
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmds []tea.Cmd
		cmd tea.Cmd
	)

	m.textInputName, cmd = m.textInputName.Update(msg)
	cmds = append(cmds, cmd)
	m.textInputPassword, cmd = m.textInputPassword.Update(msg)
	cmds = append(cmds, cmd)
	m.textInputPasswordRepeat, cmd = m.textInputPasswordRepeat.Update(msg)
	cmds = append(cmds, cmd)

	switch msg := msg.(type) {
		case tea.KeyMsg:
			key := msg.String()
			switch m.state {
				case menuView:
					switch key {
					case "up", "k":
						if m.choicesCursor > 0 {
							m.choicesCursor--
						}
					case "down", "j":
						if m.choicesCursor  < len(m.choicesMainMenu)-1 {
							m.choicesCursor ++
						}
					case "enter", " ":
						switch m.choicesCursor {
							case 0:
								m.state = createView
							case 1:
								m.state = listView
						}
						case "q", "ctrl+c":
							return m, tea.Quit
					}
				case createView:
					switch key {
						case "g":
						// TODO: GENERATE PASSWORD
						case "esc":
							m.state = menuView
					}

				case passwordView:
					switch key {
						case "esc":
							m.state = menuView
					}

				case listView:
					switch key {
						case "esc":
							m.state = menuView
						case "n":
							m.textInputName.SetValue("")
							m.textInputPassword.SetValue("")
							m.textInputPasswordRepeat.SetValue("")
							m.textInputName.Focus()
							m.currPassword = lib.Password{}
							m.state = createView
						case "up", "k":
							if m.listIndex > 0 {
								m.listIndex--
							}
						case "down", "j":
							if m.listIndex < len(m.passwords)-1 {
								m.listIndex++
							}
						case "enter":
							m.currPassword = m.passwords[m.listIndex]
							m.state = passwordView
							m.textInputName.SetValue(m.currPassword.Name)
							m.textInputPassword.Focus()
							m.textInputPassword.CursorEnd()
						
					}
			}
	
		}
	return m, tea.Batch(cmds...)
}
		