package moveZeros

func sol2(nums []int) {
	zeroCount := 0
	lastNonZeroIndex := -1

	for index, num := range nums {
		if zeroCount == len(nums)-(lastNonZeroIndex+1) {
			break
		} else if num == 0 {
			zeroCount++
		} else {
			lastNonZeroIndex++

			if index != lastNonZeroIndex {
				nums[lastNonZeroIndex] = num
				nums[index] = 0
			}
		}
	}

	return
}
