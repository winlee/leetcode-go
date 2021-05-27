package geek_demo

import (
	"math"
	"sort"
)

func permute(nums []int) [][]int {
	var res [][]int
	visited := map[int]bool{}

	var dfs func(path []int)
	dfs = func(path []int) {
		if len(path) == len(nums) {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}

		for _, n := range nums {
			if visited[n] {
				continue
			}

			path = append(path, n)
			visited[n] = true
			dfs(path)
			path = path[:len(path)-1]
			visited[n] = false
		}
	}

	dfs([]int{})
	return res
}

func permute2(nums []int) [][]int {
	var res [][]int
	visited := map[int]bool{}
	var dfs func(path []int)
	dfs = func(path []int) {
		if len(path) == len(nums) {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}

		for _, n := range nums {
			if visited[n] {
				continue
			}

			path = append(path, n)
			visited[n] = true
			dfs(path)
			path = path[:len(path)-1]
			visited[n] = false
		}
	}

	dfs([]int{})
	return res
}

// tmpCoinChangeDp
// dp[i] =
//	i == coins[j] dp[i] = 1
//  i != coins[j] dp[i] = min(dp[i], dp[i-coins[j]]+1)
func tmpCoinChangeDp(coins []int, amount int) int {
	if len(coins) == 0 || amount == 0 {
		return 0
	}

	//sort.Ints(coins)
	dp := make([]int, amount+1)
	for i := 1; i < amount; i++ {
		dp[i] = math.MaxInt32
		for j := 0; j < len(coins); j++ {
			if i == coins[j] {
				dp[i] = 1
			} else {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}
	return dp[amount]
}

var tmpMinCount = math.MaxInt32

// 深度遍历
func tmpCoinChangeRecursion(coins []int, amount int) int {
	sort.Ints(coins)
	tmpRecursion(coins, amount, 0, len(coins)-1)
	if tmpMinCount == math.MaxInt32 {
		return -1
	}
	return tmpMinCount
}

func tmpRecursion(coins []int, amount int, count int, index int) {
	if index < 0 || (count+amount/coins[index]) > tmpMinCount {
		return
	}

	if amount%coins[index] == 0 {
		tmpMinCount = min(minCount, count+amount/coins[index])
		return
	}

	for i := amount / coins[index]; i >= 0; i++ {
		tmpRecursion(coins, amount-i*coins[index], count+i, index-1)
	}
}

func coinChange3(coins []int, amount int) int {
	count := math.MaxInt32
	if amount == 0 {
		return 0
	}

	var dfs func(compose []int, amt int)
	dfs = func(compose []int, amt int) {
		if amt == 0 {
			count = min(count, len(compose))
			return
		}

		for i := len(coins) - 1; i >= 0; i-- {
			v := coins[i]
			div := amt / v
			if div >= 1 {
				for j := 0; j < div; j++ {
					compose = append(compose, v)
				}
				dfs(compose, amt-div*v)
				compose = compose[:len(compose)-v]
			} else {
				dfs(compose, amt)
			}
		}
	}
	dfs([]int{}, amount)
	if count == math.MaxInt32 {
		return -1
	}

	return count

}
