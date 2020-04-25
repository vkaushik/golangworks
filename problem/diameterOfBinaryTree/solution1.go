package diameterOfBinaryTree

func diameterOfBinaryTree(root *TreeNode) int {
	_, diameter := helper(root)
	return diameter
}

func helper(root *TreeNode) (depth, width int) {
	if (root == nil) || (root.Left == nil && root.Right == nil) {
		return 0, 0
	}
	leftDepth, leftWidth, rightDepth, rightWidth := 0, 0, 0, 0
	if root.Left != nil {
		leftDepth, leftWidth = helper(root.Left)
		leftDepth = leftDepth + 1
	}
	if root.Right != nil {
		rightDepth, rightWidth = helper(root.Right)
		rightDepth = rightDepth + 1
	}
	depth = maxOfTwo(leftDepth, rightDepth)
	width = maxOfThree(leftDepth+rightDepth, leftWidth, rightWidth)
	return
}

func maxOfTwo(first, second int) int {
	if first > second {
		return first
	} else {
		return second
	}
}

func maxOfThree(first, second, third int) int {
	if first > second {
		if first > third {
			return first
		} else {
			return third
		}
	} else {
		if second > third {
			return second
		} else {
			return third
		}
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
