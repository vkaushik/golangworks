package maximumSubArray

func maxSubArray(nums []int) int {
	maxSum := 0
	for _, n := range nums {
		maxSum = maxSum + n
	}
	totalSum := maxSum
	excludedSum := 0

	for i := len(nums) - 1; i > 0; i-- {
		excludedSum = excludedSum + nums[i]

		sum := getMaxSum(i, totalSum-excludedSum, nums)

		if sum > maxSum {
			maxSum = sum
		}

	}

	return maxSum
}

func getMaxSum(size, startingSum int, numbers []int) int {
	maxSum := startingSum
	left := 0
	right := size - 1

	for right < len(numbers)-1 {
		startingSum = startingSum - numbers[left]
		left++
		right++
		startingSum = startingSum + numbers[right]

		if startingSum > maxSum {
			maxSum = startingSum
		}
	}

	return maxSum
}
