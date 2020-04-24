package backSpaceStringCompare

import "testing"

func TestSol1(t *testing.T) {
	testDataSet := createTestDataSet()

	for _, testData := range testDataSet {
		actual := backspaceCompare(testData.arg1, testData.arg2)

		if actual == testData.expected {
			pass(t, testData.expected, actual, testData.arg1, testData.arg2)
		} else {
			fail(t, testData.expected, actual, testData.arg1, testData.arg2)
		}
	}
}


func createTestDataSet() []TestData {
	return []TestData{
		{
			arg1:     "bxj##tw",
			arg2:     "bxj###tw",
			expected: false,
		},
		{
			arg1:     "nzp#o#g",
			arg2:     "b#nzp#o#g",
			expected: true,
		},
		{
			arg1:     "ab#c",
			arg2:     "ad#c",
			expected: true,
		},
		{
			arg1:     "ab##",
			arg2:     "c#d#",
			expected: true,
		},
		{
			arg1:     "a##c",
			arg2:     "#a#c",
			expected: true,
		},
		{
			arg1:     "a#c",
			arg2:     "b",
			expected: false,
		},
		{
			arg1:     "####",
			arg2:     "#",
			expected: true,
		},
		{
			arg1:     "#a",
			arg2:     "#b",
			expected: false,
		},
		{
			arg1:     "a#c",
			arg2:     "b",
			expected: false,
		},
		{
			arg1:     "#",
			arg2:     "###",
			expected: true,
		},
		{
			arg1:     "a",
			arg2:     "b",
			expected: false,
		},
		{
			arg1:     "a",
			arg2:     "a",
			expected: true,
		},
		{
			arg1:     "a###",
			arg2:     "b###",
			expected: true,
		},
		{
			arg1:     "###a",
			arg2:     "###b",
			expected: false,
		},
		{
			arg1:     "a#fb#c#",
			arg2:     "#a#fb#c#",
			expected: true,
		},
	}
}

const passString string = "Pass | Expected: %+v | Actual: %+v | %v | %v |"
const failString string = "Fail | Expected: %+v | Actual: %+v | %v | %v |"

func pass(t *testing.T, expected interface{}, actual interface{}, arg1, arg2 string) {
	t.Logf(passString, expected, actual, arg1, arg2)
}

func fail(t *testing.T, expected interface{}, actual interface{}, arg1, arg2 string) {
	t.Errorf(failString, expected, actual, arg1, arg2)
}

type TestData struct {
	arg1     string
	arg2     string
	expected bool
}
