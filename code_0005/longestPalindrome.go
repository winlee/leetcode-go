package code_0005

//func main() {
//	a := "cbbd"
//	squareRoot := longestPalindrome(a)
//	fmt.Println(squareRoot)
//}

//5. 最长回文子串
//给你一个字符串 s，找到 s 中最长的回文子串。
// longestPalindrome
// fn(n) = fn(n-1) || fn(n-1)+(s[0] == s[n-1])
func longestPalindrome(s string) string {
	if len(s) == 0 || len(s) == 1 {
		return s
	}
	if len(s) == 2 && s[0] == s[1] {
		return s
	}

	subStr := s[1 : len(s)-1]
	sub := longestPalindrome(subStr)
	if s[0] == s[len(s)-1] {
		return s
	}

	return sub
}

func longestPalindromePassed(s string) string {
	length := len(s)
	if length <= 1 {
		return s
	}

	dp := make([][]bool, length)
	start := 0
	maxLen := 1

	for r := 0; r < length; r++ {
		dp[r] = make([]bool, length)
		dp[r][r] = true

		for l := 0; l < r; l++ {
			if s[l] == s[r] && (r-l <= 2 || dp[l+1][r-1]) {
				dp[l][r] = true
			} else {
				dp[l][r] = false
			}

			if dp[l][r] {
				curLen := r - l + 1
				if curLen > maxLen {
					maxLen = curLen
					start = l
				}
			}
		}

	}

	return s[start : start+maxLen]
}

// 该func 错误在i 为右边界，j为左边界
func longestPalindrome2(s string) string {
	length := len(s)
	if length <= 1 {
		return s
	}

	dp := make([][]bool, length)
	start := 1
	maxLen := 1
	for i := 0; i < length; i++ {
		dp[i] = make([]bool, length)
		dp[i][i] = true

		for j := 0; j < i; j++ {
			if s[i] == s[j] && (j-i <= 2 || dp[i+1][j-1]) {
				dp[i][j] = true
			} else {
				dp[i][j] = false
			}

			if dp[i][j] {
				curLen := j - i + 1
				if curLen > maxLen {
					start = i
					maxLen = curLen
				}
			}
		}
	}
	return s[start : start+maxLen]
}

func longestPalindrome3(s string) string {
	if s == "" {
		return ""
	}

	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		left1, right1 := expandAroundCenter(s, i, i)
		left2, right2 := expandAroundCenter(s, i, i+1)
		if right1-left1 > end-start {
			start, end = left1, right1
		}

		if right2-left2 > end-start {
			start, end = left2, right2
		}
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, left int, right int) (int, int) {
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {
	}
	return left + 1, right - 1
}
