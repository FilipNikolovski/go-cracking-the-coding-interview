package ch01

import (
	"unicode"
)

func IsPalindromePerm(s string) bool {
	var letters int
	m := make(map[rune]int)

	for _, r := range s {
		if !unicode.IsSpace(r) {
			m[unicode.ToLower(r)]++
			letters++
		}
	}

	if letters%2 == 0 {
		// if there's even number of letters
		// we should have an even number of each letter
		// otherwise we cannot form a palindrome
		for _, v := range m {
			if v%2 != 0 {
				return false
			}
		}
	}

	if letters%2 == 1 {
		// for odd number of letters we should only have one odd letter
		// to put in the middle of the palindrome. The rest of the letters
		// should be of even number
		oneOdd := false
		for _, v := range m {
			if v == 1 {
				if oneOdd {
					return false
				}
				oneOdd = true
				continue
			}

			if v%2 != 0 {
				return false
			}
		}
	}
	return true
}
