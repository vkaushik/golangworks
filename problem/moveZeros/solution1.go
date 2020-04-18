package problem/moveZeros

func sol1(nums []int) {
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

	for index := len(nums) - 1; zeroCount > 0; index, zeroCount = index-1, zeroCount-1 {
		nums[index] = 0
	}

	return
}
