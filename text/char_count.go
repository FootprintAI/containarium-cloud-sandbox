package text

import "unicode/utf8"

// CharCount returns the number of Unicode code points (runes) in s, not its
// byte length.
func CharCount(s string) int {
	return utf8.RuneCountInString(s)
}
