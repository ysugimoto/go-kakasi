package kakasi

import (
	"testing"
)

func TestOption(t *testing.T) {
	tests := []struct {
		option   TransformOption
		expected string
	}{
		{
			option:   WithWakatigaki(),
			expected: "-w",
		},
		{
			option:   WithCapitalize(),
			expected: "-C",
		},
		{
			option:   WithUppercase(),
			expected: "-U",
		},
		{
			option:   WithFurigana(),
			expected: "-f",
		},
		{
			option:   WithHiragana(ASCII),
			expected: "-Ha",
		},
		{
			option:   WithKatakana(ASCII),
			expected: "-Ka",
		},
		{
			option:   WithKatakanaJIS0201(ASCII),
			expected: "-ka",
		},
		{
			option:   WithKanji(ASCII),
			expected: "-Ja",
		},
		{
			option:   WithASCII(JISROMAN),
			expected: "-aj",
		},
		{
			option:   WithJISROMAN(ASCII),
			expected: "-ja",
		},
		{
			option:   WithGraphic(ASCII),
			expected: "-ga",
		},
		{
			option:   WithSign(ASCII),
			expected: "-Ea",
		},
	}

	for i, tt := range tests {
		v := tt.option()
		if v != tt.expected {
			t.Fatalf("Test[%d] wrong result, expected=%q, got=%q\n", i, tt.expected, v)
		}
	}
}
