package problem/rottingOranges

import (
	"reflect"
	"testing"
)

type testData struct {
	inputOrangeGrid [][]int
	output          int
}

func TestOrangeRotting(t *testing.T) {
	testData := []testData{
		{[][]int{{1, 2}}, 1},
		{[][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}, 4},
		{[][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}, -1},
		{[][]int{{0, 2}}, 0},
		{[][]int{{}}, 0},
	}

	for _, data := range testData {
		if actual := orangesRotting(data.inputOrangeGrid); reflect.DeepEqual(actual, data.output) {
			t.Logf("Pass >> Data: %+v | result: %+v", data, actual)
		} else {
			t.Errorf("Fail >> Data: %+v | result: %+v", data, actual)
		}
	}
}
