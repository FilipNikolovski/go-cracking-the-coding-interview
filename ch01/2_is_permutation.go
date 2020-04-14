package ch01

func IsPermutation(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	s1 := sortStr(a)
	s2 := sortStr(b)

	for i := 0; i < len(a); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}

func IsPermutationv2(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	m := make(map[rune]int)

	for _, r := range a {
		m[r]++
	}

	for _, r := range b {
		m[r]--
		if m[r] < 0 {
			return false
		}
	}

	return true
}
