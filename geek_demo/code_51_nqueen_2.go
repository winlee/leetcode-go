package geek_demo

var (
	colm  = map[int]struct{}{}
	mainm = map[int]struct{}{}
	subm  = map[int]struct{}{}
	ns2   int
)

func solveNQueens2(n int) [][]string {
	ns2 = n
	var path []int

	dfs2(0, path)
	return res
}

func dfs2(row int, path []int) {
	if row == ns2 {
		board := convert2Board(path)
		res = append(res, board)
		return
	}

	for j := 0; j < ns2; j++ {
		if !contains(colm, j) && !contains(mainm, row-j) && !contains(subm, row+j) {
			path = append(path, j)
			colm[j] = struct{}{}
			mainm[row-j] = struct{}{}
			subm[row+j] = struct{}{}

			dfs2(row+1, path)

			delete(subm, row+j)
			delete(mainm, row-j)
			delete(colm, j)
			path = path[:len(path)-1]
		}
	}

}

func contains(m map[int]struct{}, key int) bool {
	if _, ok := m[key]; ok {
		return ok
	}

	return false
}
