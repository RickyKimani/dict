package scripts

import (
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
	"os"
)

var (
	width, height, _ = term.GetSize(int(os.Stdin.Fd())) // Get terminal size

	focusedStyle           = lipgloss.NewStyle().Foreground(lipgloss.Color("212")) // Pink
	unfocusedStyle         = lipgloss.NewStyle().Foreground(lipgloss.Color("245")) // Grey
	slightlyUnfocusedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#C8A2C8"))
	headerStyle            = lipgloss.NewStyle().Foreground(lipgloss.Color("90"))
	linkStyle              = lipgloss.NewStyle().Foreground(lipgloss.Color("33"))

	containerStyle = lipgloss.NewStyle().Width(width).Height(height).Align(lipgloss.Center, lipgloss.Center)
)

func (m Model) View() string {
	var inputView, viewportView, footer string

	h := "Simply a dictionary\n\n"
	header := headerStyle.Render(h)

	if m.focused {
		inputView = focusedStyle.Render(m.textInput.View())
		viewportView = slightlyUnfocusedStyle.Render(m.viewport.View())
	} else {
		inputView = slightlyUnfocusedStyle.Render(m.textInput.View())
		viewportView = focusedStyle.Render(m.viewport.View())
	}

	footnote := "Results fetched from urban dictionary are most likely satirical.\n"
	f := "\n\nesc - quit || enter - search || tab - shift focus || ↑/↓ - scroll\n\n" +
		footnote + "\n\n"

	link := linkStyle.Render("github.com/rickyKimani")
	footer = unfocusedStyle.Render(f)

	ui := lipgloss.JoinVertical(lipgloss.Center, header, inputView, viewportView, footer, link)

	return containerStyle.Render(ui) // Center everything
}
