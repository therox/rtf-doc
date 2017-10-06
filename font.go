package rtfdoc

import "fmt"

const (
	FONT_TIMES_NEW_ROMAN = "font_times_new_roman"
	FONT_SYMBOL          = "font_symbol"
	FONT_ARIAL           = "font_arial"
	FONT_COMIC_SANS_MS   = "font_comic_sans_ms"
)

// AddFont returns font instance
func (ft *FontTable) AddFont(family string, cs int, prq int, name string, code string) *FontTable {
	if prq == 0 {
		prq = 2
	}
	*ft = append(*ft, Font{family: family, charset: cs, prq: prq, name: name, code: code})
	return ft
}

func (f *Font) encode() string {
	return fmt.Sprintf("\\f%s\\fprq%d\\fcharset%d %s;", f.family, f.prq, f.charset, f.name)

}

// FontTable

// NewFontTable - returns new font table
func NewFontTable() *FontTable {
	return &FontTable{}
}

func (ft FontTable) encode() string {
	var fontInfo string
	for i := range ft {
		fontInfo += fmt.Sprintf("{\\f%d%s}", i, ft[i].encode())

	}
	return fontInfo
}
