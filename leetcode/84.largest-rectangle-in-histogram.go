package leetcode

/*
 * @lc app=leetcode id=84 lang=golang
 *
 * [84] Largest Rectangle in Histogram
 *
 * https://leetcode.com/problems/largest-rectangle-in-histogram/description/
 *
 * algorithms
 * Hard (47.05%)
 * Likes:    18191
 * Dislikes: 324
 * Total Accepted:    1.2M
 * Total Submissions: 2.5M
 * Testcase Example:  '[2,1,5,6,2,3]'
 *
 * Given an array of integers heights representing the histogram's bar height
 * where the width of each bar is 1, return the area of the largest rectangle
 * in the histogram.
 *
 *
 * Example 1:
 *
 *
 * Input: heights = [2,1,5,6,2,3]
 * Output: 10
 * Explanation: The above is a histogram where width of each bar is 1.
 * The largest rectangle is shown in the red area, which has an area = 10
 * units.
 *
 *
 * Example 2:
 *
 *
 * Input: heights = [2,4]
 * Output: 4
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= heights.length <= 10^5
 * 0 <= heights[i] <= 10^4
 *
 *
 */

// @lc code=start
func largestRectangleArea(heights []int) int {
	maxArea := 0
	heights = append([]int{0}, heights...)
	heights = append(heights, 0)
	st := []int{}

	for i := 0; i < len(heights); i++ {
		for len(st) > 0 && heights[i] < heights[st[len(st)-1]] {
			h := heights[st[len(st)-1]]
			st = st[:len(st)-1]
			w := i - st[len(st)-1] - 1
			area := h * w
			if area > maxArea {
				maxArea = area
			}
		}
		st = append(st, i)
	}
	return maxArea
}

// @lc code=end

// [2,1,5,6,2,3]
// [0,2,1,5,6,2,3,0]
// st
// i0, i1
// i0, area=2
// i0, i2, i3,i4,i5
// i0, i2, i3,i4 , 5-3
// 4, 3-1
