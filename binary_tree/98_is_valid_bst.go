package binarytree

/*中序遍历*/ 

// 
func isValidBST(root *TreeNode) bool {
	var  res []int 
	var traverse func(node *TreeNode)
	traverse = func (node *TreeNode) {
		if node == nil {
			return 
		}
		traverse(node.Left)
		res = append(res, node.Val)
		traverse(node.Right) // 
	}
	traverse(root)

	for i:=1;i < len(res); i++ {
		if res[i]<res[i-1] {
			return false 
		}
	}

	return true 
}
