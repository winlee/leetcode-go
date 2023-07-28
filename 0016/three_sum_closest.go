package _016

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	minDiff := math.MaxInt
	ans := 0
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		j, k := i+1, len(nums)-1
		for j < k {
			s := nums[i] + nums[j] + nums[k]
			if s == target {
				return target
			}

			if s > target {
				if s-target < minDiff {
					minDiff = s - target
					ans = s
				}
				k--
			} else {
				if target-s < minDiff {
					minDiff = target - s
					ans = s
				}
				j++
			}
		}

	}
	return ans
}

func threeSumClosest2(nums []int, target int) (ans int) {
	sort.Ints(nums)
	n := len(nums)
	minDiff := math.MaxInt
	for i, x := range nums[:n-2] {
		if i > 0 && x == nums[i-1] {
			continue // 优化三
		}

		// 优化一
		s := x + nums[i+1] + nums[i+2]
		if s > target { // 后面无论怎么选，选出的三个数的和不会比 s 还小
			if s-target < minDiff {
				ans = s // 由于下面直接 break，这里无需更新 minDiff
			}
			break
		}

		// 优化二
		s = x + nums[n-2] + nums[n-1]
		if s < target { // x 加上后面任意两个数都不超过 s，所以下面的双指针就不需要跑了
			if target-s < minDiff {
				minDiff = target - s
				ans = s
			}
			continue
		}

		// 双指针
		j, k := i+1, n-1
		for j < k {
			s = x + nums[j] + nums[k]
			if s == target {
				return target
			}
			if s > target {
				if s-target < minDiff { // s 与 target 更近
					minDiff = s - target
					ans = s
				}
				k--
			} else { // s < target
				if target-s < minDiff { // s 与 target 更近
					minDiff = target - s
					ans = s
				}
				j++
			}
		}
	}
	return ans
}

func threeSumClosest3(nums []int, target int) int {
	n, res, diff := len(nums), 0, math.MaxInt32
	if n <= 2 {
		return res
	}

	sort.Ints(nums)
	for i := 0; i < n-2; i++ {
		for j, k := i+1, n-1; j < k; {
			sum := nums[i] + nums[j] + nums[k]
			if abs(sum-target) < diff {
				res, diff = sum, abs(sum-target)
			}
			if sum == target {
				return res
			} else if sum > target {
				k--
			} else {
				j++
			}
		}
	}

	return res
}

func abs(a int) int {
	if a > 0 {
		return a
	}

	return -a
}
