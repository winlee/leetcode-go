package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs_FAILED(head *ListNode) *ListNode{
	dummy := &ListNode{}
	if head == nil || head.Next == nil {
		return head
	}

	slow, fast := head, head.Next
	for slow != nil && fast != nil {
		tmp := fast.Next
		dummy.Next = fast 
		fast.Next = slow 
		slow.Next = tmp

		slow = fast.Next
		if fast.Next != nil {
			fast = fast.Next.Next
		}
	}

	return dummy.Next
}

func swapPairs(head *ListNode) *ListNode {
    dummy := &ListNode{Next: head}
    prev := dummy
    
    for prev.Next != nil && prev.Next.Next != nil {
        first := prev.Next
        second := prev.Next.Next
        
        // 交换
        prev.Next = second
        first.Next = second.Next
        second.Next = first
        
        prev = first
    }
    
    return dummy.Next
}
