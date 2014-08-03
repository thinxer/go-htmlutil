package htmlutil

import (
	"code.google.com/p/go.net/html"
	"code.google.com/p/go.net/html/atom"
)

func Walk(n *html.Node, before, after func(*html.Node)) {
	if before != nil {
		before(n)
	}
	var next *html.Node
	for c := n.FirstChild; c != nil; c = next {
		next = c.NextSibling
		Walk(c, before, after)
	}
	if after != nil {
		after(n)
	}
}

func FindOne(n *html.Node, a atom.Atom) *html.Node {
	if n.DataAtom == a {
		return n
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if r := FindOne(c, a); r != nil {
			return r
		}
	}
	return nil
}

func FindId(n *html.Node, id string) *html.Node {
	if GetAttr(n, "id") == id {
		return n
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if r := FindId(c, id); r != nil {
			return r
		}
	}
	return nil
}

func FindOneFunc(n *html.Node, f func(*html.Node) bool) *html.Node {
	if f(n) {
		return n
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if r := FindOneFunc(c, f); r != nil {
			return r
		}
	}
	return nil
}

func FindAllFunc(n *html.Node, f func(*html.Node) bool) (nodes []*html.Node) {
	if f == nil || f(n) {
		nodes = append(nodes, n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		nodes = append(nodes, FindAllFunc(c, f)...)
	}
	return
}

func Children(n *html.Node) (ns []*html.Node) {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ns = append(ns, c)
	}
	return
}

func FindTags(n *html.Node, f func(*html.Node), atoms ...atom.Atom) {
	x := map[atom.Atom]struct{}{}
	for _, atom := range atoms {
		x[atom] = struct{}{}
	}
	Walk(n, func(n *html.Node) {
		if _, ok := x[n.DataAtom]; ok {
			f(n)
		}
	}, nil)
}
