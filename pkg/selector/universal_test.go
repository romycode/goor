package selector

import (
	"reflect"
	"testing"
)

func TestUniversalParser_Parse(t1 *testing.T) {
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
		t1.Run(tt.name, func(t1 *testing.T) {
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
