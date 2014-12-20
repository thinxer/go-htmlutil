package htmlutil

import "golang.org/x/net/html"

func GetAttr(n *html.Node, name string) string {
	for _, attr := range n.Attr {
		if attr.Key == name {
			return attr.Val
		}
	}
	return ""
}

func SetAttr(n *html.Node, name, val string) {
	for i := range n.Attr {
		if n.Attr[i].Key == name {
			n.Attr[i].Val = val
			return
		}
	}
	n.Attr = append(n.Attr, html.Attribute{Key: name, Val: val})
}

func DelAttr(n *html.Node, name string) {
	i := 0
	for j := range n.Attr {
		if n.Attr[j].Key != name {
			n.Attr[i] = n.Attr[j]
			i++
		}
	}
	n.Attr = n.Attr[:i]
}
