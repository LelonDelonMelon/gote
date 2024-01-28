// internal/editor/editor.go
package editor

// Editor represents the core logic of the text editor.
type Editor struct {
	buffer    []rune // Text buffer
	cursorPos int    // Cursor position in the buffer
}

// NewEditor initializes a new Editor instance.
func NewEditor() *Editor {
	return &Editor{
		buffer:    make([]rune, 0),
		cursorPos: 0,
	}
}

// InsertChar inserts a character at the current cursor position.
func (e *Editor) InsertChar(char rune) {
	// Ensure cursor position is within the buffer bounds
	if e.cursorPos < 0 {
		e.cursorPos = 0
	} else if e.cursorPos > len(e.buffer) {
		e.cursorPos = len(e.buffer) - 1
	}

	// Insert the character at the cursor position
	e.buffer = append(e.buffer[:e.cursorPos], append([]rune{char}, e.buffer[e.cursorPos:]...)...)
	e.cursorPos++
}

// DeleteChar deletes the character at the current cursor position.
func (e *Editor) DeleteChar() {
	if e.cursorPos < 0 || e.cursorPos >= len(e.buffer) {
		// Cursor is out of bounds, nothing to delete
		return
	}

	if e.cursorPos == len(e.buffer)-1 {
		e.buffer = e.buffer[:e.cursorPos]
	} else {
		e.buffer = append(e.buffer[:e.cursorPos], e.buffer[e.cursorPos+1:]...)

	}
}

// MoveCursorLeft moves the cursor to the left.
func (e *Editor) MoveCursorLeft() {
	if e.cursorPos > 0 {
		e.cursorPos--
	}
}

// MoveCursorRight moves the cursor to the right.
func (e *Editor) MoveCursorRight() {
	if e.cursorPos < len(e.buffer) {
		e.cursorPos++
	}
}

// GetBuffer returns the current text buffer.
func (e *Editor) GetBuffer() string {
	return string(e.buffer)
}

// GetCursorPos returns the current cursor position.
func (e *Editor) GetCursorPos() int {
	return e.cursorPos
}
