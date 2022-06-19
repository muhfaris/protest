package ui

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	styleItem int = iota
	styleSelectedItem
)

type item struct {
	Key   string
	Value string
}

func (this item) FilterValue() string { return "" }

type itemDelegate struct{}

func NewItemDelegate() *itemDelegate {
	return &itemDelegate{}
}

func (d itemDelegate) Height() int                               { return 1 }
func (d itemDelegate) Spacing() int                              { return 0 }
func (d itemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i.Value)

	fn := menuStyle.Render
	if index == m.Index() {
		fn = func(s string) string {
			//d.selectedStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#3BACB6"))
			return selectedMenuStyle.Render(s)
		}
	}

	fmt.Fprintf(w, fn(str))
}
