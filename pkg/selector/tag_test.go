package selector

import (
	"reflect"
	"strings"
	"testing"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func TestTagParser_Parse(t1 *testing.T) {
	tests := []struct {
		name     string
		selector string
		want     Sel
	}{
		{
			name:     "parse a basic id",
			selector: "section",
			want:     &TagSelector{tag: "section"},
		},
		{
			name:     "parse a basic id with escaped letter",
			selector: `s\ection`,
			want:     &TagSelector{tag: "section"},
		},
		{
			name:     "parse a basic id with \\n",
			selector: "sect\nion",
			want:     &TagSelector{tag: "section"},
		},
		{
			name:     "parse a basic id with \\r",
			selector: "sect\rion",
			want:     &TagSelector{tag: "section"},
		},
		{
			name:     "parse a basic id with \\t",
			selector: "sect\tion",
			want:     &TagSelector{tag: "section"},
		},
		{
			name:     "parse a basic id with \\r\\n",
			selector: "sect\r\nion",
			want:     &TagSelector{tag: "section"},
		},
		{
			name:     "parse a basic id with whitespace",
			selector: "sect ion",
			want:     &TagSelector{tag: "section"},
		},
		{
			name:     "parse id with escaped element with 6 digits '\\000073' (s)",
			selector: `\000073ection`,
			want:     &TagSelector{tag: "section"},
		},
		{
			name:     "parse id with multiple escaped element with 6 digits '\\000073' (s) and '\\000069' (i)",
			selector: `\000073ect\000069on`,
			want:     &TagSelector{tag: "section"},
		},
		{
			name:     "parse id with escaped element with 2 digits and whitespace '\\73' (s)",
			selector: `\73 ection`,
			want:     &TagSelector{tag: "section"},
		},
		{
			name:     "parse id with multiple escaped element with 2 digits and whitespace '\\73' (s) and '\\69' (i)",
			selector: `\73 ection`,
			want:     &TagSelector{tag: "section"},
		},
		{
			name:     "parse id with escaped UNICODE element '\\U+000073' (s)",
			selector: `\U+000073ection`,
			want:     &TagSelector{tag: "section"},
		},
		{
			name:     "parse id with multiple escaped UNICODE element '\\U+000073' (s) and '\\U+000069' (i)",
			selector: `\U+000073ect\U+000069on`,
			want:     &TagSelector{tag: "section"},
		},
		{
			name:     "parse id with escaped UNICODE element with 4 digits and whitespace '\\U+0073' (s)",
			selector: `\U+0073 ection`,
			want:     &TagSelector{tag: "section"},
		},
		{
			name:     "parse id with multiple escaped UNICODE element with 4 digits and whitespace '\\U+0073' (s) and '\\U+0069' (i)",
			selector: `\U+0073 ect\U+0069 on`,
			want:     &TagSelector{tag: "section"},
		},
	}
	for _, tt := range tests {
		t1.Run(
			tt.name, func(t1 *testing.T) {
				t := NewTagParser(tt.selector)
				if got, _ := t.Parse(); !reflect.DeepEqual(got, tt.want) {
					t1.Errorf("Parse() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func TestTagParser_Match(t *testing.T) {
	tests := []struct {
		name string
		html string
		tag  string
		want bool
	}{
		{
			name: "match 'a' selector for <a>",
			html: "<a></a>",
			tag:  "a",
			want: true,
		},
		{
			name: "match 'main' selector for <main>",
			html: "<main></main>",
			tag:  "main",
			want: true,
		},
		{
			name: "not match 'li' selector for <ul>",
			html: "<ul></ul>",
			tag:  "li",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t1 *testing.T) {
			n, _ := html.ParseFragment(strings.NewReader(tt.html), &html.Node{
				Type:     html.ElementNode,
				DataAtom: atom.Body,
				Data:     "body",
			})
			t := NewTagParser(tt.tag)
			got, _ := t.Parse()
			res := got.Match(n[0])

			if res != tt.want {
				t1.Errorf("Match() = %v, want %v for html '%s' with selector %s", res, tt.want, tt.html, tt.tag)
				return
			}
		})
	}
}
