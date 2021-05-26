package geek_demo

import "strings"

/*
无负数，数字含有0也正常返回
 1216000
 翻转 0006121

*/

func reverseNumber(num int64) string {
	var stack []string

	for num > 0 {
		stack = append(stack, (string)((int)(num)%10))
		num /= 10
	}

	return strings.Join(stack, "")
}
