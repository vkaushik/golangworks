package groupAnagrams

import "sort"

func groupAnagrams(strs []string) [][]string {
	sortedWordSet := make(map[string][]string)
	result := [][]string{}

	for _, str := range strs {
		key := sortString(str)
		if _, isPresent := sortedWordSet[key]; isPresent {
			sortedWordSet[key] = append(sortedWordSet[key], str)
		} else {
			sortedWordSet[key] = []string{str}
		}
	}

	for _, words := range sortedWordSet {
		result = append(result, words)
	}

	return result
}

type sortedRune []rune

func (sr sortedRune) Less(i, j int) bool {
	return sr[i] < sr[j]
}

func (sr sortedRune) Swap(i, j int) {
	sr[i], sr[j] = sr[j], sr[i]
}

func (sr sortedRune) Len() int {
	return len(sr)
}

func sortString(str string) string {
	sr := sortedRune(str)
	sort.Sort(sortedRune(sr))
	return string(sr)
}
