package leetcode

func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true 

	for i := 1; i <=n; n++ {
		for _, word := range wordDict {
			if len(word) <=i && s[i-len(word):i] == word && dp[i-len(word)]{
				dp[i] = true 
				break 
			}
		}
	}

	return dp[n]
}