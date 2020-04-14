package ch01

import (
	"sort"
)

func HasUniqueChars(s string) bool {
	m := make(map[rune]bool)

	for _, r := range s {
		if _, ok := m[r]; ok {
			return false
		}

		m[r] = true
	}

	return true
}

func HasUniqueCharsv2(s string) bool {
	s = sortStr(s)

	for i, j := 0, 1; i < len(s)-1 && j < len(s); i, j = i+1, j+1 {
		if s[i] == s[j] {
			return false
		}
	}
	return true
}

func sortStr(s string) string {
	d := sortRunes(s)
	sort.Sort(d)
	return string(d)
}

type sortRunes []rune

func (sr sortRunes) Less(i, j int) bool {
	return sr[i] < sr[j]
}

func (sr sortRunes) Swap(i, j int) {
	sr[i], sr[j] = sr[j], sr[i]
}

func (sr sortRunes) Len() int {
	return len(sr)
}
