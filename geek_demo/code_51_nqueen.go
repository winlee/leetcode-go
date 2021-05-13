package geek_demo

import (
	"strings"
)

//N皇后问题

var (
	ns   int
	col  []bool //记录某一列是否放置了皇后
	main []bool //记录主对角线上的单元格是否放置了皇后
	sub  []bool //记录副对角线上的单元格是否放置了皇后

	res [][]string
)

//solveNQueens
func solveNQueens(n int) [][]string {
	if n == 0 {
		return res
	}
	ns = n
	col = make([]bool, n)
	main = make([]bool, 2*n-1)
	sub = make([]bool, 2*n-1)

	path := []int{}
	dfs(0, path)
	return res
}

func dfs(row int, path []int) {
	if row == ns {
		board := convert2Board(path)
		res = append(res, board)
		return
	}

	for j := 0; j < ns; j++ {
		if !col[j] && !main[row-j+ns-1] && !sub[row+j] {
			path = append(path, j)
			col[j] = true
			main[row-j+ns-1] = true
			sub[row+j] = true

			dfs(row+1, path)

			sub[row+j] = false
			main[row-j+ns-1] = false
			col[j] = false
			path = path[:len(path)-1]
		}
	}
}

func convert2Board(path []int) []string {
	var board []string
	for _, v := range path {
		s := strings.Repeat(".", max(0, ns2))
		bytes := []byte(s)
		bytes[v] = 'Q'
		board = append(board, string(bytes))
	}
	return board
}
