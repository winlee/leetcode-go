package leetcode

func climbStairs(n int) int {
	if n <= 1 {
		return 1
	}

	dp := make([]int, n+1)
	dp[0] = 1
	stairs := []int{1, 2}
	for i := 1; i <= n; i++ {
		for _, stair := range stairs {
			if i >= stair {
				dp[i] += dp[i-stair]
			}
		}
	}

	return dp[n]
}
