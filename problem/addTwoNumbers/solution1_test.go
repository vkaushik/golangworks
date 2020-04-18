package addTwoNumbers

import "testing"

func TestAddTwoNumbers(t *testing.T) {
	testDataSet := createTestDataSet()

	for _, testData := range testDataSet {
		actual := addTwoNumbers(testData.listOne, testData.listTwo)

		if testData.expected.Equals(actual) {
			t.Logf("expected: %+v | actual: %+v", testData.expected.toSlice(), actual.toSlice())
			t.Logf("Pass: %+v testData | actual: %+v", testData, actual)
		} else {
			t.Logf("expected: %+v | actual: %+v", testData.expected.toSlice(), actual.toSlice())
			t.Errorf("Fail: %+v testData | actual: %+v", testData, actual)
		}
	}
}

type testData struct {
	listOne  *ListNode
	listTwo  *ListNode
	expected *ListNode
}

func createTestDataSet() []testData {
	return []testData{
		{
			listOne:  createList(2, 4, 3),
			listTwo:  createList(5, 6, 4),
			expected: createList(7, 0, 8),
		},
		{
			listOne:  createList(5),
			listTwo:  createList(5),
			expected: createList(0, 1),
		},
		{
			listOne:  createList(9, 8),
			listTwo:  createList(1),
			expected: createList(0, 9),
		},
		{
			listOne:  createList(9, 9),
			listTwo:  createList(1),
			expected: createList(0, 0, 1),
		},
	}
}

func createList(numbers ...int) *ListNode {
	var head, prevNode *ListNode

	for index, value := range numbers {
		if index == 0 {
			head = &ListNode{
				Val:  value,
				Next: nil,
			}
			prevNode = head
		} else {
			newNode := &ListNode{
				Val:  value,
				Next: nil,
			}
			prevNode.Next = newNode
			prevNode = newNode
		}
	}

	return head
}

func (this *ListNode) Equals(that *ListNode) bool {
	isEqual := false

	for this != nil && that != nil {
		if this.Val != that.Val {
			break
		} else {
			this = this.Next
			that = that.Next
		}
	}

	if this == nil && that == nil {
		isEqual = true
	}

	return isEqual
}

func (this *ListNode) toSlice() []int {
	list := []int{}

	for this != nil {
		list = append(list, this.Val)
		this = this.Next
	}

	return list
}
