package geek_demo

//func main() {
//	a := []int{3, 3, 1}
//	result := findList(a)
//	for _, i2 := range result {
//		fmt.Print(i2, ",")
//	}
//}

// findList 输出每个索引位置的值都比左边小，比右边大
// input:9,8,7,3,4,2,1 output:9,8,7,2,1
// input:3,3,1 output: 1
func findList(a []int) []int {
	var stack []int

	for i := 0; i < len(a); i++ {

		if len(stack) == 0 {
			stack = append(stack, a[i])
		} else {
			add := true
			for len(stack) > 0 && stack[len(stack)-1] <= a[i] {
				stack = stack[:len(stack)-1]
				add = false
			}
			if add {
				stack = append(stack, a[i])
			}
		}
	}

	return stack
}

func findList2(a []int) []int {
	var stack []int
	for i := 0; i < len(a); i++ {
		if len(stack) == 0 {
			stack = append(stack, a[i])
			continue
		}

		add := true
		for len(stack) > 0 && stack[len(stack)-1] <= a[i] {
			add = false
			stack = stack[:len(stack)-1]
		}

		if add {
			stack = append(stack, a[i])
		}
	}

	return stack
}
