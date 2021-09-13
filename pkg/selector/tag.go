package selector

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

type TagSelector struct {
	tag string
}

func (t TagSelector) Match(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == t.tag
}

type TagParser struct {
	SelParser
}

func NewTagParser(sel string) *TagParser {
	return &TagParser{
		SelParser{
			sel:    sel,
			selLen: len(sel),
		},
	}
}

func (t *TagParser) Parse() (Sel, error) {
	if !(t.isValidTagNameChar(t.sel[t.pos]) || t.sel[t.pos] == '\\') {
		return nil, fmt.Errorf("expected id selector (key), found '%c'", t.sel[t.pos])
	}

	tag := ""
	for t.pos < t.selLen {
		char := t.sel[t.pos]

		switch {
		case t.isValidTagNameChar(char): // get current "i" if is a valid name character
			tag += string(char)
			t.pos++
			break
		case char == '\\': // sel have an escaped element https://drafts.csswg.org/css-syntax-3/#escaping
			c, err := t.parseEscape()
			if err != nil {
				return nil, err
			}
			tag += c
			break
		case char == '\r':
			t.pos++
			if t.sel[t.pos] == '\n' {
				t.pos++ // for end of lines \r\n
			}
			break
		case char == ' ' || char == '\n' || char == '\t':
			t.pos++
			break
		}
	}

	return &TagSelector{
		tag: strings.ToLower(tag),
	}, nil
}
