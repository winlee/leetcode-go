package geek_demo

//描述：(字节一面面试题)
//递增数组中目标值出现的次数，复杂度<O(n)
//输入：1, 2, 3, 3, 4, 4, 4, 4, 4, 5, 5, 8, 10, 12 和 target = 4
//输出：target出现的次数，5次
//
//findRepeatTarget 统计重复数字个数
func findRepeatTarget(arr []int, target int) int {
	if len(arr) == 0 || arr[len(arr)-1] < target || arr[0] > target {
		return 0
	}

	start, end := 0, 0
	left, right := 0, len(arr)-1
	//先找start
	for left <= right {
		mid := left + (right-left)/2
		if left < mid && arr[mid] < target {
			left = mid
		} else if mid < right && arr[mid] > target {
			right = mid
		} else {
			//mid=target
			if arr[mid] == target {
				if mid == 0 || (mid > 0 && arr[mid-1] < target) {
					start = mid
					break
				} else {
					right = mid
				}
			} else if left == mid && arr[mid+1] == target {
				start = mid + 1
				break
			} else {
				//找不到
				return 0
			}
		}
	}

	left, right = start, len(arr)-1
	if left == right {
		return 1
	}

	for left < right {
		mid := left + (right-left)/2
		if left < mid && arr[mid] < target {
			left = mid
		} else if mid < right && arr[mid] > target {
			right = mid
			if arr[mid-1] == target {
				end = mid - 1
				break
			}
		} else {
			//mid=target
			if arr[mid] == target {
				if mid == len(arr)-1 || (mid < len(arr)-1 && arr[mid+1] > target) {
					end = mid
					break
				} else {
					left = mid
				}
			} else if right == mid && arr[mid-1] == target {
				end = mid - 1
				break
			} else {
				//找不到
				return 0
			}
		}
	}

	return end - start + 1
}
