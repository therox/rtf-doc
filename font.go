package rtfdoc

import "fmt"

// AddFont returns font instance
func (ft *FontTable) AddFont(family string, cs int, prq int, name string, code string) *FontTable {
	*ft = append(*ft, Font{Family: family, Charset: cs, Prq: prq, Name: name, Code: code})
	return ft
}

func (f *Font) encode() string {
	var prq, charset string
	if f.Prq != 0 {
		prq = fmt.Sprintf("\\fprq%d", f.Prq)
	}
	if f.Charset != 0 {
		charset = fmt.Sprintf("\\fcharset%d", f.Charset)
	}
	return fmt.Sprintf("\\f%s%s%s %s;", f.Family, prq, charset, f.Name)

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
