package rtfdoc

import "fmt"

func (color Color) Compose() string {
	return fmt.Sprintf("\\red%d\\green%d\\blue%d;", color.Red, color.Green, color.Blue)
}

func (cTbl ColorTable) Compose() string {
	var res string
	for i := range cTbl {
		res += cTbl[i].Compose()
	}
	return res
}

func (cTbl *ColorTable) AddColor(color Color) {
	*cTbl = append(*cTbl, color)

}

func (cTbl *ColorTable) SetColor(color Color) {
	*cTbl = []Color{color}

}
