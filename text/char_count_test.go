package text

import "testing"

func TestCharCount(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"", 0},
		{"abc", 3},
		{"é", 1},
		{"café", 4},
	}
	for _, tc := range cases {
		if got := CharCount(tc.in); got != tc.want {
			t.Errorf("CharCount(%q) = %d, want %d", tc.in, got, tc.want)
		}
	}
}
