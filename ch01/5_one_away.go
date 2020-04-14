package ch01

func OneAway(a, b string) bool {
	if a == b {
		return true
	}

	numEdits := -(len(a) - len(b))

	if numEdits > 1 {
		return false
	}

	if numEdits == 0 {
		edited := false
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				if edited {
					return false
				}
				edited = true
			}
		}

		return true
	}

	if len(a) > len(b) {
		return canInsert(a, b)
	}

	return canInsert(b, a)
}

func canInsert(a, b string) bool {
	for i, j := 0, 0; i < len(a) && j < len(b); {
		if a[i] != b[j] {
			if i != j {
				return false
			}
			i++
			continue
		}

		i++
		j++
	}
	return true
}
