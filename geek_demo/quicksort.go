package geek_demo

func QuickSort(a []int) {
	if len(a) == 0 {
		return
	}

	left, right := 0, len(a)-1
	if left < right {

	}
}

//func main() {
//	a := []int{3, 4, 5, 6, 3, 2, 5, 8}
//	quickSort(a, 0, len(a)-1)
//	fmt.Println(a)
//}

func quickSort(a []int, low, high int) {
	if low < high {
		l, r, tmp := low, high, a[low]
		for l < r {
			for l < r && a[r] >= tmp {
				r--
			}
			a[l] = a[r]

			for l < r && a[l] <= tmp {
				l++
			}
			a[r] = a[l]
		}
		a[l] = tmp
		quickSort(a, low, l-1)
		quickSort(a, l+1, high)
	}
}
