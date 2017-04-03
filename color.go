package rtfdoc

import "fmt"

func NewColor(red, green, blue int) Color {
	return Color{
		Red:   red,
		Green: green,
		Blue:  blue,
	}
}

func (color Color) String() string {
	return fmt.Sprintf("\\red%d\\green%d\\blue%d;", color.Red, color.Green, color.Blue)
}

func (cTbl ColorTable) String() string {
	var res string
	for i := range cTbl {
		res += cTbl[i].String()
	}
	return res
}

func (cTbl *ColorTable) Add(color Color) {
	*cTbl = append(*cTbl, color)

}

func (cTbl *ColorTable) Set(color Color) {
	*cTbl = []Color{color}

}
