package lengthOfLongestSubstring

func solutionThree(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	leftHand := 0
	maxSubstringLen := 0
	charSet := map[string]int{}
	indexToCharMap := map[int]string{}

	for index, char := range s {
		indexToCharMap[index] = string(char)
	}

	for rightHand := 0; rightHand < len(s); rightHand++ {
		// fmt.Println(rightHand, string(s[rightHand]), leftHand, maxSubstringLen)
		currentChar := string(s[rightHand])
		substringLen := 0

		if duplicateCharIndex, isDuplicateChar := charSet[currentChar]; isDuplicateChar {
			// recalibrate charSet
			for index := leftHand; index <= duplicateCharIndex; index++ {
				charToRemove := indexToCharMap[index]
				delete(charSet, charToRemove)
			}

			// recalibrate leftHand
			leftHand = duplicateCharIndex + 1
		}

		charSet[currentChar] = rightHand

		substringLen = rightHand - leftHand + 1

		if substringLen > maxSubstringLen {
			maxSubstringLen = substringLen
		}
	}

	return maxSubstringLen
}

/*
s -> awkafg

0 a
1 w
2 k
3 a
4 f
5 g

maxLen 3
leftHand 1
rightHand 3
currentChar a
subStringLen 3

charSet
w 1
k 2
a 3

duplicteCharIndex 0

*/
