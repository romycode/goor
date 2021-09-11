package selector

import (
	"fmt"

	"golang.org/x/net/html"
)

type UniversalSelector struct{}

func (t UniversalSelector) Match(_ *html.Node) bool {
	return true
}

type UniversalParser struct {
	SelParser
}

func NewUniversalParser(sel string) *UniversalParser {
	return &UniversalParser{
		SelParser{
			sel:    sel,
			selLen: len(sel),
		},
	}
}

func (t UniversalParser) Parse() (Sel, error) {
	if t.sel[t.pos] != '*' {
		return nil, fmt.Errorf("expected universal selector (*), found '%c'", t.sel[t.pos])
	}

	return &UniversalSelector{}, nil
}
