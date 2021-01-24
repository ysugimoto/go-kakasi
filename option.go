package kakasi

type Character string

// kakasi's Transformation character modes.
// See kakasi's documentation.
const (
	ASCII           Character = "a"
	JISROMAN                  = "j"
	GRAPHIC                   = "g"
	KatakanaJIS0201           = "k"
	Katakana                  = "K"
	Kanji                     = "J"
	Hiragana                  = "H"
	Sign                      = "E"
)

type TransformOption func() string

// Enable wakatigaki mode
func WithWakatigaki() TransformOption {
	return func() string {
		return "-w"
	}
}

// Enable capitalize romaji mode
func WithCapitalize() TransformOption {
	return func() string {
		return "-C"
	}
}

// Enable upcape romaji mode
func WithUppercase() TransformOption {
	return func() string {
		return "-U"
	}
}

// Enable furigana mode
func WithFurigana() TransformOption {
	return func() string {
		return "-f"
	}
}

// Set Hiragana transformation mode.
// Enable characters are:
// - ASCII
// - JISROMAN
// - KatakanaJIS0201
// - Katakana
func WithHiragana(c Character) TransformOption {
	return func() string {
		return "-" + string(Hiragana) + string(c)
	}
}

// Set Katakana transformation mode.
// Enable characters are:
// - ASCII
// - JISROMAN
// - KatakanaJIS0201
// - Hiragana
func WithKatakana(c Character) TransformOption {
	return func() string {
		return "-" + string(Katakana) + string(c)
	}
}

// Set KatakanaJIS0201 transformation mode.
// Enable characters are:
// - ASCII
// - JISROMAN
// - Katakana
// - Hiragana
func WithKatakanaJIS0201(c Character) TransformOption {
	return func() string {
		return "-" + string(KatakanaJIS0201) + string(c)
	}
}

// Set Kanji transformation mode.
// Enable characters are:
// - ASCII
// - JISROMAN
// - Katakana
// - KatakanaJIS0201
// - Hiragana
func WithKanji(c Character) TransformOption {
	return func() string {
		return "-" + string(Kanji) + string(c)
	}
}

// Set ASCII transformation mode.
// Enable characters are:
// - JISROMAN
// - Sign
func WithASCII(c Character) TransformOption {
	return func() string {
		return "-" + string(ASCII) + string(c)
	}
}

// Set JISROMAN transformation mode.
// Enable characters are:
// - ASCII
// - Sign
func WithJISROMAN(c Character) TransformOption {
	return func() string {
		return "-" + string(JISROMAN) + string(c)
	}
}

// Set Graphic transformation mode.
// Enable characters are:
// - ASCII
// - JISROMAN
// - Sign
func WithGraphic(c Character) TransformOption {
	return func() string {
		return "-" + string(GRAPHIC) + string(c)
	}
}

// Set Sign transformation mode.
// Enable characters are:
// - ASCII
// - JISROMAN
func WithSign(c Character) TransformOption {
	return func() string {
		return "-" + string(Sign) + string(c)
	}
}
