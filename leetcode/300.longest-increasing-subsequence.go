package leetcode

/*
 * @lc app=leetcode id=300 lang=golang
 *
 * [300] Longest Increasing Subsequence
 *
 * https://leetcode.com/problems/longest-increasing-subsequence/description/
 *
 * algorithms
 * Medium (57.53%)
 * Likes:    21734
 * Dislikes: 481
 * Total Accepted:    2.1M
 * Total Submissions: 3.7M
 * Testcase Example:  '[10,9,2,5,3,7,101,18]'
 *
 * Given an integer array nums, return the length of the longest strictly
 * increasing subsequence.
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [10,9,2,5,3,7,101,18]
 * Output: 4
 * Explanation: The longest increasing subsequence is [2,3,7,101], therefore
 * the length is 4.
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [0,1,0,3,2,3]
 * Output: 4
 *
 *
 * Example 3:
 *
 *
 * Input: nums = [7,7,7,7,7,7,7]
 * Output: 1
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 2500
 * -10^4 <= nums[i] <= 10^4
 *
 *
 *
 * Follow up: Can you come up with an algorithm that runs in O(n log(n)) time
 * complexity?
 *
 */

// @lc code=start
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	n := len(nums)
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}

	maxLen := 1
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		maxLen = max(maxLen, dp[i])
	}

	return maxLen

	// if len(nums) == 0  {
	// 	return 0
	// }

	// max := 0
	// n := len(nums)
	// dp := make([]int,n)
	// for i :range dp {
	// 	dp[i] = 1
	// }

	// maxLen := 1
	// for i := 1; i < n; i++ {
	// 	for j:=0; j<i; j++ {
	// 		if nums[i]>nums[j] {
	// 			dp[i] = max(dp[i],dp[j]+1)
	// 		}
	// 	}
	// }

	// return maxLen

}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// @lc code=end
