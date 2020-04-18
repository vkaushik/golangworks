package bestTimeToBuySellStock2

import "testing"

func TestSol2(t *testing.T) {
	testDataSet := createTestDataSet()

	for _, testData := range testDataSet {
		actual := maxProfit1(testData.arg1)

		if actual == testData.expected {
			pass(t, testData.expected, actual)
		} else {
			fail(t, testData.expected, actual)
		}
	}
}
