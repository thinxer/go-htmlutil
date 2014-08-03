package htmlutil

import (
	"strings"

	"code.google.com/p/go.net/html/atom"
)

type AtomSet map[atom.Atom]struct{}

func NewAtomSet(atoms []atom.Atom) AtomSet {
	s := AtomSet{}
	for _, t := range atoms {
		s[t] = struct{}{}
	}
	return s
}

func NewAtomSetFromString(s string) AtomSet {
	atoms := []atom.Atom{}
	for _, token := range strings.Fields(s) {
		a := atom.Lookup([]byte(token))
		if a > 0 {
			atoms = append(atoms, a)
		}
	}
	return NewAtomSet(atoms)
}

func (s AtomSet) Has(a atom.Atom) bool {
	if s == nil {
		return false
	}
	_, ok := s[a]
	return ok
}

func (s AtomSet) HasName(name []byte) bool {
	a := atom.Lookup(name)
	if a == 0 {
		return false
	}
	return s.Has(a)
}

var (
	Codes      = NewAtomSetFromString("style script link")
	Blocks     = NewAtomSetFromString("address article aside audio blockquote br canvas dd div dl fieldset figcaption figure figcaption footer form h1 h2 h3 h4 h5 h6 header hgroup hr noscript ol output p pre section table tfoot ul video")
	Containers = NewAtomSetFromString("article aside div footer header hgroup section")
	Headers    = NewAtomSetFromString("h1 h2 h3 h4 h5 h6 h7")
	Voids      = NewAtomSetFromString("area base br col embed hr img input keygen link meta param source track wbr")
)
