package kakasi

import (
	"testing"
)

func TestKakasiTransform(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		options  []TransformOption
	}{
		{
			input:    "群馬県",
			expected: "gunmaken",
			options: []TransformOption{
				WithKanji(ASCII),
			},
		},
	}

	for i, tt := range tests {
		transformed, err := Transform(tt.input, tt.options...)
		if err != nil {
			t.Fatalf("Test[%d] unexpected transform error: %s", i, err)
		}
		if transformed != tt.expected {
			t.Fatalf("Test[%d] wrong result, expected=%q, got=%q\n", i, tt.expected, transformed)
		}
	}
}
