package leetcode

import (
	"math"
	"sort"
)

/*
描述
给你一个长度为 n 的整数数组 nums 和一个目标值 target。请在 nums 中找出和最接近 target 的三个数，返回它们的和。
示例
输入: nums = [-1,2,1,-4], target = 1
输出: 2  (-1 + 2 + 1 = 2)
*/

func threeSumClosest_FAIL(nums []int, target int) int{
	
	sort.Slice(nums,func(i,j int)  bool{
		return nums[i]<nums[j]
	})

	left, right := 0,len(nums)-1

	delta := math.MaxInt32
	res:=make([]int,3)
	for left < right {
		for j :=left+1; j<right;j++ {
			d := nums[left]+nums[right]+nums[j] -target
			if abs(d) <  delta {
				delta = abs(d)
				res[0]=left
				res[1]=right
				res[2]=j
			} 
		}
	}

	return nums[res[0]]+nums[res[1]]+nums[res[2]]
}

func threeSumClosest(nums []int, target int) int {
    sort.Ints(nums)
    closest := nums[0] + nums[1] + nums[2]
    
    for i := 0; i < len(nums)-2; i++ {
        left, right := i+1, len(nums)-1
        for left < right {
            sum := nums[i] + nums[left] + nums[right]
            if abs(sum-target) < abs(closest-target) {
                closest = sum
            }
            if sum < target {
                left++
            } else if sum > target {
                right--
            } else {
                return target // 刚好等于，直接返回
            }
        }
    }
    return closest
}

