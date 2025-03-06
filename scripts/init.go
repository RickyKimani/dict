package scripts

import (
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"golang.org/x/term"
	"os"
)

func (m Model) Init() tea.Cmd {
	width, height, _ := term.GetSize(int(os.Stdin.Fd()))
	if width == 0 {
		width = 80 // Default width
	}
	if height == 0 {
		height = 24 // Default height
	}

	m.viewport = viewport.New(width, height-5) // Make viewport full width, leave space for input

	return tea.Batch(textinput.Blink, tea.SetWindowTitle("A dictionary for the terminal"))
}
