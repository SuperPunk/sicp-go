package ch03

func Interleave(a, b []string) [][]string {
	if len(a) == 0 {
		return [][]string{b}
	}
	if len(b) == 0 {
		return [][]string{a}
	}

	var ret [][]string
	for _, v := range Interleave(a[1:], b) {
		ret = append(ret, append([]string{a[0]}, v...))
	}
	for _, v := range Interleave(a, b[1:]) {
		ret = append(ret, append([]string{b[0]}, v...))
	}
	return ret
}
