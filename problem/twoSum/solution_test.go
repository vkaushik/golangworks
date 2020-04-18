package problem/twoSum

import (
	"reflect"
	"testing"
)

type testData struct {
	inputNums     []int
	inputTarget   int
	outputIndexes []int
}

func TestTwoSum(t *testing.T) {
	testData := []testData{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
	}

	for _, data := range testData {
		actual := twoSum(data.inputNums, data.inputTarget)
		if actual != nil && reflect.DeepEqual(actual, data.outputIndexes) {
			t.Logf("Pass >> %+v  actual: %v", data, actual)
		} else {
			t.Errorf("Pass >> %+v  actual: %v", data, actual)
		}
	}
}
