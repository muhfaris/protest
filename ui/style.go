package ui

import "github.com/charmbracelet/lipgloss"

var (
	menuStyle         = lipgloss.NewStyle().PaddingLeft(2)
	selectedMenuStyle = lipgloss.NewStyle().PaddingLeft(2).Bold(true).Foreground(lipgloss.Color("#3BACB6"))
)
