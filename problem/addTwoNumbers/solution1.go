package problem/addTwoNumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var prevNode, currentNode, resultList *ListNode
	carry := 0
	sum := 0

	for l1 != nil && l2 != nil {
		sum, carry = add(l1.Val, l2.Val, carry)
		currentNode = createNode(sum)
		linkToPreviousNodeOrSetHeadOfTheList(&prevNode, currentNode, &resultList)

		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		sum, carry = add(l1.Val, carry)
		currentNode = createNode(sum)
		linkToPreviousNodeOrSetHeadOfTheList(&prevNode, currentNode, &resultList)

		l1 = l1.Next
	}

	for l2 != nil {
		sum, carry = add(l2.Val, carry)
		currentNode = createNode(sum)
		linkToPreviousNodeOrSetHeadOfTheList(&prevNode, currentNode, &resultList)

		l2 = l2.Next
	}

	if carry != 0 {
		currentNode = createNode(carry)
		prevNode.Next = currentNode
	}

	return resultList
}

func add(numbers ...int) (sum int, carry int) {
	for _, number := range numbers {
		sum = sum + number
	}
	carry = sum / 10
	sum = sum % 10
	return
}

func createNode(value int) *ListNode {
	return &ListNode{
		Val:  value,
		Next: nil,
	}
}

func linkToPreviousNodeOrSetHeadOfTheList(prevNode **ListNode, currentNode *ListNode, resultList **ListNode) {
	if (*prevNode) != nil {
		(*prevNode).Next = currentNode
	} else {
		*resultList = currentNode
	}
	*prevNode = currentNode
}
