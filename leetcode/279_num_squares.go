package leetcode

func numSquares(n int) int {
    dp := make([]int, n+1)
    for i := 1; i <= n; i++ {
        dp[i] = i
    }

    squares := []int{}
    for i := 1; i*i <= n; i++ {
        squares = append(squares, i*i)
    }

    for i := 1; i <= n; i++ {
        for _, sq := range squares {
            if sq <= i {
                dp[i] = min(dp[i], dp[i-sq]+1)
            }
        }
    }

    return dp[n]
}