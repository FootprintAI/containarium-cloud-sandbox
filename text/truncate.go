package text

import "strings"

// Truncate shortens s for display to at most n runes, including the
// trailing "…" it adds when a cut is made. If s already fits within n
// runes, it is returned unchanged. Otherwise it is cut at the last word
// boundary at or before the limit so a word isn't split, falling back to
// a hard cut when no boundary exists (e.g. a single word longer than n).
func Truncate(s string, n int) string {
	runes := []rune(s)
	if len(runes) <= n {
		return s
	}
	if n <= 0 {
		return ""
	}
	budget := n - 1
	cut := budget
	for cut > 0 && runes[cut] != ' ' {
		cut--
	}
	if cut == 0 {
		cut = budget
	}
	return strings.TrimRight(string(runes[:cut]), " ") + "…"
}
