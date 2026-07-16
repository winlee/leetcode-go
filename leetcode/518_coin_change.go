package leetcode

/**
题目： 给定不同面额的硬币和一个总金额，计算凑成总金额的组合数（每种硬币可以无限次使用）。

示例：

amount = 5, coins = [1, 2, 5]
5 = 5
5 = 2 + 2 + 1
5 = 2 + 1 + 1 + 1
5 = 1 + 1 + 1 + 1 + 1
共 4 种组合
答案 = 4



提示：
这是什么类型的背包问题？
对比 494，dp[j] 表示什么？
组合数 vs 排列数，遍历顺序有什么不同？
*/

func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1 // 面额0，1种方法

	// 从物品的角度看，
	// 先看物品1如何放满
	// 再看物品2如何放满
	// 再看物品n如何放满
	for _, coin := range coins {
		for j := coin; j <= amount; j++ {
			dp[j] += dp[j-coin]
		}
	}

	return dp[amount]
}
