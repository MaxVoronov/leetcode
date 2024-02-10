package reverse_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

// ToSlice convert linked list to slice (only for tests)
func (n *ListNode) ToSlice() []int {
	node := n
	result := make([]int, 0)
	for node != nil {
		result = append(result, node.Val)
		node = node.Next
	}

	return result
}

// SliceToLinkedList return head node of linked list (only for tests)
func SliceToLinkedList(data []int) *ListNode {
	var node *ListNode
	for i := len(data) - 1; i >= 0; i-- {
		node = &ListNode{Val: data[i], Next: node}
	}

	return node
}
