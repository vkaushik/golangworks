package problem/moveZeros

import "testing"

func TestSol3(t *testing.T) {
	testDataSet := createTestDataSet()

	for _, testData := range testDataSet {
		sol3(testData.arg1)
		actual := testData.arg1

		if isEqual(testData.expected, actual) {
			pass(t, testData.expected, actual)
		} else {
			fail(t, testData.expected, actual)
		}
	}
}
