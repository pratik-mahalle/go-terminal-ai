package gui

import (
	"github.com/mahalle-pratik/go-terminal-ai/internal/terminal"
	"github.com/rivo/tview"
)

// StartGUI initializes and starts the GUI
func StartGUI() {
	app := tview.NewApplication()
	terminal.StartTerminal(app)
}