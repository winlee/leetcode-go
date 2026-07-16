package leetcode

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, v := range nums {
		sum+=v
	}

	if (sum+target)%2 == 1 || abs(target) > sum {
		return 0
	}

	p := (sum+target)/2
	dp := make([]int, p+1)
	dp[0]=1 // 凑出和为0，有1种方案（空集）

	for _, num := range nums {
		for j := p; j >= num; j-- {
			dp[j] += dp[j-num]
		}
	}

	return dp[p]
}

func abs(x int ) int {
	if x < 0 {
		return -x
	}
	return x
}