package htmlutil

import (
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func Lift(n *html.Node) {
	if n.Parent == nil {
		return
	}
	for {
		c := n.FirstChild
		if c == nil {
			break
		}
		n.RemoveChild(c)
		n.Parent.InsertBefore(c, n)
	}
	Drop(n)
}

func Drop(n *html.Node) *html.Node {
	if n.Parent != nil {
		n.Parent.RemoveChild(n)
	}
	return n
}

func Create(tag string) *html.Node {
	return &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.Lookup([]byte(tag)),
		Data:     tag,
	}
}

func Replace(original, repl *html.Node) {
	if original.Parent == nil {
		panic("replacing not possible: parent is nil")
	}
	if original == repl {
		return
	}
	original.Parent.InsertBefore(repl, original)
	original.Parent.RemoveChild(original)
}

// BodyContext can be used as the context when parsing fragments.
var BodyContext *html.Node

func init() {
	doc, _ := html.Parse(strings.NewReader(`<html><body></body></html>`))
	BodyContext = FindOne(doc, atom.Body)
}
