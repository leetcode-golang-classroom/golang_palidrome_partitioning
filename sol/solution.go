package sol

import "unsafe"

func Clone(s string) string {
	b := make([]byte, len(s))
	copy(b, s)
	return *(*string)(unsafe.Pointer(&b))
}
func partition(s string) [][]string {
	sLen := len(s)
	// check
	check := make([][]bool, sLen)
	for i := range check {
		check[i] = make([]bool, sLen)
	}
	result := [][]string{}
	var dfs func(start int, cur []string)
	dfs = func(start int, cur []string) {
		if start >= sLen {
			temp := make([]string, len(cur))
			copy(temp, cur)
			result = append(result, temp)
			return
		}
		// check s[start: end+1]
		for end := start; end < sLen; end++ {
			if s[start] == s[end] && (end-start <= 2 || check[start+1][end-1]) {
				check[start][end] = true
				cur := append(cur, Clone(s[start:end+1]))
				dfs(end+1, cur)
				cur = cur[:len(cur)-1]
			}
		}
	}
	dfs(0, []string{})
	return result
}
