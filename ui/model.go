package ui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	list     list.Model
	selected string
	quit     bool
}

func New() model {
	const (
		defaultWidth = 20
		listHeight   = 14
	)

	listMenu := list.New(
		[]list.Item{
			item{Key: "a", Value: "AA"},
			item{Key: "b", Value: "BBB"},
		},
		itemDelegate{},
		defaultWidth,
		listHeight,
	)

	return model{
		list: listMenu,
	}
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quit = true
			return m, tea.Quit

		case "enter":
			i, ok := m.list.SelectedItem().(item)
			if ok {
				m.selected = i.Value
			}

			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	if m.quit {
		return "thanks"
	}

	return "\n" + m.list.View()
}
