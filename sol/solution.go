package sol

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
				cur := append(cur, s[start:end+1])
				dfs(end+1, cur)
				cur = cur[:len(cur)-1]
			}
		}
	}
	dfs(0, []string{})
	return result
}
