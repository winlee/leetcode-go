package geek_demo

func partition(arr []int, start, end int) int {
	var (
		pivot = arr[start]
		left  = start
		right = end
	)

	for left != right {
		for left < right && arr[right] >= pivot {
			right--
		}

		for left < right && arr[left] <= pivot {
			left++
		}

		if left < right {
			arr[left], arr[right] = arr[right], arr[left]
		}
	}

	arr[start], arr[left] = arr[left], arr[start]
	return left
}

func partitionv2(arr []int, start, end int) int {
	var (
		mark  = start
		pivot = arr[start]
		point = start + 1
	)

	for point < len(arr) {
		if arr[point] < pivot {
			arr[point], arr[mark] = arr[mark], arr[point]
			mark = point
		}
		point++
	}

	arr[start], arr[mark] = arr[mark], arr[start]
	return mark
}

func quickSortSingle(arr []int, start, end int) int {
	var (
		mark  = start
		pivot = arr[start]
		point = start + 1
	)

	for point < len(arr) {
		if arr[point] < pivot {
			mark++
			arr[mark], arr[point] = arr[point], arr[mark]
		}
		point++
	}

	arr[start], arr[mark] = arr[mark], arr[start]
	return mark
}

func Fib3(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 1
	default:
		return Fib3(n-1) + Fib3(n-2)
	}
}
