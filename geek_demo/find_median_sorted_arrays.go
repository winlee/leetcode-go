package geek_demo

//func main() {
//	nums1 := []int{1, 3}
//	nums2 := []int{2}
//	squareRoot := findMedianSortedArrays(nums1, nums2)
//	fmt.Println(squareRoot)
//}
// findMedianSortedArrays 先合并，再求中位数
// https://leetcode-cn.com/problems/median-of-two-sorted-arrays/solution/xun-zhao-liang-ge-you-xu-shu-zu-de-zhong-wei-s-114/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 && len(nums2) == 0 {
		return 0
	}

	result := []int{}
	i, j := 0, 0
	for i < len(nums1) || j < len(nums2) {
		if i >= len(nums1) {
			result = append(result, nums2[j:]...)
			break
		}

		if j >= len(nums2) {
			result = append(result, nums1[i:]...)
			break
		}

		if nums1[i] > nums2[j] {
			result = append(result, nums2[j])
			j++
		} else {
			result = append(result, nums1[i])
			i++
		}
	}

	median := len(result) % 2
	div := len(result) / 2
	if median > 0 {
		return float64(result[div])
	} else {
		return float64(result[div-1]+result[div]) / 2
	}
}
