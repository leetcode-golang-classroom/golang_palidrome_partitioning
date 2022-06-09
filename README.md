# golang_palindrome_partitioning

Given a string `s`, partition `s` such that every substring of the partition is a **palindrome**. Return all possible palindrome partitioning of `s`.

A **palindrome** string is a string that reads the same backward as forward.

## Examples

**Example 1:**

```
Input: s = "aab"
Output: [["a","a","b"],["aa","b"]]

```

**Example 2:**

```
Input: s = "a"
Output: [["a"]]

```

**Constraints:**

- `1 <= s.length <= 16`
- `s` contains only lowercase English letters.

## 解析

題目給定一個 字串 s

要求實作一個演算法找出所有拆分 s 成的回文子字串陣列 result

從要求可以發現幾個限制

1. 所有在 result的元素 都是一個 字串陣列， 且字串陣列做串接會是 s
2. 每個字串陣列內的字串必須是回文

要窮舉的作法

可以從 index = 0 為起點找出該起點所有可能的長度字串

當 search 的 index = s 的長度 代表已經找完

如下圖：

![](https://i.imgur.com/N7zjQYl.png)

而判斷回文的作法可以透過

![](https://i.imgur.com/xQi2PxF.png)

## 程式碼
```go
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

```
## 困難點

1. 要想出如何窮舉出所有可能的方法
2. 要想出如何透過已找過的結果來加速判斷字串是否為迴文

## Solve Point

- [x]  Understand what problem to solve
- [x]  Analysis Complexity