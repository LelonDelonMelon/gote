// internal/editor/buffer.go
package editor


// RopeNode represents a node in the rope data structure.
type RopeNode struct {
	Value    string
	Left     *RopeNode
	Right    *RopeNode
	Weight   int // Number of characters in the subtree rooted at this node
	Balance  int // Balance factor for balancing the tree during edits
}

// NewRopeNode initializes a new RopeNode with the given value.
func NewRopeNode(value string) *RopeNode {
	return &RopeNode{
		Value: value,
		Weight: len(value),
		Balance: 0,
	}
}

// TextBuffer represents the text buffer using a rope data structure.
type TextBuffer struct {
	root *RopeNode
}

// NewTextBuffer initializes a new TextBuffer instance.
func NewTextBuffer() *TextBuffer {
	return &TextBuffer{
		root: nil,
	}
}

// Insert inserts text into the buffer at the specified position.
func (b *TextBuffer) Insert(position int, text string) {
	if b.root == nil {
		b.root = NewRopeNode(text)
	} else {
		b.root = b.insertNode(b.root, position, text)
	}
}

// Helper function to insert a node into the rope data structure.
func (b *TextBuffer) insertNode(node *RopeNode, position int, text string) *RopeNode {
	if node == nil {
		return NewRopeNode(text)
	}

	leftWeight := 0
	if node.Left != nil {
		leftWeight = node.Left.Weight
	}

	if position <= leftWeight {
		node.Left = b.insertNode(node.Left, position, text)
	} else {
		node.Right = b.insertNode(node.Right, position-leftWeight-1, text)
	}

	// Update weights and balance factor
	node.Weight = leftWeight + len(node.Value) + node.Balance + b.calculateWeight(node.Right)
	node.Balance = b.calculateBalance(node.Left) - b.calculateBalance(node.Right)

	// Rebalance the tree if necessary
	if node.Balance > 1 {
		if b.calculateBalance(node.Left) >= 0 {
			node = b.rotateRight(node)
		} else {
			node.Left = b.rotateLeft(node.Left)
			node = b.rotateRight(node)
		}
	} else if node.Balance < -1 {
		if b.calculateBalance(node.Right) <= 0 {
			node = b.rotateLeft(node)
		} else {
			node.Right = b.rotateRight(node.Right)
			node = b.rotateLeft(node)
		}
	}

	return node
}

// Helper function to calculate the weight of a node.
func (b *TextBuffer) calculateWeight(node *RopeNode) int {
	if node == nil {
		return 0
	}
	return node.Weight
}

// Helper function to calculate the balance factor of a node.
func (b *TextBuffer) calculateBalance(node *RopeNode) int {
	if node == nil {
		return 0
	}
	return b.calculateWeight(node.Left) - b.calculateWeight(node.Right)
}

// Helper function to perform a right rotation on a node.
func (b *TextBuffer) rotateRight(node *RopeNode) *RopeNode {
	newRoot := node.Left
	node.Left = newRoot.Right
	newRoot.Right = node

	// Update weights and balance factors
	node.Weight = b.calculateWeight(node.Left) + len(node.Value) + b.calculateWeight(node.Right)
	newRoot.Weight = b.calculateWeight(newRoot.Left) + len(newRoot.Value) + b.calculateWeight(newRoot.Right)

	return newRoot
}

// Helper function to perform a left rotation on a node.
func (b *TextBuffer) rotateLeft(node *RopeNode) *RopeNode {
	newRoot := node.Right
	node.Right = newRoot.Left
	newRoot.Left = node

	// Update weights and balance factors
	node.Weight = b.calculateWeight(node.Left) + len(node.Value) + b.calculateWeight(node.Right)
	newRoot.Weight = b.calculateWeight(newRoot.Left) + len(newRoot.Value) + b.calculateWeight(newRoot.Right)

	return newRoot
}

// GetText returns the text content of the buffer.
func (b *TextBuffer) GetText() string {
	if b.root == nil {
		return ""
	}
	return b.root.Value
}

