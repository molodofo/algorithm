package stringmatching

func getNext(t string) []int {
	l := len(t)
	next := make([]int, l)

	next[0] = -1
	i, j := 0, -1
	for i < (l - 1) {
		if j == -1 || t[i] == t[j] {
			// if j == -1 then that mean the prefix and suffix of current substring are not equal
			// if t[i] == t[j] then that mean the last char of substring prefix is equal the last char of suffix
			i++
			j++
			next[i] = j
		} else {
			j = next[j]
		}
	}
	return next
}

func getNewNext(t string) []int {
	l := len(t)
	next := make([]int, l)

	next[0] = -1
	i, j := 0, -1
	for i < (l - 1) {
		if j == -1 || t[i] == t[j] {
			i++
			j++
			if t[i] == t[j] {
				next[i] = next[j]
			} else {
				next[i] = j
			}
		} else {
			j = next[j]
		}
	}
	return next
}

func KMP(p, t string) int {
	m, n := len(p), len(t)
	if m == 0 || m < n {
		return -1
	}

	next := getNewNext(t)
	i, j := 0, 0
	for i < m && j < n {
		if j == -1 || p[i] == t[j] {
			i++
			j++
		} else {
			j = next[j]
		}
	}

	if j == n {
		return i - j
	}
	return -1
}
