package rtfdoc

import (
	"fmt"
	"image/color"
)

func (col colorItem) encode() string {
	r, g, b, _ := col.rgbColor.RGBA()
	return fmt.Sprintf("\\red%d\\green%d\\blue%d;", r/256, g/256, b/256)
}

func (cTbl colorTable) encode() string {
	var res string
	for i := range cTbl {
		res += cTbl[i].encode()
	}
	return res
}

// AddColor adds color to color table
func (cTbl *colorTable) AddColor(c color.RGBA, name string) *colorTable {
	*cTbl = append(*cTbl, colorItem{c, name})
	return cTbl
}
