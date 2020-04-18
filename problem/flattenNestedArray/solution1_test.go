package flattenNestedArray

import "testing"

func TestFlattenArray(t *testing.T) {
	input := make([]int, 3)
	input[0] = []int{1, 1}
	input[1] = 2
	input[2] = []int{1, 1}

	ans := flattenArray(input)

	t.Log(ans)
	t.Log(input)
}
