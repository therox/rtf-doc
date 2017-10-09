package rtfdoc

import "fmt"

// AddFont returns font instance
func (ft *fontTable) AddFont(family string, charset int, prq int, name string, code string) *fontTable {
	if prq == 0 {
		prq = 2
	}
	*ft = append(*ft, Font{family: family, charset: charset, prq: prq, name: name, code: code})
	return ft
}

func (f *Font) encode() string {
	return fmt.Sprintf("\\f%s\\fprq%d\\fcharset%d %s;", f.family, f.prq, f.charset, f.name)

}

// FontTable

// NewFontTable - returns new font table
func NewFontTable() *fontTable {
	return &fontTable{}
}

func (ft fontTable) encode() string {
	var fontInfo string
	for i := range ft {
		fontInfo += fmt.Sprintf("{\\f%d%s}", i, ft[i].encode())

	}
	return fontInfo
}
