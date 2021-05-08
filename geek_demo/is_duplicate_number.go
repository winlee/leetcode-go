package geek_demo

// IsDuplicateNumber 类似计数排序
func IsDuplicateNumber(a []int, n int) bool {
	if a == nil {
		return false
	}

	var tmp int
	for i := 0; i < n; i++ {
		for a[i] != i {
			if a[a[i]] == a[i] {
				return true
			}
			tmp = a[a[i]]
			a[a[i]] = a[i]
			a[i] = tmp
		}
	}

	return false
}
