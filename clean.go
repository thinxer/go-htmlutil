package htmlutil

import (
	"bufio"
	"bytes"
	"io"

	"code.google.com/p/go.net/html"
	"code.google.com/p/go.net/html/atom"
	"code.google.com/p/go.net/html/charset"
	"code.google.com/p/go.text/transform"
)

// Clean and parse the HTML, so that:
// 0. Encoding converted to UTF8 if not already.
// 1. Comment and DOCTYPE removed.
// 2. Code blocks (script, style, link) removed.
func ParseClean(r io.Reader) (*html.Node, error) {
	buffered := bufio.NewReader(r)
	peek, _ := buffered.Peek(1024)
	e, _, certain := charset.DetermineEncoding(peek, "text/html")
	if !certain {
		e, _ = charset.Lookup("UTF-8")
	}
	r = transform.NewReader(buffered, e.NewDecoder())
	// Parse the document.
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	// Rewrite the document.
	var clean func(*html.Node)
	clean = func(n *html.Node) {
		switch n.Type {
		case html.ElementNode:
			if n.DataAtom == atom.Script || n.DataAtom == atom.Style || n.DataAtom == atom.Link {
				Drop(n)
			} else if n.DataAtom == atom.Noscript {
				// It is complicated whether to extract the Noscript tag.
				// Sometimes Noscript contains only hints that you should enable JavaScript.
				// Sometimes Noscript contains the real document.
				// The current policy is that we extract the tag only if it is longer than 512 bytes.
				data := GetText(n)
				if len(data) > 512 {
					nodes, err := html.ParseFragment(bytes.NewReader([]byte(data)), n.Parent)

					if err == nil {
						for _, node := range nodes {
							n.Parent.InsertBefore(node, n)
							Walk(node, clean, nil)
						}
					}
				}
				Drop(n)
			} else if n.DataAtom == atom.Meta {
				// Drop HTTP-EQUIVs.
				if GetAttr(n, "http-equiv") != "" {
					Drop(n)
				}
			}
		case html.CommentNode, html.DoctypeNode:
			Drop(n)
		case html.TextNode:
			// Do nothing.
		case html.ErrorNode:
			// XXX Never seen this yet. Drop this?
			panic(n)
		}
	}
	Walk(doc, clean, nil)
	return doc, nil
}
