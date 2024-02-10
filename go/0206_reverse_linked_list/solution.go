package reverse_linked_list

func reverseList(head *ListNode) *ListNode {
	var newList, nextNode *ListNode
	currentNode := head

	for currentNode != nil {
		nextNode = currentNode.Next
		currentNode.Next = newList
		newList = currentNode
		currentNode = nextNode
	}

	return newList
}
