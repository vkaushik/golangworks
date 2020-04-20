package maximumSubArray

import "testing"

func TestMaxSubArray3(t *testing.T) {
	testDataSet := createTestDataSet()

	for _, testData := range testDataSet {
		actual := maxSubArray3(testData.arg1)

		if testData.expected == actual {
			pass(t, testData.expected, actual)
		} else {
			fail(t, testData.expected, actual)
		}
	}
}
