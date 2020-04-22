package middleOfLinkedList

func middleNode(head *ListNode) *ListNode {
	fastPtr, slowPtr *ListNode

	for fastPtr, slowPtr = head; true {
		if fastPtr.Next == nil {
			break
		}
		slowPtr = slowPtr.Next
		fastPtr = fastPtr.Next
		if fastPtr.Next == nil {
			break
		}
		fastPtr = fastPtr.Next
	}

	return slowPtr
}
