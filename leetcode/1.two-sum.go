/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
    m := make(map[int]int)
	for i,v := range nums {
		if j, ok := m[target-v];ok {
			return []int{j,i}
		}else {
			m[v] = i 
		}
	}
	return nil 
}
// @lc code=end

