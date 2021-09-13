package selector

import (
	"fmt"
	"strconv"

	"golang.org/x/net/html"
)

type Sel interface {
	Match(n *html.Node) bool
}

type Parser interface {
	Parse(sel string) Sel
}

type SelParser struct {
	sel    string
	selLen int
	pos    int
}

// parseEscape parses backslash escaped character (formats '\000026' or '\26 ' and 'U+000026' or 'U+0026')
func (s *SelParser) parseEscape() (string, error) {
	if '\\' != s.sel[s.pos] {
		return "", fmt.Errorf("expected escape element (\\), found '%c'", s.sel[s.pos])
	}
	s.pos++

	start := s.pos // this is for the pattern "\000026" and also works with "\26 "
	if s.sel[start] == 'U' {
		start = s.pos + 2 // this is for the pattern "\U+000026"
	}

	var i int
	for i = start; i < start+6 && i < s.selLen && s.isHexChar(s.sel[i]); i++ {
	}

	if i-start < 6 && s.sel[i] != ' ' {
		s.pos += i - start
		return s.sel[start:i], nil
	}

	v, err := strconv.ParseUint(s.sel[start:i], 16, 21)
	if err != nil {
		return "", err
	}

	s.pos = i

	return string(rune(v)), nil
}

// isValidTagNameChar checks if is valid character for id name
// as defined in https://html.spec.whatwg.org/dev/syntax.html#syntax-tag-name
func (s SelParser) isValidTagNameChar(char byte) bool {
	return 'a' <= char && char <= 'z' || 'A' <= char && char <= 'Z' || '0' <= char && char <= '9'
}

// isValidIdentifierChar checks if is valid character for selector
// as defined in https://www.w3.org/TR/CSS2/syndata.html#characters and https://drafts.csswg.org/selectors-4/#case-sensitive
func (s SelParser) isValidIdentifierChar(char byte) bool {
	return s.isValidTagNameChar(char) || char == '_' || char == '-' || char > 127
}

// isHexChar checks if is a hexadecimal character
// as defined in https://infra.spec.whatwg.org/#code-point
func (s SelParser) isHexChar(char byte) bool {
	return 'A' <= char && char <= 'F' || 'a' <= char && char <= 'f' || '0' <= char && char <= '9'
}
