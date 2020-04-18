package maximumSubArray

import "testing"

func TestMaxSubArray(t *testing.T) {
	testDataSet := createTestDataSet()

	for _, testData := range testDataSet {
		actual := maxSubArray(testData.arg1)

		if testData.expected == actual {
			pass(t, testData.expected, actual)
		} else {
			fail(t, testData.expected, actual)
		}
	}
}

func runIt(testData TestData) int {
	return maxSubArray(testData.arg1)
}

func createTestDataSet() []TestData {
	return []TestData{
		TestData{
			arg1:     []int{},
			expected: 0,
		},
		TestData{
			arg1:     []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			expected: 6,
		},
		TestData{
			arg1:     []int{1},
			expected: 1,
		},
		TestData{
			arg1:     []int{1, 3},
			expected: 4,
		},
		TestData{
			arg1:     []int{1, -1},
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
