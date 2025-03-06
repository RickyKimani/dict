package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/rickykimani/dict/scripts"
	"log"
)

func main() {
	p := tea.NewProgram(scripts.NewModel(), tea.WithAltScreen())
	_, err := p.Run()
	if err != nil {
		log.Fatal(err)
	}
}
