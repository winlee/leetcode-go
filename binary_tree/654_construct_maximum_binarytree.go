package binarytree

import "math"

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil 
	}

	rootVal,rootPos := math.MinInt32, 0
	for i,v := range nums {
		if v > rootVal {
			rootVal = v 
			rootPos = i
		}
	}

	return &TreeNode{
		Val:   rootVal,
		Left:  constructMaximumBinaryTree(nums[:rootPos]),
		Right: constructMaximumBinaryTree(nums[rootPos+1:]),
	}
}