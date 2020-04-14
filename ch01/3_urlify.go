package ch01

import (
	"unicode"
)

func URLify(s string, n int) string {
	spaces := 0
	for i := 0; i < n; i++ {
		if unicode.IsSpace(rune(s[i])) {
			spaces++
		}
	}

	len := n + spaces*2
	runes := make([]rune, len)
	index := len

	for i := n - 1; i >= 0; i-- {
		if unicode.IsSpace(rune(s[i])) {
			runes[index-1] = '0'
			runes[index-2] = '2'
			runes[index-3] = '%'
			index -= 3
			continue
		}

		runes[index-1] = rune(s[i])
		index--
	}

	return string(runes)
}
