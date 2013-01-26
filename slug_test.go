package slug

import "testing"

func TestIsSlugAscii(t *testing.T) {
	tests := []struct {
		s string
		b bool
	}{
		{"", false},
		{"-", false},
		{"A", false},
		{"a", true},
		{"-a", false},
		{"a-", false},
		{"a-0", true},
		{"aa", true},
		{"a--0", false},
	}

	for _, test := range tests {
		if IsSlugAscii(test.s) != test.b {
			t.Error(test.s, "!=", test.b)
		}
	}
}
