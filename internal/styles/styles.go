package styles

import "github.com/charmbracelet/lipgloss"

var (
	BaseStyle = lipgloss.NewStyle().
			Width(80).
			Foreground(lipgloss.Color("7")).
			TabWidth(2)
	HeaderStyle = BaseStyle.
			Bold(true).
			Foreground(lipgloss.Color("12"))
	TextStyle = BaseStyle.
			Foreground(lipgloss.Color("10"))
	SynonymStyle = BaseStyle.
			Foreground(lipgloss.Color("10"))
	AntonymStyle = BaseStyle.
			Foreground(lipgloss.Color("9"))
	ExampleStyle = BaseStyle.
			PaddingLeft(2)
	WordStyle = HeaderStyle.
			Foreground(lipgloss.Color("15"))
)
