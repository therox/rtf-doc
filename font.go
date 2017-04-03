package rtfdoc

import "fmt"

func NewFont(family string, cs int, prq int, name string) Font {
	return Font{Family: family, Charset: cs, Prq: prq, Name: name}
}

func (f Font) String() string {
	return fmt.Sprintf("%s;", f.Name)

}

func (ft FontTable) String() string {
	var fontinfo string
	for i := range ft {
		fontinfo += fmt.Sprintf("{\\f%d%s%s%s;}", i, ft[i].Family, ft[i].Prq, ft[i].Charset)

	}
	return fontinfo
}

func (ft *FontTable) Add(font Font) int {
	*ft = append(*ft, font)
	return len(*ft)
}
