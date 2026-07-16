package binarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
示例：
3
/ \
9 20
/ \
15 7
输出：[[3], [9,20], [15,7]]
*/
/**
* Definition for a binary tree node.
* type TreeNode struct {
* Val int
* Left *TreeNode
* Right *TreeNode
* }
 */
func levelOrder(root *TreeNode) [][]int {
	// 你的代码
	queue := []*TreeNode{root}
	res := [][]int{}
	for len(queue) > 0 {
		levelSize := len(queue)
		level := []int{}

		for i := 0; i < levelSize; i++ {
			cur := queue[0]
			queue = queue[1:]
			level = append(level, cur.Val)

			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}

		res = append(res, level)
	}

	return res
}
