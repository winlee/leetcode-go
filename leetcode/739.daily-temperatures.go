/*
 * @lc app=leetcode id=739 lang=golang
 *
 * [739] Daily Temperatures
 *
 * https://leetcode.com/problems/daily-temperatures/description/
 *
 * algorithms
 * Medium (66.86%)
 * Likes:    13803
 * Dislikes: 349
 * Total Accepted:    1.3M
 * Total Submissions: 1.9M
 * Testcase Example:  '[73,74,75,71,69,72,76,73]'
 *
 * Given an array of integers temperatures represents the daily temperatures,
 * return an array answer such that answer[i] is the number of days you have to
 * wait after the i^th day to get a warmer temperature. If there is no future
 * day for which this is possible, keep answer[i] == 0 instead.
 * 
 * 
 * Example 1:
 * Input: temperatures = [73,74,75,71,69,72,76,73]
 * Output: [1,1,4,2,1,1,0,0]
 * Example 2:
 * Input: temperatures = [30,40,50,60]
 * Output: [1,1,1,0]
 * Example 3:
 * Input: temperatures = [30,60,90]
 * Output: [1,1,0]
 * 
 * 
 * Constraints:
 * 
 * 
 * 1 <= temperatures.length <= 10^5
 * 30 <= temperatures[i] <= 100
 * 
 * 
 */

// @lc code=start
func dailyTemperatures(temperatures []int) []int {
    n := len(temperatures)
	answer := make([]int,n)
	st := []int {}

	// 遍历
	for j := 0; j < n; j++ {
		// 如果 st 有数据，且比栈顶大，那就出栈，然后再入栈（）
		for len(st) > 0 && temperatures[j] > temperatures[st[len(st)-1]] {
			i := st[len(st)-1]
			answer[i] = j - i 
			st  = st[:len(st)-1] 
		}

		st = append(st, j)
	}
// 73, 
// 74, 75,
// 75,
// 75,71,
// 75,71,69,

	return answer
}
// @lc code=end

