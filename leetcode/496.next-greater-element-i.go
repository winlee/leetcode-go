/*
 * @lc app=leetcode id=496 lang=golang
 *
 * [496] Next Greater Element I
 *
 * https://leetcode.com/problems/next-greater-element-i/description/
 *
 * algorithms
 * Easy (74.26%)
 * Likes:    8801
 * Dislikes: 932
 * Total Accepted:    1M
 * Total Submissions: 1.4M
 * Testcase Example:  '[4,1,2]\n[1,3,4,2]'
 *
 * The next greater element of some element x in an array is the first greater
 * element that is to the right of x in the same array.
 * 
 * You are given two distinct 0-indexed integer arrays nums1 and nums2, where
 * nums1 is a subset of nums2.
 * 
 * For each 0 <= i < nums1.length, find the index j such that nums1[i] ==
 * nums2[j] and determine the next greater element of nums2[j] in nums2. If
 * there is no next greater element, then the answer for this query is -1.
 * 
 * Return an array ans of length nums1.length such that ans[i] is the next
 * greater element as described above.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: nums1 = [4,1,2], nums2 = [1,3,4,2]
 * Output: [-1,3,-1]
 * Explanation: The next greater element for each value of nums1 is as follows:
 * - 4 is underlined in nums2 = [1,3,4,2]. There is no next greater element, so
 * the answer is -1.
 * - 1 is underlined in nums2 = [1,3,4,2]. The next greater element is 3.
 * - 2 is underlined in nums2 = [1,3,4,2]. There is no next greater element, so
 * the answer is -1.
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: nums1 = [2,4], nums2 = [1,2,3,4]
 * Output: [3,-1]
 * Explanation: The next greater element for each value of nums1 is as follows:
 * - 2 is underlined in nums2 = [1,2,3,4]. The next greater element is 3.
 * - 4 is underlined in nums2 = [1,2,3,4]. There is no next greater element, so
 * the answer is -1.
 * 
 * 
 * 
 * Constraints:
 * 
 * 
 * 1 <= nums1.length <= nums2.length <= 1000
 * 0 <= nums1[i], nums2[i] <= 10^4
 * All integers in nums1 and nums2 are unique.
 * All the integers of nums1 also appear in nums2.
 * 
 * 
 * 
 * Follow up: Could you find an O(nums1.length + nums2.length) solution?
 */

// @lc code=start
func nextGreaterElement(nums1 []int, nums2 []int) []int {

	ans := make([]int,len(nums1))
	mp := make(map[int]int)
	st := []int {}
	for i := len(nums2)-1; i>=0; i-- {
		num := nums2[i]
		mp[num] = -1 
		for len(st)>0 && num >= st[len(st)-1] {
			st = st[:len(st)-1]
		}

		if len(st) > 0 {
			mp[num] = st[len(st)-1]
		}
		
		st = append(st, num )
	}
	// 1,3,4,2,5
    // stack
	// 5
	// 5,2 
	// 5,4,3


	// for j:=0; j < len(nums2); j++ {
	// 	num := nums2[j]
	// 	mp[num] = -1 
	// 	for len(st)>0 && nums2[j] > nums2[st[len(st)-1]] {
	
	// 		mp[nums2[st[len(st)-1]]] = num
	// 		st = st[:len(st)-1]
	// 	}

	// 	st = append(st, j) // 递增栈
	// }
	
	for i, val:= range nums1 {
		ans[i] = mp[val]
	}
	return ans
}
// @lc code=end

