package ch01

import "testing"

func TestPalindromePerm(t *testing.T) {
	tests := []struct {
		input  string
		expect bool
	}{
		{"Tact Coa", true},
		{"Abcd", false},
		{"Aabb", true},
	}

	for _, tt := range tests {
		res := IsPalindromePerm(tt.input)
		if res != tt.expect {
			t.Errorf("expected %t got %t for input %s", tt.expect, res, tt.input)
		}
	}
}
