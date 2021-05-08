package geek_demo

//题目
//输入整数数组 arr ，找出其中最小的 k 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。
//示例 1：
//输入：arr = [3,2,1], k = 2 输出：[1,2] 或者 [2,1] 示例 2：
//输入：arr = [0,1,2,1], k = 1 输出：[0]
//限制：
//0 <= k <= arr.length <= 10000 0 <= arr[i] <= 10000

func getLeastNumbers(arr []int, k int) []int {
	if k == 0 || len(arr) == 0 {
		return []int{}
	}

	counter := make([]int, 10000)

	for i := 0; i < len(arr); i++ {
		counter[arr[i]]++
	}

	result := []int{}
	count := 0
	for i := 0; i < len(counter); i++ {
		if count >= k {
			break
		}
		if counter[i] > 0 {
			for ; counter[i] > 0; counter[i]-- {
				result = append(result, i)
				count++
			}
		}
	}

	return result
}
