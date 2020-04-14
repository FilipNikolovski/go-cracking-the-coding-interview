package ch01

import "testing"

func TestHasUniqueChars(t *testing.T) {
	tests := []struct {
		input  string
		expect bool
	}{
		{"abcdef", true},
		{"abdeaf", false},
		{"dbefghabalaa", false},
	}

	for _, tt := range tests {
		u := HasUniqueChars(tt.input)
		if u != tt.expect {
			t.Errorf("HasUniqueChars(): expected %t got %t for input %s", tt.expect, u, tt.input)
		}
		u = HasUniqueCharsv2(tt.input)
		if u != tt.expect {
			t.Errorf("HasUniqueCharsv2(): expected %t got %t for input %s", tt.expect, u, tt.input)
		}
	}
}
