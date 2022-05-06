package gtree

import (
	"bufio"
	"encoding/json"
	"io"

	toml "github.com/pelletier/go-toml/v2"
	"gopkg.in/yaml.v2"
)

// 関心事はtreeを出力すること
type spreader interface {
	spread(io.Writer, []*Node) error
}

func newSpreader(encode encode) spreader {
	if s, ok := spreaders[encode]; ok {
		return s
	}
	return &defaultSpreader{}
}

// NOTE: 微妙。分岐は無くせるし追加は楽そうだが、グローバルに持ってしまってる。
var spreaders = map[encode]spreader{
	encodeDefault: &defaultSpreader{},
	encodeJSON:    &jsonSpreader{},
	encodeTOML:    &tomlSpreader{},
	encodeYAML:    &yamlSpreader{},
}

type encode int

const (
	encodeDefault encode = iota
	encodeJSON
	encodeYAML
	encodeTOML
)

type defaultSpreader struct{}

func (ds *defaultSpreader) spread(w io.Writer, roots []*Node) error {
	branches := ""
	for _, root := range roots {
		branches += ds.spreadBranch(root, "")
	}
	return ds.write(w, branches)
}

func (*defaultSpreader) spreadBranch(current *Node, out string) string {
	out += current.prettyBranch()
	for _, child := range current.children {
		out = (*defaultSpreader)(nil).spreadBranch(child, out)
	}
	return out
}

func (*defaultSpreader) write(w io.Writer, in string) error {
	buf := bufio.NewWriter(w)
	if _, err := buf.WriteString(in); err != nil {
		return err
	}
	return buf.Flush()
}

type jsonSpreader struct{}

func (*jsonSpreader) spread(w io.Writer, roots []*Node) error {
	enc := json.NewEncoder(w)
	for _, root := range roots {
		jRoot := root.toJSONNode(nil)
		if err := enc.Encode(jRoot); err != nil {
			return err
		}
	}
	return nil
}

type jsonNode struct {
	Name     string      `json:"value"`
	Children []*jsonNode `json:"children"`
}

func (parent *Node) toJSONNode(jParent *jsonNode) *jsonNode {
	if len(parent.children) == 0 {
		return nil
	}
	if jParent == nil {
		jParent = &jsonNode{Name: parent.name}
	}

	jParent.Children = make([]*jsonNode, len(parent.children))
	for i := range parent.children {
		jParent.Children[i] = &jsonNode{Name: parent.children[i].name}
		parent.children[i].toJSONNode(jParent.Children[i])
	}

	return jParent
}

type tomlSpreader struct{}

func (*tomlSpreader) spread(w io.Writer, roots []*Node) error {
	enc := toml.NewEncoder(w)
	for _, root := range roots {
		tRoot := root.toTOMLNode(nil)
		if err := enc.Encode(tRoot); err != nil {
			return err
		}
	}
	return nil
}

type tomlNode struct {
	Name     string      `toml:"value"`
	Children []*tomlNode `toml:"children"`
}

func (parent *Node) toTOMLNode(tParent *tomlNode) *tomlNode {
	if len(parent.children) == 0 {
		return nil
	}
	if tParent == nil {
		tParent = &tomlNode{Name: parent.name}
	}

	tParent.Children = make([]*tomlNode, len(parent.children))
	for i := range parent.children {
		tParent.Children[i] = &tomlNode{Name: parent.children[i].name}
		parent.children[i].toTOMLNode(tParent.Children[i])
	}

	return tParent
}

type yamlSpreader struct{}

func (*yamlSpreader) spread(w io.Writer, roots []*Node) error {
	enc := yaml.NewEncoder(w)
	for _, root := range roots {
		yRoot := root.toYAMLNode(nil)
		if err := enc.Encode(yRoot); err != nil {
			return err
		}
	}
	return nil
}

type yamlNode struct {
	Name     string      `yaml:"value"`
	Children []*yamlNode `yaml:"children"`
}

func (parent *Node) toYAMLNode(yParent *yamlNode) *yamlNode {
	if len(parent.children) == 0 {
		return nil
	}
	if yParent == nil {
		yParent = &yamlNode{Name: parent.name}
	}

	yParent.Children = make([]*yamlNode, len(parent.children))
	for i := range parent.children {
		yParent.Children[i] = &yamlNode{Name: parent.children[i].name}
		parent.children[i].toYAMLNode(yParent.Children[i])
	}

	return yParent
}
