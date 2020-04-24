package lastStoneWeight

import "testing"

func TestSol1(t *testing.T) {
	testDataSet := createTestDataSet()

	for _, testData := range testDataSet {
		actual := lastStoneWeight(testData.arg1)

		if actual == testData.expected {
			pass(t, testData.expected, actual)
		} else {
			fail(t, testData.expected, actual)
		}
	}
}

func createTestDataSet() []TestData {
	return []TestData{
		{
			arg1:     []int{2, 7, 4, 1, 8, 1},
			expected: 1,
		},
		{
			arg1:     []int{1, 1, 3, 3, 5, 5, 7, 7},
			expected: 0,
		},
		{
			arg1:     []int{1, 3, 2, 3, 5, 0},
			expected: 0,
		},
		{
			arg1:     []int{1},
			expected: 1,
		},
		{
			arg1:     []int{9, 9},
			expected: 0,
		},
		{
			arg1:     []int{7, 6, 4, 3, 1},
			expected: 1,
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

type TestData struct {
	arg1     []int
	expected int
}
