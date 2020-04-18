package medianOfTwoSortedArrays

import "testing"

func TestFindMedianSortedArrays(t *testing.T) {
	testDataSet := createTestDataSet()

	for _, testData := range testDataSet {
		actual := findMedianSortedArrays(testData.arg1, testData.arg2)

		if testData.expected == actual {
			pass(t, testData.expected, actual)
		} else {
			fail(t, testData.expected, actual)
		}
	}
}

type TestData struct {
	arg1     []int
	arg2     []int
	expected float64
}

func createTestDataSet() []TestData {
	return []TestData{
		{
			arg1:     []int{1, 5, 9, 10},
			arg2:     []int{2, 4, 5},
			expected: 5,
		},
		{
			arg1:     []int{1, 5, 9, 10},
			arg2:     []int{2, 4, 4, 6},
			expected: 4.5,
		},
		{
			arg1:     []int{1},
			arg2:     []int{2},
			expected: 1.5,
		},
		{
			arg1:     []int{5},
			arg2:     []int{},
			expected: 5,
		},
		{
			arg1:     []int{},
			arg2:     []int{5},
			expected: 5,
		},
	}
}

const passString string = "Pass | Expected: %+v | Actual: %+v"
const failString string = "Fail | Expected: %+v | Actual: %+v"

func pass(t *testing.T, expected interface{}, actual interface{}) {
	t.Logf(passString, expected, actual)
}

func fail(t *testing.T, expected interface{}, actual interface{}) {
	t.Errorf(failString, expected, actual)
}
