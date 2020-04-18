package problem/lengthOfLongestSubstring

type empty struct{}

func solutionTwo(s string) int {
	lengthOfLongestSubstring := 0

	for index, _ := range s {
		substringLength := 0
		isSubstringEnded := false
		charSet := make(map[string]empty)

		for subIndex := index; subIndex < len(s); subIndex++ {

			char := string(s[subIndex])
			_, isSubstringEnded = charSet[char]

			if isSubstringEnded {
				break
			} else {
				charSet[char] = empty{}
				substringLength++
			}
		}

		if substringLength > lengthOfLongestSubstring {
			lengthOfLongestSubstring = substringLength
		}
	}

	return lengthOfLongestSubstring
}
