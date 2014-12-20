package htmlutil

import (
	"bytes"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func GetValue(n *html.Node) string {
	switch n.DataAtom {
	case atom.Meta:
		return strings.TrimSpace(GetAttr(n, "content"))
	case atom.Time:
		if a := strings.TrimSpace(GetAttr(n, "datetime")); a != "" {
			return a
		}
	case atom.Abbr:
		if full := GetAttr(n, "title"); full != "" {
			return full
		}
	}
	return strings.TrimSpace(GetText(n))
}

func GetText(n *html.Node) string {
	var text bytes.Buffer

	Walk(n, func(n *html.Node) {
		if n.Type == html.TextNode {
			text.WriteString(n.Data)
		}
	}, func(n *html.Node) {
		if n.Type == html.ElementNode {
			if Blocks.Has(n.DataAtom) {
				text.WriteString("\n\n")
			} else {
				text.WriteString(" ")
			}
		}
	})
	return text.String()
}
