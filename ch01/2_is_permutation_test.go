package ch01

import "testing"

func TestIsPermutation(t *testing.T) {
	tests := []struct {
		a      string
		b      string
		expect bool
	}{
		{"abcdef", "fedcba", true},
		{"abdeaf", "aaab", false},
		{"aaab", "abaa", true},
	}

	for _, tt := range tests {
		res := IsPermutation(tt.a, tt.b)
		if res != tt.expect {
			t.Errorf("expected %t to be %t for inputs %s and %s", res, tt.expect, tt.a, tt.b)
		}

		res = IsPermutationv2(tt.a, tt.b)
		if res != tt.expect {
			t.Errorf("expected %t to be %t for inputs %s and %s", res, tt.expect, tt.a, tt.b)
		}
	}
}
