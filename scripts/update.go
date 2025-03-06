package scripts

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/x/term"
	"github.com/muesli/reflow/wordwrap"
	"os"
	"strings"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			return m, tea.Quit
		case "enter":
			m.focused = false // Unfocus text input when searching
			t := m.textInput.Value()
			return m, GetResults(t) // Fetch results
		case "tab": // Toggle focus between text input and viewport
			m.focused = !m.focused
			return m, nil
		}

		// **Handle scrolling only when text input is NOT focused**
		if !m.focused {
			switch msg.String() {
			case "up", "k":
				m.viewport.LineUp(1)
				return m, nil
			case "down", "j":
				m.viewport.LineDown(1)
				return m, nil
			}
		}
	}

	// **Handle API response (TermsMsg)**
	if termsMsg, ok := msg.(TermsMsg); ok {
		if termsMsg.err != nil {
			m.err = termsMsg.err
			return m, nil
		}
		m.terms = termsMsg.terms

		//var definitions string
		//for _, term := range m.terms.List {
		//	definitions += fmt.Sprintf("%s\n\n", term.Definition)
		//}
		//m.viewport.SetContent(definitions)
		//m.viewport.GotoTop()

		w, _, _ := term.GetSize(os.Stdout.Fd())
		if w == 0 {
			w = 80
		}

		padding := 4
		indent := strings.Repeat(" ", padding+3)
		wrapWidth := w - padding - 1

		var formattedDefinitions []string
		for i, trm := range m.terms.List {
			wrapped := wordwrap.String(trm.Definition, wrapWidth)
			wrappedIndented := strings.ReplaceAll(wrapped, "\n", "\n"+indent)
			formattedDefinitions = append(formattedDefinitions, fmt.Sprintf("%02d. %s\n\n", i+1, wrappedIndented))
		}
		m.viewport.SetContent(strings.Join(formattedDefinitions, ""))
		m.viewport.GotoTop() // Ensure viewport resets to the top

		return m, nil
	}

	// **Update text input only when focused**
	if m.focused {
		m.textInput, cmd = m.textInput.Update(msg)
		cmds = append(cmds, cmd)
	}

	m.viewport, cmd = m.viewport.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}
