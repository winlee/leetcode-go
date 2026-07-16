package leetcode

// 求数组内的等和子集，说明数组总和可以分成两份
func canPartition(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	if sum%2 == 1 {
		return false
	}

	target := sum / 2

	// 一维数组
	dp := make([]bool, target+1)
	dp[0] = true // 和为0，不选任何元素

	for _, num := range nums {
		// 倒序遍历！
		for j := target; j >= num; j-- {
			dp[j] = dp[j] || dp[j-num]
		}
	}

	return dp[target]
}
