package geek_demo

import (
	"math"
	"sort"
)

//322. 零钱兑换
//给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。
//
//你可以认为每种硬币的数量是无限的。
//
//示例 1：
//
//输入：coins = [1, 2, 5], amount = 11
//输出：3
//解释：11 = 5 + 5 + 1
//示例 2：
//
//输入：coins = [2], amount = 3
//输出：-1
//示例 3：
//
//输入：coins = [1], amount = 0
//输出：0
//示例 4：
//
//输入：coins = [1], amount = 1
//输出：1
//示例 5：
//
//输入：coins = [1], amount = 2
//输出：2
//
//
//提示：
//
//1 <= coins.length <= 12
//1 <= coins[i] <= 231 - 1
//0 <= amount <= 104

func coinChange(coins []int, amount int) int {
	count := math.MaxInt32
	if len(coins) == 0 || amount == 0 {
		return 0
	}

	sort.Ints(coins)
	var dfs func(compose []int, amt int)
	dfs = func(compose []int, amt int) {
		if amt == 0 {
			count = min(count, len(compose))
			return
		}

		for i := len(coins) - 1; i >= 0; i-- {
			v := coins[i]
			if amt >= v {
				compose = append(compose, v)
				dfs(compose, amt-v)
				compose = compose[:len(compose)-1]
			}
		}
	}

	dfs([]int{}, amount)
	if count == math.MaxInt32 {
		return -1
	}

	return count
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func getSum(compose []int) int {
	sum := 0
	for v := range compose {
		sum += v
	}
	return sum
}

// coinChange2 dfs+剪枝 比dp好
func coinChange2(coins []int, amount int) int {
	sort.Ints(coins)
	recursion(coins, amount, 0, len(coins)-1)

	if minCount == math.MaxInt32 {
		return -1
	}
	return minCount
}

var (
	minCount = math.MaxInt32
)

/**
 * 1、按金额从大到小，从多到少（排序，用余数一步到位）
 * 2、预判低于最优解，终止递归（可以返回最优解，不过提升有限，意义不大）
 * 3、能整除即可返回
 * @道法自然 dfs+剪枝 比dp好
 */
func recursion(coins []int, amount int, count int, index int) {
	if index < 0 || (count+amount/coins[index]) > minCount {
		return
	}

	if amount%coins[index] == 0 {
		minCount = min(minCount, count+amount/coins[index])
		return
	}

	for i := amount / coins[index]; i >= 0; i++ {
		recursion(coins, amount-i*coins[index], count+i, index-1)
	}
}

func coinChangeDp(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	n := len(coins)
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i < amount+1; i++ {
		dp[i] = math.MaxInt32
		for j := 0; j < n; j++ {
			if i == coins[j] {
				dp[i] = 1
			} else if i > coins[j] {
				//总金额>面额 个数=min【dp（总金额-面额）+1,dp[总金额]】
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}

	if dp[amount] == math.MaxInt32 {
		return -1
	}

	return dp[amount]
}
