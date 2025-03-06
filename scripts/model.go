package scripts

import (
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/x/term"
	"os"
)

type Model struct {
	textInput textinput.Model
	viewport  viewport.Model
	terms     Terms
	err       error
	focused   bool // Track whether the text input is active
}

func NewModel() Model {
	t := textinput.New()
	t.Placeholder = "Enter a search term"
	t.Focus()

	width, height, err := term.GetSize(os.Stdout.Fd())
	if err != nil {
		width = 80
		height = 24
	}

	vp := viewport.New(width, height-15) // Default viewport size

	return Model{
		textInput: t,
		viewport:  vp,
		focused:   true, // Start with text input focused
	}
}
