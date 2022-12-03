package gtree

import (
	"io"

	"github.com/fatih/color"
)

type Program struct {
	stdout io.Writer
	stderr io.Writer // TODO: まだ使ってない.そもそもいらないかも

	cfg *config
}

func NewProgram(stdout, stderr io.Writer) *Program {
	return &Program{
		stdout: stdout,
		stderr: stderr,
		cfg: &config{
			lastNodeFormat: &branchFormat{
				directly:   "└──",
				indirectly: "    ",
			},
			intermedialNodeFormat: &branchFormat{
				directly:   "├──",
				indirectly: "│   ",
			},
			space:  spacesTab,
			encode: encodeDefault,
		},
	}
}

// Output outputs tree to Program.stdout.
// This function requires node generated by NewRoot function.
func (p *Program) Output(root *Node) error {
	if root == nil {
		return ErrNilNode
	}
	if !root.isRoot() {
		return ErrNotRoot
	}

	idxCounter.reset()

	tree := newTree(p.cfg, []*Node{root})
	if err := tree.grow(); err != nil {
		return err
	}
	return tree.spread(p.stdout)
}

// Mkdir makes directories.
// This function requires node generated by NewRoot function.
func (p *Program) Mkdir(root *Node) error {
	if root == nil {
		return ErrNilNode
	}
	if !root.isRoot() {
		return ErrNotRoot
	}

	idxCounter.reset()

	tree := newTree(p.cfg, []*Node{root})
	tree.enableValidation()
	// when detect invalid node name, return error. process end.
	if err := tree.grow(); err != nil {
		return err
	}
	if p.cfg.dryrun {
		// when detected no invalid node name, output tree.
		return tree.spread(color.Output)
	}
	// when detected no invalid node name, no output tree.
	return tree.mkdir()
}

// FormatIntermedialNode is for branch format.
func (p *Program) FormatIntermedialNode(directly, indirectly string) *Program {
	p.cfg.intermedialNodeFormat.directly = directly
	p.cfg.intermedialNodeFormat.indirectly = indirectly
	return p
}

// FormatLastNode is for branch format.
func (p *Program) FormatLastNode(directly, indirectly string) *Program {
	p.cfg.lastNodeFormat.directly = directly
	p.cfg.lastNodeFormat.indirectly = indirectly
	return p
}

// EncodeJSON is for output json format.
func (p *Program) EncodeJSON() *Program {
	p.cfg.encode = encodeJSON
	return p
}

// EncodeYAML is for output yaml format.
func (p *Program) EncodeYAML() *Program {
	p.cfg.encode = encodeYAML
	return p
}

// EncodeTOML is for output toml format.
func (p *Program) EncodeTOML() *Program {
	p.cfg.encode = encodeTOML
	return p
}

// DryRun is for dry run. Detects node that is invalid for directory generation.
func (p *Program) DryRun() *Program {
	p.cfg.dryrun = true
	return p
}

// FileExtensions is for creating as a file instead of a directory.
func (p *Program) FileExtensions(extensions []string) *Program {
	p.cfg.fileExtensions = extensions
	return p
}
