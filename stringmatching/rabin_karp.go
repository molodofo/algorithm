package stringmatching

var PRIME = 101
var BASE = 256

func hash(s string) int {
	h := 0
	for _, c := range s {
		h = hashLengthen(h, int(c)) // because BASE is 256, it can change to ((h + int(c)) << 8) % PRIME
	}
	return h
}

func hashLengthen(h int, new int) int {
	return ((h * BASE) + new) % PRIME
}

func hashShorten(h, old int) int {
	return h + PRIME - (old*((BASE%PRIME)*BASE))%PRIME
}

func RabinKarp(source, patten string) int {
	m, n := len(source), len(patten)
	if m < n {
		return -1
	}

	ph := hash(patten)
	sh := hash(source[0:n])
	if ph == sh && patten == source[0:n] {
		return 0
	}

	for i := 1; i < m-n+1; i++ {
		sh = hashLengthen(hashShorten(sh, int(source[i-1])), int(source[n+i-1]))
		if ph == sh && patten == source[i:n+i] {
			return i
		}
	}

	return -1
}

func RabinKarpMulti(source string, pattenLen int, pattens ...string) int {
	pSet := make(map[string]bool)
	phSet := make(map[int]bool)
	for _, p := range pattens {
		pSet[p] = true
		phSet[hash(p)] = true
	}

	sh := hash(source[0:pattenLen])
	i := 0
	check := func() int {
		if _, ok := phSet[sh]; ok {
			if _, ok := pSet[source[i:pattenLen+i]]; ok {
				return i
			}
		}
		return -1
	}

	if check() == i {
		return i
	}

	i++
	for ; i < len(source)-pattenLen+1; i++ {
		sh = hashLengthen(hashShorten(sh, int(source[i-1])), int(source[pattenLen+i-1]))
		if check() == i {
			return i
		}
	}

	return -1
}
