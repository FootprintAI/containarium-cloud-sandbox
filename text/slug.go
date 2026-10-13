// Package text holds the one small function this sandbox's engineering
// fixture issue is about.
package text

import "strings"

// Slugify turns a title into a URL-safe slug: lower case, with every run of
// characters that are not ASCII letters or digits collapsed into a single
// "-", and no leading or trailing "-".
func Slugify(s string) string {
	var b strings.Builder
	inSep := false
	for _, r := range strings.ToLower(s) {
		switch {
		case r >= 'a' && r <= 'z', r >= '0' && r <= '9':
			b.WriteRune(r)
			inSep = false
		default:
			if !inSep && b.Len() > 0 {
				b.WriteByte('-')
			}
			inSep = true
		}
	}
	return strings.TrimSuffix(b.String(), "-")
}
