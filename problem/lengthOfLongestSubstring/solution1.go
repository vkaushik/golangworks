package lengthOfLongestSubstring

func solutionOne(s string) int {
	lengthOfLongestSubstring := 0

	for index, _ := range s {
		substringLength := 0
		isSubstringEnded := false

		for subIndex := index; subIndex < len(s); subIndex++ {

			for lookUpIndex := index; lookUpIndex < subIndex; lookUpIndex++ {
				if s[subIndex] == s[lookUpIndex] {
					isSubstringEnded = true
					break
				}
			}

			if isSubstringEnded {
				break
			} else {
				substringLength++
			}
		}

		if substringLength > lengthOfLongestSubstring {
			lengthOfLongestSubstring = substringLength
		}
	}

	return lengthOfLongestSubstring
}

/*
s => "abcabcd"

0 a
1 b
2 c
3 a
4 b
5 c
6 d

lengthOfLongestSubstring 0
index 0
substringLength 3
isSubstringEnded false
subIndex 3
lookUpIndex 0
*/
