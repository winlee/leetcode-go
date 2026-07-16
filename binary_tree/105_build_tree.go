package binarytree

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder)==0{
		return nil 
	}
if len(preorder) == 0 {
        return nil
    }
    
    rootVal := preorder[0]
    rootPos := 0
    for i, v := range inorder {
        if v == rootVal {
            rootPos = i
            break
        }
    }

	var res = &TreeNode{
			Val:  rootVal,
			Left:  buildTree(preorder[1:1+rootPos],inorder[:rootPos]),
			Right:  buildTree(preorder[1+rootPos:],inorder[rootPos+1:]),
		}
	return res 
}

func buildTree2(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0|| len(inorder) == 0 {
        return nil 
    }

    rootVal := preorder[0]
    rootPos := 0
    for i, v := range inorder {
        if v == rootVal {
            rootPos = i
            break
        }
    }

    var res = &TreeNode{
    	Val:   rootVal,
    	Left:  buildTree2(preorder[1:1+rootPos],inorder[:rootPos]),
    	Right: buildTree2(preorder[1+rootPos:],inorder[rootPos+1:]),
    }

    return res 
}