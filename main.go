// main.go
package main

import (
	"github.com/gdamore/tcell"
	"gote/internal/editor"
	"gote/internal/ui"
)

func main() {
	// Create a new text editor
	textEditor := editor.NewEditor()

	// Initialize the Tcell screen
	screen, err := tcell.NewScreen()
	if err != nil {
		// Handle error
		panic(err)
	}

	if err := screen.Init(); err != nil {
		// Handle error
		panic(err)
	}

	defer screen.Fini()

	// Create a new UI instance
	textUI := ui.NewUI(screen, textEditor)


	// Start the event loop
	textUI.EventLoop()
}
