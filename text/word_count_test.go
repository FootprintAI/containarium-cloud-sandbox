package text

import "testing"

func TestWordCount(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"", 0},
		{"   \t\n  ", 0},
		{"Hello", 1},
		{"Hello there friend", 3},
		{"Hello\tthere\n\nfriend", 3},
		{"word, word.", 2},
	}
	for _, tc := range cases {
		if got := WordCount(tc.in); got != tc.want {
			t.Errorf("WordCount(%q) = %d, want %d", tc.in, got, tc.want)
		}
	}
}
