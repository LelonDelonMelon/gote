// internal/ui/ui.go
package ui

import (
	"github.com/gdamore/tcell"
	"gote/internal/editor"
)

// UI represents the user interface for the text editor.
type UI struct {
	screen   tcell.Screen
	editor   *editor.Editor
	inputStr string // String to store user input
}

// NewUI initializes a new UI instance.
func NewUI(screen tcell.Screen, editor *editor.Editor) *UI {
	return &UI{
		screen:   screen,
		editor:   editor,
		inputStr: "",
	}
}

// ... (rest of the code remains the same)

// handleInputChar handles the input of printable characters.
func (ui *UI) handleInputChar(char rune) {
	ui.inputStr += string(char)
	ui.editor.InsertChar(char)
	ui.editor.MoveCursorRight() // Assuming there's a MoveCursorRight method in your editor
}

// handleBackspace handles the backspace key.
func (ui *UI) handleBackspace() {
	if len(ui.inputStr) > 0 {
		ui.inputStr = ui.inputStr[:len(ui.inputStr)-1]
		ui.editor.DeleteChar() // Assuming there's a DeleteChar method in your editor
		ui.editor.MoveCursorLeft() // Assuming there's a MoveCursorLeft method in your editor
	}
}

// handleEnter handles the enter key.
func (ui *UI) handleEnter() {
	ui.inputStr += "\n"
	ui.editor.InsertChar('\n') // Assuming there's an InsertChar method in your editor
	ui.editor.MoveCursorRight() // Assuming there's a MoveCursorRight method in your editor
}

// redraw updates the screen with the current content.
func (ui *UI) redraw() {
	ui.screen.Clear()
	ui.screen.Show()

	// Draw text content
	ui.drawEditorContent()

	// Draw input string
	ui.drawInputString()

	ui.screen.Show()
}

// drawEditorContent draws the text content of the editor.
func (ui *UI) drawEditorContent() {
	editorContent := ui.editor.GetBuffer() // Assuming there's a GetBuffer method in your editor
	x, y := 2, 4
	for _, char := range editorContent {
		ui.screen.SetContent(x, y, char, nil, tcell.StyleDefault)
		x++
		if char == '\n' {
			x = 2
			y++
		}
	}

	// Draw cursor manually using the cursor position
	cursorPos := ui.editor.GetCursorPos() // Assuming there's a GetCursorPos method in your editor
	ui.screen.SetContent(cursorPos+2, y, '█', nil, tcell.StyleDefault)
}

// drawInputString draws the user input string.
func (ui *UI) drawInputString() {
	_, screenHeight := ui.screen.Size()
	x, y := 2, screenHeight-2
	for _, char := range ui.inputStr {
		ui.screen.SetContent(x, y, char, nil, tcell.StyleDefault)
		x++
	}
}
// handleKeyEvent handles specific key events.
func (ui *UI) handleKeyEvent(ev *tcell.EventKey) {
    switch ev.Key() {
    case tcell.KeyRune:
        ui.handleInputChar(ev.Rune())
    case tcell.KeyBackspace, tcell.KeyBackspace2:
        ui.handleBackspace()
    case tcell.KeyEnter:
        ui.handleEnter()
    }

    // Redraw the screen with updated content
    ui.redraw()
}


// EventLoop manages the main event loop for the UI.
func (ui *UI) EventLoop() {
    for {
        ev := ui.screen.PollEvent()
        switch ev := ev.(type) {
        case *tcell.EventKey:
            // Handle key events
            if ev.Key() == tcell.KeyCtrlC {
                return
            }

            ui.handleKeyEvent(ev)
        }
    }
}
