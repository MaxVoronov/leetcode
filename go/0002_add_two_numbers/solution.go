package add_two_numbers

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var buf int
	headNode := &ListNode{}
	node := headNode
	for l1 != nil || l2 != nil {
		var v1, v2 int
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}

		sum := v1 + v2 + buf
		buf = sum / 10

		node.Next = &ListNode{Val: sum % 10}
		node = node.Next
	}

	if buf > 0 {
		node.Next = &ListNode{Val: buf}
	}

	return headNode.Next
}
