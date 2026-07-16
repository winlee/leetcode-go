package binarytree

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}

	root1.Val += root2.Val
	root1.Left = mergeTrees(root1.Left, root2.Left)
	root1.Right = mergeTrees(root1.Right, root2.Right)
	return root1
}

func mergeTrees2(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}

	ifv := func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		return node.Val
	}

	ifn := func(node *TreeNode, left bool) *TreeNode {
		if node == nil {
			return nil
		}
		if left {
			return node.Left
		}
		return node.Right
	}

	return &TreeNode{
		Val:   ifv(root1) + ifv(root2),
		Left:  mergeTrees2(ifn(root1, true), ifn(root2, true)),
		Right: mergeTrees2(ifn(root1, false), ifn(root2, false)),
	}
}
