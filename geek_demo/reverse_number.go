package geek_demo

import (
	"strconv"
)

/*
无负数，数字含有0也正常返回
 1216000
 翻转 0006121

*/

func reverseNumber(num int64) string {
	var stack []int

	for num > 0 {
		stack = append(stack, (int)(num)%10)
		num /= 10
	}

	var result string
	for _, v := range stack {
		result += strconv.Itoa(v)
	}

	return result
}
