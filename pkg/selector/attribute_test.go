package selector

import (
	"reflect"
	"strings"
	"testing"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func TestAttributeParser_Parse(t *testing.T) {
	tests := []struct {
		name     string
		selector string
		want     Sel
	}{
		{
			name:     "parse a basic attribute sel",
			selector: `[key]`,
			want: &AttrSelector{
				key: "key",
				op:  "",
				val: "",
			},
		},
		{
			name:     "parse a basic attribute sel with =",
			selector: `[key=val]`,
			want: &AttrSelector{
				key: "key",
				op:  "=",
				val: "val",
			},
		},
		{
			name:     "parse a basic attribute sel with ~=",
			selector: `[key~="val"]`,
			want: &AttrSelector{
				key: "key",
				op:  "~=",
				val: "val",
			},
		},
		{
			name:     "parse a basic attribute sel with |=",
			selector: `[key|="val"]`,
			want: &AttrSelector{
				key: "key",
				op:  "|=",
				val: "val",
			},
		},
		{
			name:     "parse a basic attribute sel with |=",
			selector: `[key|=val]`,
			want: &AttrSelector{
				key: "key",
				op:  "|=",
				val: "val",
			},
		},
		{
			name:     "parse a basic attribute sel with ^=",
			selector: `[key^="val"]`,
			want: &AttrSelector{
				key: "key",
				op:  "^=",
				val: "val",
			},
		},
		{
			name:     "parse a basic attribute sel with $=",
			selector: `[key$="val"]`,
			want: &AttrSelector{
				key: "key",
				op:  "$=",
				val: "val",
			},
		},
		{
			name:     "parse a basic attribute sel with *=",
			selector: `[key*="val"]`,
			want: &AttrSelector{
				key: "key",
				op:  "*=",
				val: "val",
			},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t1 *testing.T) {
				t := NewAttributeParser(tt.selector)
				if got, _ := t.Parse(); !reflect.DeepEqual(got, tt.want) {
					t1.Errorf("Parse() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func TestAttributeParser_Match(t *testing.T) {
	tests := []struct {
		name string
		html string
		id   string
		want bool
	}{
		{
			name: "match element with attribute 'width' for <input width='100px'/>",
			html: `<input width="100px"/>`,
			id:   `[width]`,
			want: true,
		},
		{
			name: "match element with attribute 'width' with value '200px' for <input width='200px'/>",
			html: `<input width="200px"/>`,
			id:   `[width="200px"]`,
			want: true,
		},
		{
			name: "match element with attribute 'title' that contains 'substring' for <input title='un input con substring en el texto'/>",
			html: `<input title="un input con substring en el texto"/>`,
			id:   `[title~="substring"]`,
			want: true,
		},
		{
			name: "match element with attribute 'title' with value that start with 'es-' for <input title='es-ES'/>",
			html: `<input title="es-ES"/>`,
			id:   `[title|="es"]`,
			want: true,
		},
		{
			name: "match element with attribute 'title' with value 'es' for <input title='es'/>",
			html: `<input title="es"/>`,
			id:   `[title|="es"]`,
			want: true,
		},
		{
			name: "match element with attribute 'title' that begins with 'es' for <input title='essential'/>",
			html: `<input title="essential"/>`,
			id:   `[title^="es"]`,
			want: true,
		},
		{
			name: "match element with attribute 'title' that ends with 'ade' for <input title='esplanade'/>",
			html: `<input title="esplanade"/>`,
			id:   `[title$="ade"]`,
			want: true,
		},
		{
			name: "match element with attribute 'title' that contains 'ass' for <input title='class'/>",
			html: `<input title="class"/>`,
			id:   `[title*="class"]`,
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t1 *testing.T) {
			n, _ := html.ParseFragment(strings.NewReader(tt.html), &html.Node{
				Type:     html.ElementNode,
				DataAtom: atom.Body,
				Data:     "body",
			})
			t := NewAttributeParser(tt.id)
			got, _ := t.Parse()
			res := got.Match(n[0])

			if res != tt.want {
				t1.Errorf("Match() = %v, want %v for html '%s' with selector %s", res, tt.want, tt.html, tt.id)
				return
			}
		})
	}
}
