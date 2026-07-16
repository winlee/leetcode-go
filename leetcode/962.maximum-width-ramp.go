package leetcode

/*
 * @lc app=leetcode id=962 lang=golang
 *
 * [962] Maximum Width Ramp
 *
 * https://leetcode.com/problems/maximum-width-ramp/description/
 *
 * algorithms
 * Medium (55.68%)
 * Likes:    2736
 * Dislikes: 90
 * Total Accepted:    191.1K
 * Total Submissions: 343.1K
 * Testcase Example:  '[6,0,8,2,1,5]'
 *
 * A ramp in an integer array nums is a pair (i, j) for which i < j and nums[i]
 * <= nums[j]. The width of such a ramp is j - i.
 *
 * Given an integer array nums, return the maximum width of a ramp in nums. If
 * there is no ramp in nums, return 0.
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [6,0,8,2,1,5]
 * Output: 4
 * Explanation: The maximum width ramp is achieved at (i, j) = (1, 5): nums[1]
 * = 0 and nums[5] = 5.
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [9,8,1,0,1,9,4,0,4,1]
 * Output: 7
 * Explanation: The maximum width ramp is achieved at (i, j) = (2, 9): nums[2]
 * = 1 and nums[9] = 1.
 *
 *
 *
 * Constraints:
 *
 *
 * 2 <= nums.length <= 5 * 10^4
 * 0 <= nums[i] <= 5 * 10^4
 *
 *
 */

// @lc code=start
// func maxWidthRamp(nums []int) int {
// 	max := 0
// 	n := len(nums)
// 	for i := 0; i < n; i++ {
// 		for len(st) > 0 && nums[i] >= nums[st[len(st)-1]] {
// 			j := st[len(st)-1]
// 			if i-j > max {
// 				max = i - j
// 			}
// 			st = st[:len(st)-1]
// 		}

// 		st = append(st, i)
// 	}

// }

// @lc code=end

// [6,0,8,2,1,5]
// 6 st i0/6 ,i1/0
// 0 st i0/6 ,i1/0
// 8 st i0/6=2,i1/0=1 i3/8
// 2 st i0/6=2,i1/0=1 i3/8
