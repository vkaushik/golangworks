package moveZeros

import "fmt"

func sol3(nums []int) {
	lastNonZeroIndex := -1
	fmt.Println("==============", nums)
	for index := 0; index < len(nums); index++ {
		if nums[index] != 0 {
			lastNonZeroIndex++
			nums[lastNonZeroIndex], nums[index] = nums[index], nums[lastNonZeroIndex]
			fmt.Println("swapping", nums[lastNonZeroIndex], nums[index])
		}
	}

	return
}
