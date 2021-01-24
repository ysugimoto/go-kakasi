# go-kakasi

go-kakasi is a C binding library of [kakasi](http://kakasi.namazu.org/index.html.ja).
Just wraps pre-built library, you can use kakasi in your Go program.

## Requirements

- Go: We built with `1.16beta1`, I appreciate if you can ensure it's be able to build with older version :-)
- gcc

## Usage

This package bundles prebuilt static library, therefore you don't worry what cgo behaves:

```Go
package main

import (
  "log"
  "github.com/ysugimoto/go-kakasi"
)

func main() {
  text := "日本語"
  transformed, err := kakasi.Transform(text, kakasi.WithKanji(kakasi.ASCII))
  if err != nil {
    log.Fatalln(err)
  }
  log.Printf("Transformed: %s -> %s\n", text, transformed)
}
```

## Customize / Transform Options

go-kakasi accepts transform option as `kakasi.WithXXX` which corresponds to original kakasi transformation mode.

| Option                         | Accept Character                                                                        | remarks                                                 |
| :----------------------------- | :-------------------------------------------------------------------------------------- | :------------------------------------------------------ |
| WithWakatigaki()               | -                                                                                       | Enable wakatigaki mode                                  |
| WithCapitalize()               | -                                                                                       | Enable Capitalize mode, only enables -Ja or -Jj option) |
| WithUppercase()                | -                                                                                       | Enable Upcase mode, only enables -Ja or -Jj option)     |
| WithFurigana                   | -                                                                                       | Enable Furigana mode.                                   |
| WithHiragana(Character)        | kakasi.ASCII, kakasi.JISROMAN, kakasi.KatakanaJIS0201, kakasi.Katakana                  | Set Hiragana transformation mode.                       |
| WithKatakana(Character)        | kakasi.ASCII, kakasi.JISROMAN, kakasi.KatakanaJIS0201, kakasi.Hiragana                  | Set Katakana transformation mode.                       |
| WithKatakanaJIS0201(Character) | kakasi.ASCII, kakasi.JISROMAN, kakasi.Katakana, kakasi.Hiragana                         | Set KatakanaJIS0201 transformation mode.                |
| WithKanji(Character)           | kakasi.ASCII, kakasi.JISROMAN, kakasi.KatakanaJIS0201, kakasi.Katakana, kakasi.Hiragana | Set Kanji transformation mode.                          |
| WithASCII(Character)           | kakasi.JISROMAN, kakasi.Sign                                                            | Set ASCII transformation mode.                          |
| WithJISROMAN(Character)        | kakasi.ASCII, kakasi.Sign                                                               | Set JISROMAN transformation mode.                       |
| WithGraphic(Character)         | kakasi.ASCII, kakasi.JISROMAN, kakasi.Sign                                              | Set Graphic transformation mode.                        |
| WithSign(Character)            | kakasi.ASCII, kakasi.JISROMAN                                                           | Set Sign transformation mode.                           |

## Development

Currently, this package uses kakasi stable version of `2.3.6`, and build environments are:

- **darwin**: OS X 10.15.7 (Catalina)
- **linux**: Ubuntu 16.04 (Xenial, in docker)

If you want to build kakasi for your environment, You can rebuilt static library as following:

```shell
git checkout https://github.com/ysugimoto/go-kakasi
cd go-kakasi
# Build kakasi in your environment
make deps
# Note that if you're linux, it will be overriden `deps/linux/*` files.
```

I appreciate it you let me know some other environment also works fine (especially Windows!)

## License

GNU General Public License, vesion 2
https://www.gnu.org/licenses/old-licenses/gpl-2.0.en.html

## Author

Yoshiaki Sugimoto

And, very thanks for kakasi's anthor: Hironobu Takahasi
