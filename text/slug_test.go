package text

import "testing"

func TestSlugify(t *testing.T) {
	cases := []struct{ in, want string }{
		{"Hello World", "hello-world"},
		{"already-a-slug", "already-a-slug"},
		{"Go 1.22", "go-1-22"},
		{"", ""},
		{"Hello,  World!!", "hello-world"},
		{"  Padded  ", "padded"},
		{"--x--", "x"},
	}
	for _, tc := range cases {
		if got := Slugify(tc.in); got != tc.want {
			t.Errorf("Slugify(%q) = %q, want %q", tc.in, got, tc.want)
		}
	}
}
