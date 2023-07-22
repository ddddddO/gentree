//go:build !wasm

package gtree

import (
	"errors"
	"io"
)

var (
	idxCounter = newCounter()
)

var (
	// ErrNilNode is returned if the argument *gtree.Node of OutputProgrammably / MkdirProgrammably function is nill.
	ErrNilNode = errors.New("nil node")
	// ErrNotRoot is returned if the argument *gtree.Node of OutputProgrammably / MkdirProgrammably function is not root of the tree.
	ErrNotRoot = errors.New("not root node")
)

// OutputProgrammably outputs tree to w.
// This function requires node generated by NewRoot function.
func OutputProgrammably(w io.Writer, root *Node, options ...Option) error {
	if root == nil {
		return ErrNilNode
	}
	if !root.isRoot() {
		return ErrNotRoot
	}

	idxCounter.reset()
	cfg := newConfig(options)

	tree := newTreeSimple(cfg)
	if cfg.massive {
		tree = newTreePipeline(cfg)
	}
	return tree.outputProgrammably(w, root, cfg)
}

var (
	// ErrExistPath is returned if the argument *gtree.Node of MkdirProgrammably function is path already exists.
	ErrExistPath = errors.New("path already exists")
)

// MkdirProgrammably makes directories.
// This function requires node generated by NewRoot function.
func MkdirProgrammably(root *Node, options ...Option) error {
	if root == nil {
		return ErrNilNode
	}
	if !root.isRoot() {
		return ErrNotRoot
	}

	idxCounter.reset()
	cfg := newConfig(options)

	tree := newTreeSimple(cfg)
	if cfg.massive {
		tree = newTreePipeline(cfg)
	}
	return tree.mkdirProgrammably(root, cfg)
}

// NewRoot creates a starting node for building tree.
func NewRoot(text string) *Node {
	return newNode(text, rootHierarchyNum, idxCounter.next())
}

// Add adds a node and returns an instance of it.
// If a node with the same text already exists in the same hierarchy of the tree, that node will be returned.
func (parent *Node) Add(text string) *Node {
	if child := parent.findChildByText(text); child != nil {
		return child
	}

	current := newNode(text, parent.hierarchy+1, idxCounter.next())
	current.setParent(parent)
	parent.addChild(current)
	return current
}
