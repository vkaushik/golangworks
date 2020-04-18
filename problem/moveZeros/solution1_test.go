package moveZeros

import "testing"

func TestSol1(t *testing.T) {
	testDataSet := createTestDataSet()

	for _, testData := range testDataSet {
		sol1(testData.arg1)
		actual := testData.arg1

		if isEqual(testData.expected, actual) {
			pass(t, testData.expected, actual)
		} else {
			fail(t, testData.expected, actual)
		}
	}
}

func isEqual(first, second []int) bool {

	if (first == nil) != (second == nil) {
		return false
	} else if len(first) != len(second) {
		return false
	}

	for i, ifirst := range first {
		if second[i] != ifirst {
			return false
		}
	}

	return true
}

func createTestDataSet() []TestData {
	return []TestData{
		TestData{
			arg1:     []int{},
			expected: []int{},
		},
		TestData{
			arg1:     []int{0, 1, 0, 3, 12},
			expected: []int{1, 3, 12, 0, 0},
		},
		TestData{
			arg1:     []int{6, 1, 7, 3, 12},
			expected: []int{6, 1, 7, 3, 12},
		},
		TestData{
			arg1:     []int{6, 1, 7, 0, 0},
			expected: []int{6, 1, 7, 0, 0},
		},
		TestData{
			arg1:     []int{0, 0, 7, 0, 0},
			expected: []int{7, 0, 0, 0, 0},
		},
		TestData{
			arg1:     []int{1},
			expected: []int{1},
		},
		TestData{
			arg1:     []int{0, 0},
			expected: []int{0, 0},
		},
		TestData{
			arg1:     []int{0},
			expected: []int{0},
		},
		TestData{
			arg1:     []int{1, 0},
			expected: []int{1, 0},
		},
		TestData{
			arg1:     []int{0, 1},
			expected: []int{1, 0},
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
	expected []int
}
