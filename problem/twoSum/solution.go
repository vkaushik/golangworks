package twoSum

func twoSum(nums []int, target int) []int {
	numberToIndexMap := make(map[int]int)

	for index, num := range nums {
		if _, isNumAlreadyAdded := numberToIndexMap[num]; !isNumAlreadyAdded {
			numberToIndexMap[num] = index
		}
	}

	for index, num := range nums {
		pairNumberNeeded := target - num
		if pairNumberIndex, isPairNumberPresent := numberToIndexMap[pairNumberNeeded]; isPairNumberPresent && index != pairNumberIndex {
			return []int{index, pairNumberIndex}
		}
	}

	return nil
}
