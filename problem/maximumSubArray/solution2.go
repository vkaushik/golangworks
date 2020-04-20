package maximumSubArray

func maxSubArray2(nums []int) int {

	return helper(nums, 0, len(nums)-1)
}

func helper(nums []int, firstIndex, lastIndex int) int {
	if lastIndex < firstIndex {
		return 0
	} else if lastIndex == firstIndex { // just one number
		return nums[firstIndex]
	} else if lastIndex-firstIndex == 1 { // two number
		return max(nums[firstIndex], nums[lastIndex], nums[firstIndex] + nums[lastIndex])
	} else {
		mid := (firstIndex + lastIndex) / 2
		firstPartMax := helper(nums, firstIndex, mid - 1)
		secondPartMax := helper(nums, mid + 1, lastIndex)
		midPartMax := midMax(nums, firstIndex, mid, lastIndex)
		return max(firstPartMax, secondPartMax, midPartMax)
	}
}

func midMax(nums []int, firstIndex, mid, lastIndex int) int {
	maxSum := nums[mid]
	sum := 0

	for index := mid; index >= firstIndex; index-- {
		sum = sum + nums[index]
		if sum > maxSum {
			maxSum = sum
		}
	}

	for index := mid + 1; index <= lastIndex; index++ {
		sum = sum + nums[index]
		if sum > maxSum {
			maxSum = sum
		}
	}

	return maxSum
}

func max(first, second, third int) int {
	if first > second {
		if first > third {
			return first
		} else {
			return third
		}
	} else {
		if second > third {
			return second
		} else {
			return third
		}
	}
}
