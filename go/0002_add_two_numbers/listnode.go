package add_two_numbers

type ListNode struct {
	Val  int
	Next *ListNode
}

// Convert linked list to slice (only for tests)
func (n *ListNode) ToSlice() []int {
	result := make([]int, 0)
	node := n
	for {
		result = append(result, node.Val)
		if node.Next == nil {
			break
		}
		node = node.Next
	}

	return result
}

// Return head node of linked list
func SliceToLinkedList(data []int) *ListNode {
	var node *ListNode
	for i := len(data) - 1; i >= 0; i-- {
		node = &ListNode{Val: data[i], Next: node}
	}

	return node
}
