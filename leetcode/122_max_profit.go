package leetcode

//变化： 可以买卖无限次，但同一时间最多持有一股（必须先卖再买）。
// 示例：
// prices = [7, 1, 5, 3, 6, 4]
// 操作：1买入 → 5卖出（赚4）→ 3买入 → 6卖出（赚3）
// 总利润 = 4 + 3 = 7

func maxProfit(prices []int) int {
	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}
	return profit
}

func maxProfit2(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}

	dp := make([][2]int, 0, len(prices))
	dp[0][0] = 0          // 不持有
	dp[0][1] = -prices[0] // 持有

	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i]) // 不动 或 卖出
		dp[i][1] = max(dp[i-1][1], dp[i][0]-prices[i])   // 不动 或 买入
	}
	return dp[n-1][0]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
