package diameterOfBinaryTree

import "testing"

func TestDiameterOfBinaryTree(t *testing.T) {
	testDataSet := createTestDataSet()

	for _, testData := range testDataSet {
		actual := diameterOfBinaryTree(testData.arg1)

		if actual == testData.expected {
			pass(t, testData.expected, actual)
		} else {
			fail(t, testData.expected, actual)
		}
	}
}

//**************************************************************

const passString string = "Pass | Expected: %+v | Actual: %+v"
const failString string = "Fail | Expected: %+v | Actual: %+v"

func pass(t *testing.T, expected interface{}, actual interface{}) {
	t.Logf(passString, expected, actual)
}

func fail(t *testing.T, expected interface{}, actual interface{}) {
	t.Errorf(failString, expected, actual)
}

//**************************************************************

type TestData struct {
	arg1     *TreeNode
	expected int
}

func createTestDataSet() []TestData {
	return []TestData{
		{
			arg1:     createTree(),
			expected: 3,
		},
		{
			arg1:     createTree1(),
			expected: 4,
		},
		{
			arg1:     createTree2(),
			expected: 1,
		},
		{
			arg1:     createTree3(),
			expected: 1,
		},
		{
			arg1:     createTree4(),
			expected: 2,
		},
	}
}
//**************************************************************


func createTree() *TreeNode {
	t1 := createNode(1)
	t2 := createNode(2)
	t3 := createNode(3)
	t4 := createNode(4)
	t5 := createNode(5)

	t1.Left = t2
	t1.Right = t3
	t2.Left = t4
	t2.Right = t5

	return t1
}

func createTree1() *TreeNode {
	t1 := createNode(1)
	t2 := createNode(2)
	t3 := createNode(3)
	t4 := createNode(4)
	t5 := createNode(5)
	t6 := createNode(6)

	t1.Left = t2
	t1.Right = t3
	t2.Left = t4
	t2.Right = t5
	t3.Right = t6

	return t1
}

func createTree2() *TreeNode {
	t1 := createNode(1)
	t2 := createNode(2)

	t1.Left = t2

	return t1
}

func createTree3() *TreeNode {
	t1 := createNode(1)
	t2 := createNode(2)

	t1.Right = t2

	return t1
}

func createTree4() *TreeNode {
	t1 := createNode(1)
	t2 := createNode(2)
	t3 := createNode(3)

	t1.Left = t2
	t2.Left = t3

	return t1
}

func createNode(value int) *TreeNode {
	return &TreeNode{
		Left:  nil,
		Right: nil,
		Val:   value,
	}
}