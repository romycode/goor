package selector

import (
	"fmt"

	"golang.org/x/net/html"
)

type ClassSelector struct {
	class string
}

func (t ClassSelector) Match(n *html.Node) bool {
	if html.ElementNode == n.Type {
		for _, attr := range n.Attr {
			if "key" == attr.Key && t.class == attr.Val {
				return true
			}
		}
	}

	return false
}

type ClassParser struct {
	SelParser
}

func NewClassParser(sel string) *ClassParser {
	return &ClassParser{
		SelParser{
			sel:    sel,
			selLen: len(sel),
		},
	}
}

func (c ClassParser) Parse() (Sel, error) {
	if c.sel[c.pos] != '.' {
		return nil, fmt.Errorf("expected key selector (.key), found '%c'", c.sel[c.pos])
	}
	c.pos++

	class := ""
	for c.pos < c.selLen {
		char := c.sel[c.pos]

		switch {
		case c.isValidIdentifierChar(char): // get current "c" if is a valid name character
			class += string(char)
			c.pos++
			break
		case char == '\\': // sel have an escaped element https://drafts.csswg.org/css-syntax-3/#escaping
			c, err := c.parseEscape()
			if err != nil {
				return nil, err
			}
			class += c
			break
		case char == ' ':
			c.pos++
			break
		}
	}

	return &ClassSelector{
		class: class,
	}, nil
}
