package lengthOfLongestSubstring

import (
	"testing"
)

func TestSolutionOne(t *testing.T) {
	testData := []testData{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"", 0},
		{"a", 1},
		{"awkafg", 5},
		{"agwkamhgxyzq", 10},
	}

	for _, test := range testData {
		actual := solutionOne(test.input)

		if actual == test.output {
			t.Logf("PASS. Expected: %d in string: \"%s\".", test.output, test.input)
		} else {
			t.Errorf("FAIL. Expected: %d in string: \"%s\". Actual: %d.", test.output, test.input, actual)
		}
	}
}

type testData struct {
	input  string
	output int
}
