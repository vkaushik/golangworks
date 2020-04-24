package lastStoneWeight

import "sort"

func lastStoneWeight(stones []int) int {
	count := len(stones)

	for count >= 2 {
		sort.Ints(stones)
		d := stones[count-1] - stones[count-2]
		if d == 0 {
			stones = stones[0 : count-2]
		} else {
			stones = append(stones[0:count-2], d)
		}
		count = len(stones)
	}

	if count == 1 {
		return stones[0]
	} 

	return 0
}