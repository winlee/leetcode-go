package leetcode

/*
LeetCode 34. 在排序数组中查找元素的第一个和最后一个位置
链接： https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/

题目描述：
给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。

如果数组中不存在目标值 target，返回 [-1, -1]。

要求： 时间复杂度 O(log n)
核心思路： 两次二分查找
第一次找左边界（第一个 ≥ target 的位置）
第二次找右边界（最后一个 ≤ target 的位置）
这道题是二分查找的经典变体，考察对二分边界条件的精确控制。

难点在边界处理

左边界：
`
  从右向左
 if nums[mid]>=target {
   right = mid 
 }else{
   left = mid+1
 }
`
*/



func searchRange(nums []int, target int) []int {
    left := findFirst(nums, target)
    if left == len(nums) || nums[left] != target {
        return []int{-1, -1}
    }
    right := findFirst(nums, target+1) - 1 // 核心：找 target+1 的左边界再 -1
    return []int{left, right}
}

// 统一模板：找第一个 >= target 的位置
func findFirst(nums []int, target int) int {
    left, right := 0, len(nums)
    for left < right {
        mid := left + (right-left)/2
        if nums[mid] >= target {
            right = mid
        } else {
            left = mid + 1
        }
    }
    return left
}