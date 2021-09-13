package selector

import (
	"reflect"
	"strings"
	"testing"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func TestClassParser_Parse(t *testing.T) {
	tests := []struct {
		name     string
		selector string
		want     Sel
	}{
		{
			name:     "parse a basic key sel",
			selector: `.testid`,
			want:     &ClassSelector{class: "testid"},
		},
		{
			name:     "parse key sel with escaped element with 6 digits '\\000069' (s)",
			selector: `.test\000069d`,
			want:     &ClassSelector{class: "testid"},
		},
		{
			name:     "parse key sel with multiple escaped element with 6 digits '\\\000065' (s) and '\\000069' (i)",
			selector: `.t\000065st\000069d`,
			want:     &ClassSelector{class: "testid"},
		},
		{
			name:     "parse key sel with escaped element with 2 digits and whitespace '\\69' (s)",
			selector: `.test\69 d`,
			want:     &ClassSelector{class: "testid"},
		},
		{
			name:     "parse key sel with multiple escaped element with 2 digits and whitespace '\\65' (s) and '\\69' (i)",
			selector: `.t\65 st\69 d`,
			want:     &ClassSelector{class: "testid"},
		},
		{
			name:     "parse key sel with escaped UNICODE element '\\U+000069' (s)",
			selector: `.test\U+000069d`,
			want:     &ClassSelector{class: "testid"},
		},
		{
			name:     "parse key sel with multiple escaped UNICODE element '\\U+000065' (s) and '\\U+000069' (i)",
			selector: `.t\U+000065st\U+000069d`,
			want:     &ClassSelector{class: "testid"},
		},
		{
			name:     "parse key sel with escaped UNICODE element with 4 digits and whitespace '\\U+0073' (s)",
			selector: `.test\U+0069 d`,
			want:     &ClassSelector{class: "testid"},
		},
		{
			name:     "parse key sel with multiple escaped UNICODE element with 4 digits and whitespace '\\U+0073' (s) and '\\U+0069' (i)",
			selector: `.t\U+0065 st\U+0069 d`,
			want:     &ClassSelector{class: "testid"},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t1 *testing.T) {
				t := NewClassParser(tt.selector)
				if got, _ := t.Parse(); !reflect.DeepEqual(got, tt.want) {
					t1.Errorf("Parse() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func TestClassParser_Match(t *testing.T) {
	tests := []struct {
		name string
		html string
		id   string
		want bool
	}{
		{
			name: "match element with class 'match-class' for <input class='match-class'/>",
			html: `<input class='match-class'/>`,
			id:   ".match-class",
			want: true,
		},
		{
			name: "not match element with class 'match-class' for <input class='no-match-class'/>",
			html: `<input class='no-match-class'/>`,
			id:   ".match-class",
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
			t := NewClassParser(tt.id)
			got, _ := t.Parse()
			res := got.Match(n[0])

			if res != tt.want {
				t1.Errorf("Match() = %v, want %v for html '%s' with selector %s", res, tt.want, tt.html, tt.id)
				return
			}
		})
	}
}
