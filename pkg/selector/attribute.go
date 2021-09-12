package selector

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

type AttrSelector struct {
	key string
	op  string
	val string
}

func (t AttrSelector) Match(n *html.Node) bool {
	if html.ElementNode == n.Type {
		for _, attr := range n.Attr {
			switch t.op {
			case "=":
				if t.key == attr.Key {
					values := strings.Split(attr.Val, " ")
					for _, value := range values {
						if t.val == value {
							return true
						}
					}
				}
				return false
				break
			case "~=":
				if t.key == attr.Key {
					values := strings.Split(attr.Val, " ")
					for _, value := range values {
						if t.val == value {
							return true
						}
					}
				}
				return false
				break
			case "|=":
				if t.key == attr.Key {
					if strings.Index(attr.Val, t.val) == 0 {
						return true
					}
				}
				return false
				break
			case "^=":
				if t.key == attr.Key {
					value := attr.Val[:len(attr.Val)]

					if value == attr.Val {
						return true
					}
				}
				return false
				break
			case "$=":
				if t.key == attr.Key {
					value := attr.Val[len(attr.Val)-len(t.val):]

					if value == attr.Val {
						return true
					}
				}
				return false
				break
			case "*=":
				if t.key == attr.Key {
					if strings.Contains(attr.Val, t.val) {
						return true
					}
				}
				return false
				break
			default:
				if t.key == attr.Key {
					return true
				}
			}
		}
	}

	return false
}

type AttributeParser struct {
	SelParser
}

func NewAttributeParser(sel string) *AttributeParser {
	return &AttributeParser{
		SelParser{
			sel:    sel,
			selLen: len(sel),
		},
	}
}

func (a AttributeParser) Parse() (Sel, error) {
	if a.sel[a.pos] != '[' {
		return nil, fmt.Errorf("expected attribute selector ([attr=val]), found '%c'", a.sel[a.pos])
	}
	a.pos++

	key := ""
	op := ""
	var val string
	for a.pos < a.selLen {
		char := a.sel[a.pos]

		switch {
		case a.isValidTagNameChar(char): // get current "i" if is a valid name character
			key += string(char)
			a.pos++
			break
		case char == '=':
			op = string(char)
			a.pos++
			break
		case char == '~':
			if a.sel[a.pos+1] == '=' {
				op = a.sel[a.pos : a.pos+2]
			} else {
				return nil, fmt.Errorf("expected operation for attribute selector ([~,|,^,$,*]=), found '%s'",
					a.sel[a.pos:a.pos+2])
			}
			a.pos += 2
			break
		case char == '|':
			if a.sel[a.pos+1] == '=' {
				op = a.sel[a.pos : a.pos+2]
			} else {
				return nil, fmt.Errorf("expected operation for attribute selector ([~,|,^,$,*]=), found '%s'",
					a.sel[a.pos:a.pos+2])
			}
			a.pos += 2
			break
		case char == '^':
			if a.sel[a.pos+1] == '=' {
				op = a.sel[a.pos : a.pos+2]
			} else {
				return nil, fmt.Errorf("expected operation for attribute selector ([~,|,^,$,*]=), found '%s'",
					a.sel[a.pos:a.pos+2])
			}
			a.pos += 2
			break
		case char == '$':
			if a.sel[a.pos+1] == '=' {
				op = a.sel[a.pos : a.pos+2]
			} else {
				return nil, fmt.Errorf("expected operation for attribute selector ([~,|,^,$,*]=), found '%s'",
					a.sel[a.pos:a.pos+2])
			}
			a.pos += 2
			break
		case char == '*':
			if a.sel[a.pos+1] == '=' {
				op = a.sel[a.pos : a.pos+2]
			} else {
				return nil, fmt.Errorf("expected operation for attribute selector ([~,|,^,$,*]=), found '%s'",
					a.sel[a.pos:a.pos+2])
			}
			a.pos += 2
			break
		case char == '"':
			a.pos++
			found := false
			value := ""
			for i := a.pos; !found; i++ {
				char := a.sel[a.pos]

				switch {
				case a.isValidTagNameChar(char):
					value += string(char)
					a.pos++
					break
				case char == '"':
					found = true
					a.pos++
					break
				case char == '\\':
					c, err := a.parseEscape()
					if err != nil {
						return nil, err
					}
					value += c
					break
				}
			}
			val = value
			break
		case char == '\\': // sel have an escaped element https://drafts.csswg.org/css-syntax-3/#escaping
			c, err := a.parseEscape()
			if err != nil {
				return nil, err
			}
			key += c
			break
		case char == '\r':
			a.pos++
			if a.sel[a.pos] == '\n' {
				a.pos++ // for end of lines \r\n
			}
			break
		case char == ' ' || char == '\n' || char == '\t' || char == ']':
			a.pos++
			break
		}
	}

	return &AttrSelector{
		key: key,
		op:  op,
		val: val,
	}, nil
}
