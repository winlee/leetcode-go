package _002

import (
	"fmt"
	"math"
	"strings"
)

/**
给定仅有小写字母组成的字符串数组 A，返回列表中的每个字符串中都显示的全部字符（包括重复字符）组成的列表。例如，如果一个字符在每个字符串中出现 3 次，但不是 4 次，则需要在最终答案中包含该字符 3 次。

你可以按任意顺序返回答案。

示例 1：

输入：["bella","label","roller"]
输出：["e","l","l"]
示例 2：

输入：["cool","lock","cook"]
输出：["c","o"]

提示：
1 <= A.length <= 100
1 <= A[i].length <= 100
A[i][j] 是小写字母

链接：https://leetcode-cn.com/problems/find-common-characters
*/

// commonChars 官方解答
func commonChars(A []string) (ans []string) {
	minFreq := [26]int {}
	for i := range minFreq {
		minFreq[i] = math.MaxInt32
	}

	for _, word := range A {
		freq := [26]int{}
		for _, b := range word {
			freq[b - 'a']++
		}

		for i, f := range freq {
			if f < minFreq[i]{
				minFreq[i] = f
			}
		}

		fmt.Println("============")
		fmt.Println(freq)
		fmt.Println(minFreq)
	}

	for i := byte(0); i < 26; i++ {
		for j := 0; j < minFreq[i]; j++ {
			ans = append(ans, string('a'+i))
		}
	}

	return
}

func commonChars2(A []string) []string {
	charMap := map[int32]int{}
	for _, ch := range A[0] {
		charMap[ch] = strings.Count(A[0], string(ch))
	}

	for i := 1; i < len(A); i++ {
		for _,ch := range A[0] {
			count := strings.Count(A[i], string(ch))
			if charMap[ch] > count {
				charMap[ch] = count
			}
		}
	}

	var result []string
	for ch, n := range charMap {
		for j := 0; j < n; j++ {
			result = append(result, string(ch))
		}
	}

	return result
}