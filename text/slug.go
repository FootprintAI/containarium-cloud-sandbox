// Package text holds the one small function this sandbox's engineering
// fixture issue is about.
package text

import "strings"

// Slugify turns a title into a URL-safe slug: lower case, with every run of
// characters that are not ASCII letters or digits replaced by "-".
func Slugify(s string) string {
	var b strings.Builder
	for _, r := range strings.ToLower(s) {
		switch {
		case r >= 'a' && r <= 'z', r >= '0' && r <= '9':
			b.WriteRune(r)
		default:
			b.WriteByte('-')
		}
	}
	return b.String()
}
