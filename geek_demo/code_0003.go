package geek_demo

func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	m := map[byte]int{}
	maxLen, start := 0, 0
	for end := 0; end < len(s); end++ {
		if v, ok := m[s[end]]; ok {
			start = max(v+1, start)
		}

		maxLen = max(maxLen, end-start+1)
		m[s[end]] = end
	}

	return maxLen
}

//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
