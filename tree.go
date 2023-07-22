//go:build !wasm

package gtree

import "io"

type iTree interface {
	output(io.Writer, io.Reader, *config) error
	outputProgrammably(io.Writer, *Node, *config) error
	mkdir(io.Reader, *config) error
	mkdirProgrammably(*Node, *config) error
}
