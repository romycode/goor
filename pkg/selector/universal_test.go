package selector

import (
	"reflect"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestUniversalParser_Parse(t *testing.T) {
	tests := []struct {
		name     string
		selector string
		want     Sel
		wantErr  bool
	}{
		{name: "parse a good universal selector", selector: "*", want: &UniversalSelector{}, wantErr: false},
		{name: "throw error for bad universal selector", selector: "a", want: nil, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t1 *testing.T) {
			t := NewUniversalParser(tt.selector)
			got, err := t.Parse()
			if (err != nil) != tt.wantErr {
				t1.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("Parse() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUniversalParser_Match(t *testing.T) {
	tests := []struct {
		name string
		html string
		want bool
	}{
		{
			name: "parse a good universal selector",
			html: "<html><body><h1></h1></body></html>",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t1 *testing.T) {
			n, _ := html.Parse(strings.NewReader(tt.html))
			t := NewUniversalParser("*")
			got, _ := t.Parse()

			if res := got.Match(n); res != tt.want {
				t1.Errorf("Match() = %v, want %v", res, tt.want)
				return
			}
		})
	}
}
