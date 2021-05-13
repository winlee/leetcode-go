package geek_demo

import "sort"

func threeSum2(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		remain := -1 * nums[i]
		r := len(nums) - 1
		for l := i + 1; l < r; l++ {
			if l > i+1 && nums[l] == nums[l-1] {
				continue
			}

			for l < r && nums[l]+nums[r] > remain {
				r--
			}

			if l == r {
				break
			}

			if nums[l]+nums[r] == remain {
				res = append(res, []int{nums[i], nums[l], nums[r]})
			}
		}
	}

	return res
}
