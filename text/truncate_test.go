package text

import "testing"

func TestTruncate(t *testing.T) {
	cases := []struct {
		name string
		in   string
		n    int
		want string
	}{
		{"shorter than the limit is returned unchanged", "Hello", 10, "Hello"},
		{"exactly at the limit is returned unchanged", "Hello World", 11, "Hello World"},
		{"empty string is returned unchanged", "", 5, ""},
		{"cuts at the last word boundary instead of mid-word", "Hello World", 8, "Hello…"},
		{"falls back to a hard cut when the first word alone is too long", "Supercalifragilisticexpialidocious", 5, "Supe…"},
		{"trims the space left behind by a word-boundary cut", "Hello  World", 8, "Hello…"},
		{"zero limit returns empty", "Hello", 0, ""},
		{"negative limit returns empty", "Hello", -1, ""},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Truncate(tc.in, tc.n); got != tc.want {
				t.Errorf("Truncate(%q, %d) = %q, want %q", tc.in, tc.n, got, tc.want)
			}
		})
	}
}
