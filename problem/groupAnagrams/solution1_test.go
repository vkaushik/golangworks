package groupAnagrams

import "testing"

func TestSol1(t *testing.T) {
	testDataSet := createTestDataSet()

	for _, testData := range testDataSet {
		actual := groupAnagrams(testData.arg1)

		if isEqual2d(testData.expected, actual) {
			pass(t, testData.expected, actual)
		} else {
			fail(t, testData.expected, actual)
		}
	}
}
func isEqual2d(first, second [][]string) bool {
	if (first == nil) != (second == nil) {
		return false
	} else if len(first) != len(second) {
		return false
	}

	for i, _ := range first {
		if !isEqual(first[i], second[i]) {
			return false
		}
	}

	return true
}

func isEqual(first, second []string) bool {

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
			arg1:     []string{},
			expected: [][]string{},
		},
		TestData{
			arg1: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			expected: [][]string{
				{"ate", "eat", "tea"},
				{"nat", "tan"},
				{"bat"},
			},
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
	arg1     []string
	expected [][]string
}
