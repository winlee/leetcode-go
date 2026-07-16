package leetcode

func rob(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-1],dp[i-2]+nums[i])
	}
	return dp[len(nums)-1]
}


func rob2(nums []int) int {
	n := len(nums)
    if n == 0 {
        return 0
    }
    if n == 1 {
        return nums[0]
    }

	prev2, prev1 := nums[0], max(nums[0],nums[1])
	for i := 2; i < len(nums); i++ {
		curr := max(prev2+nums[i], prev1)
		prev2, prev1 = prev1, curr
	}
	
	return prev1
}