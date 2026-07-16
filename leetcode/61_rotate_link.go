package leetcode

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	// 算长度
	n := 0
	for p := head; p != nil; p = p.Next {
		n++
	}
	k %= n
	if k == 0 {
		return head
	}

	// 快指针先走 k 步
	fast, slow := head, head
	for i := 0; i < k; i++ {
		fast = fast.Next
	}
	// 一起走
	for fast.Next != nil {
		fast, slow = fast.Next, slow.Next
	}

	// 重连
	newHead := slow.Next
	slow.Next = nil
	fast.Next = head
	return newHead
}
