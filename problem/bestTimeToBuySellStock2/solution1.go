package bestTimeToBuySellStock2


type buySellDay struct {
	buyDay, sellDay int
}

var maxProfitWithBuySellDay map[buySellDay]int


func maxProfit(prices []int) int {
	maxProfitWithBuySellDay = make(map[buySellDay]int)

	return getMaxProfit(prices, 0, len(prices) - 1)
}

func getMaxProfit(valueOnDay []int, firstDay, lastDay int) int {
	maxProfit := 0
	for ; firstDay < lastDay; firstDay++ {
		profitWithSelectedStartDate := getMaxProfitHelper(valueOnDay, firstDay, lastDay)
		if profitWithSelectedStartDate > maxProfit {
			maxProfit = profitWithSelectedStartDate
		}
	}

	return maxProfit
}

func getMaxProfitHelper(valueOnDay []int, firstDay, lastDay int) int {

	maxProfit := 0

	if firstDay >= lastDay {
		maxProfit = 0
	} else if lastDay - firstDay == 1 {
		maxProfit = valueOnDay[lastDay] - valueOnDay[firstDay]
		if maxProfit < 0 {
			maxProfit = 0
		}
	} else {
		for sellDay := firstDay + 1; sellDay <= lastDay; sellDay++ {
			currentProfit := valueOnDay[sellDay] - valueOnDay[firstDay]
			if currentProfit < 0 {
				currentProfit = 0
			}
			currentBuySell := buySellDay{ 
				buyDay: firstDay, 
				sellDay: sellDay + 1,
			}
			maxProfitFromDaysAfterSellDay := 0
			if p, isPresent := maxProfitWithBuySellDay[currentBuySell]; isPresent {
				maxProfitFromDaysAfterSellDay = p
			} else {
				maxProfitFromDaysAfterSellDay = getMaxProfit(valueOnDay, sellDay + 1, lastDay)
				maxProfitWithBuySellDay[currentBuySell] = maxProfitFromDaysAfterSellDay
			}
			totalProfit := currentProfit + maxProfitFromDaysAfterSellDay
			if totalProfit > maxProfit {
				maxProfit = totalProfit
			}
		}
	}

	return maxProfit
}
