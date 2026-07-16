package leetcode

/*

题目： 给定不同面额的硬币和一个总金额，计算凑成总金额所需的最少硬币数量。如果无法凑出，返回 -1。

示例：

coins = [1, 2, 5], amount = 11
5 + 5 + 1 = 11，共 3 枚
答案 = 3



coins = [2], amount = 3
无法凑出 3
答案 = -1

*/ 

func coinChange(coins []int, amount int) int{
	dp := make([]int,amount+1)
	for i:=1; i<= amount;i++ {
		dp[i] = amount+1 // 表示无穷大
	}

	dp[0] = 0
	for _,coin := range coins {
		for j := coin; j<=amount;j++ {
			dp[j] = min(dp[j], dp[j-coin]+1)
		}
	}

	if dp[amount]> amount {
		return -1
	}

	return dp[amount]
}
