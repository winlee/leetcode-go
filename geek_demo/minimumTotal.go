package geek_demo

//实际问题解决例子：路径最小值问题
//给定一个三角形，找出自顶向下的最小路径和。每一步只能移动到下一行中相邻的结点上。
//相邻的结点 在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1 的两个结点。
//
//例如，给定三角形：
//[
//[2],
//[3,4],
//[6,5,7],
//[4,1,8,3]
//]
//自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。
/*
	f(0,0)
	f(1,0)=f(0,0)+f(1,0),f(1,1)=f(0,0)+f(1,1)
	f(2,0)=f(1,0)+f(2,0),f(2,1)=min(f(1,0),f(1,1))+f(2,1)

	dp[i][j] = min(dp[i-1][j-1],dp[i-1][j])+triangle[i][j]

*/
//minimumTotal
func minimumTotal(triangle [][]int) int {
	dp := make([][]int, len(triangle))

	for i := 0; i < len(triangle); i++ {
		dp[i] = make([]int, len(triangle[i]))
		if i == 0 {
			dp[i][0] = triangle[0][0]
			continue
		}

		for j := 0; j < len(triangle[i]); j++ {
			if j == 0 {
				dp[i][j] = dp[i-1][j] + triangle[i][j]
			} else if j == len(triangle[i])-1 {
				dp[i][j] = dp[i-1][j-1] + triangle[i][j]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i-1][j-1]) + triangle[i][j]
			}
		}
	}

	minimum := dp[len(triangle)-1][0]
	for _, v := range dp[len(triangle)-1] {
		if minimum > v {
			minimum = v
		}
	}

	return minimum
}
