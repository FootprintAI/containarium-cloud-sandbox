package text

import "strings"

// WordCount returns the number of maximal runs of non-space characters
// in s. An empty or whitespace-only s returns 0.
func WordCount(s string) int {
	return len(strings.Fields(s))
}
