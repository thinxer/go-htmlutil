package htmlutil

import (
	"fmt"
	"strings"

	"code.google.com/p/go.net/html"
	"code.google.com/p/go.net/html/atom"
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

func Drop(n *html.Node) {
	if n.Parent != nil {
		n.Parent.RemoveChild(n)
	}
}

func Create(tag string) *html.Node {
	var s string
	if Voids.HasName([]byte(tag)) {
		s = "<" + tag + ">"
	} else {
		s = fmt.Sprintf("<%s></%s>", tag, tag)
	}
	fragments, err := html.ParseFragment(strings.NewReader(s), bodyContext)
	if err != nil {
		return nil
	}
	return fragments[0]
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

var bodyContext *html.Node

func init() {
	doc, _ := html.Parse(strings.NewReader(`<html><body></body></html>`))
	bodyContext = FindOne(doc, atom.Body)
}
