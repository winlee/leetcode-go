/*
 * @lc app=leetcode id=503 lang=golang
 *
 * [503] Next Greater Element II
 *
 * https://leetcode.com/problems/next-greater-element-ii/description/
 *
 * algorithms
 * Medium (65.97%)
 * Likes:    8478
 * Dislikes: 209
 * Total Accepted:    557.7K
 * Total Submissions: 845.4K
 * Testcase Example:  '[1,2,1]'
 *
 * Given a circular integer array nums (i.e., the next element of
 * nums[nums.length - 1] is nums[0]), return the next greater number for every
 * element in nums.
 * 
 * The next greater number of a number x is the first greater number to its
 * traversing-order next in the array, which means you could search circularly
 * to find its next greater number. If it doesn't exist, return -1 for this
 * number.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: nums = [1,2,1]
 * Output: [2,-1,2]
 * Explanation: The first 1's next greater number is 2; 
 * The number 2 can't find next greater number. 
 * The second 1's next greater number needs to search circularly, which is also
 * 2.
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: nums = [1,2,3,4,3]
 * Output: [2,3,4,-1,4]
 * 
 * 
 * 
 * Constraints:
 * 
 * 
 * 1 <= nums.length <= 10^4
 * -10^9 <= nums[i] <= 10^9
 * 
 * 
 */

// @lc code=start
func nextGreaterElements(nums []int) []int {
    n := len(nums)
	ans := make([]int,len(nums))
	st := []int{}

	for i := 0; i < n; i++ {
		ans[i] = -1
		for len(st)> 0 && nums[i]>nums[st[len(st)-1]] {
			ans[st[len(st)-1]]	= nums[i]
			st = st[:len(st)-1]
		}

		st = append(st, i)
	}

	for len(st)> 0 {
		last := st[len(st)-1]
		for _, v := range nums {
			if v > nums[last] {
				ans[last] = v 
				break
			}
		}
		st = st[:len(st)-1]
	}

	return ans
}
// @lc code=end

// 5,4,3,2,1
// st 
// i1,i2,i3,i4,i5
// i5=5

// [1,2,3,4,3]
// [2,3,4,-1,4]
// st 0  if 2 > 1 { ans[0]=2}
// st i3,i4  

// for len(st)>0 {
//     
// }