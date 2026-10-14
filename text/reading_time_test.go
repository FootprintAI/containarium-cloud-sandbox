package text

import (
	"strings"
	"testing"
	"time"
)

func TestReadingTime(t *testing.T) {
	cases := []struct {
		in   string
		want time.Duration
	}{
		{"", 0},
		{"   \t\n  ", 0},
		{"Hello", 1 * time.Minute},
		{"Hello there friend", 1 * time.Minute},
		{strings.Repeat("w ", 200), 1 * time.Minute},
		{strings.Repeat("w ", 400), 2 * time.Minute},
		{strings.Repeat("w ", 450), 3 * time.Minute},
	}
	for _, tc := range cases {
		if got := ReadingTime(tc.in); got != tc.want {
			t.Errorf("ReadingTime(%q) = %v, want %v", tc.in, got, tc.want)
		}
	}
}
