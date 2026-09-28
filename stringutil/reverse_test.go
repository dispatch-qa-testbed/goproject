package stringutil

import "testing"

func TestReverse(t *testing.T) {
	tests := []struct {
		name, in, want string
	}{
		{name: "empty", in: "", want: ""},
		{name: "single rune", in: "a", want: "a"},
		{name: "ascii", in: "hello", want: "olleh"},
		{name: "multi-byte CJK", in: "Hello, 世界", want: "界世 ,olleH"},
		{name: "accented", in: "héllo", want: "olléh"},
		{name: "emoji", in: "ab😀", want: "😀ba"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse(tt.in); got != tt.want {
				t.Errorf("Reverse(%q) = %q, want %q", tt.in, got, tt.want)
			}
			if got := Reverse(Reverse(tt.in)); got != tt.in {
				t.Errorf("Reverse(Reverse(%q)) = %q, want %q", tt.in, got, tt.in)
			}
		})
	}
}
