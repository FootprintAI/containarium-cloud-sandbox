package text

import (
	"strings"
	"time"
)

const wordsPerMinute = 200

// ReadingTime estimates how long s takes to read at 200 words per
// minute, rounded up to a whole minute. Empty or whitespace-only text
// returns 0; any other non-empty text returns at least one minute.
func ReadingTime(s string) time.Duration {
	words := len(strings.Fields(s))
	if words == 0 {
		return 0
	}
	minutes := (words + wordsPerMinute - 1) / wordsPerMinute
	return time.Duration(minutes) * time.Minute
}
