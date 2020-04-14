package ch01

import "testing"

func TestOneAway(t *testing.T) {
	tests := []struct {
		a      string
		b      string
		expect bool
	}{
		{"pale", "ple", true},
		{"pales", "pale", true},
		{"pale", "bale", true},
		{"pale", "bake", false},
	}

	for _, tt := range tests {
		res := OneAway(tt.a, tt.b)
		if res != tt.expect {
			t.Errorf("expected %t got %t for inputs %s, %s", tt.expect, res, tt.a, tt.b)
		}
	}
}
