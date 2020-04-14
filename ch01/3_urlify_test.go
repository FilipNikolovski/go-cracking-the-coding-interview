package ch01

import "testing"

func TestURLify(t *testing.T) {
	tests := []struct {
		input string
		n     int
		out   string
	}{
		{"Mr John Smith", 13, "Mr%20John%20Smith"},
		{"a b c d efg", 11, "a%20b%20c%20d%20efg"},
	}

	for _, tt := range tests {
		u := URLify(tt.input, tt.n)
		if u != tt.out {
			t.Errorf("expected %s to be %s", u, tt.out)
		}
	}
}
