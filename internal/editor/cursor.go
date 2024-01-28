// internal/editor/cursor.go
package editor

// Cursor represents the cursor position in the text editor.
type Cursor struct {
	Line   int // Line number (if applicable)
	Column int // Column number
}

// NewCursor initializes a new Cursor instance.
func NewCursor(line, column int) *Cursor {
	return &Cursor{
		Line:   line,
		Column: column,
	}
}

// MoveCursorLeft moves the cursor to the left.
func (c *Cursor) MoveCursorLeft() {
	if c.Column > 0 {
		c.Column--
	}
}

// MoveCursorRight moves the cursor to the right.
func (c *Cursor) MoveCursorRight() {
	c.Column++
}

// MoveCursorUp moves the cursor up (if applicable).
func (c *Cursor) MoveCursorUp() {
	if c.Line > 0 {
		c.Line--
	}
}

// MoveCursorDown moves the cursor down (if applicable).
func (c *Cursor) MoveCursorDown() {
	c.Line++
}

// SetCursor sets the cursor to the specified position.
func (c *Cursor) SetCursor(line, column int) {
	c.Line = line
	c.Column = column
}

