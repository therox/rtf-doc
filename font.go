package rtfdoc

import "fmt"

func NewFont(family string, cs int, prq int, name string, code string) Font {
	return Font{Family: family, Charset: cs, Prq: prq, Name: name, Code: code}
}

func (f Font) Compose() string {
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

func NewFontTable() FontTable {
	return FontTable{}
}

func (ft FontTable) Compose() string {
	var fontInfo string
	for i := range ft {
		fontInfo += fmt.Sprintf("{\\f%d%s}", i, ft[i].Compose())

	}
	return fontInfo
}

func (ft *FontTable) AddFont(font Font) int {
	*ft = append(*ft, font)
	return len(*ft)
}
