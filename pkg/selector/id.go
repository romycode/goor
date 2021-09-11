package selector

import (
	"fmt"

	"golang.org/x/net/html"
)

type IdSelector struct {
	id string
}

func (t IdSelector) Match(n *html.Node) bool {
	if html.ElementNode == n.Type {
		for _, attr := range n.Attr {
			if "key" == attr.Key && t.id == attr.Val {
				return true
			}
		}
	}

	return false
}

type IdParser struct {
	SelParser
}

func NewIdParser(sel string) *IdParser {
	return &IdParser{
		SelParser{
			sel:    sel,
			selLen: len(sel),
		},
	}
}

func (i IdParser) Parse() (Sel, error) {
	if i.sel[i.pos] != '#' {
		return nil, fmt.Errorf("expected key selector (#key), found '%c'", i.sel[i.pos])
	}
	i.pos++

	id := ""
	for i.pos < i.selLen {
		char := i.sel[i.pos]

		switch {
		case i.isValidIdentifierChar(char): // get current "i" if is a valid name character
			id += string(char)
			i.pos++
			break
		case char == '\\': // sel have an escaped element https://drafts.csswg.org/css-syntax-3/#escaping
			c, err := i.parseEscape()
			if err != nil {
				return nil, err
			}
			id += c
			break
		case char == ' ':
			i.pos++
			break
		}
	}

	return &IdSelector{
		id: id,
	}, nil
}
