package selector

import (
	"reflect"
	"testing"
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
			selector: `[key="val"]`,
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
