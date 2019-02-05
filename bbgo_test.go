package bbgo_test

import (
	"testing"

	"github.com/namreg/bbgo"
)

func TestBBGO_Parse(t *testing.T) {

	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"b", `[b][b]hello[/b]`, `<strong><strong>hello</strong>`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := bbgo.New()
			got := b.Parse(tt.input)
			if tt.want != got {
				t.Fatalf("want = %s, got = %s", tt.want, got)
			}
		})
	}
}
