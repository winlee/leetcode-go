package geek_demo

type Node struct {
	Next *Node
	Val  int
}

//func main() {
//	root := &Node{
//		Next: &Node{
//			//Next: &Node{
//			//	Next: &Node{
//			//		Next: &Node{
//			//			Next: nil,
//			//			Val:  5,
//			//		},
//			//		Val: 4,
//			//	},
//			//	Val: 3,
//			//},
//			Val: 2,
//		},
//		Val: 1,
//	}
//	root1 := root
//
//	for root1 != nil {
//		fmt.Println(root1.Val)
//		root1 = root1.Next
//	}
//
//	node := reverse(root)
//
//	for node != nil {
//		fmt.Println(node.Val)
//		node = node.Next
//	}
//}

func reverse(root *Node) *Node {
	if root == nil {
		return root
	}

	dummy, next := root, root.Next
	dummy.Next = nil

	for next != nil {
		nextNext := next.Next
		next.Next = dummy
		dummy = next
		next = nextNext
	}

	return dummy
}
