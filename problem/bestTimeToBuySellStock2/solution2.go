package bestTimeToBuySellStock2

func maxProfit1(prices []int) int {
	profit := 0

	valley := findValley(prices, 0, len(prices)-1)

	for peak := valley + 1; peak < len(prices); {
		if (peak < (len(prices) - 1)) && (prices[peak] <= prices[peak + 1]) {
			peak++
		} else {
			profit = profit + (prices[peak] - prices[valley])
			valley = findValley(prices, peak+1, len(prices)-1)
			peak = valley + 1
		}
	}

	return profit
}

func findValley(prices []int, from, till int) int {
	for valley := from; valley < till; valley++ {
		if prices[valley] < prices[valley+1] {
			return valley
		}
	}

	return till
}
