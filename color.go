package rtfdoc

import (
	"fmt"
	"image/color"
)

const (
	COLOR_BLACK   = "color_black"
	COLOR_BLUE    = "color_blue"
	COLOR_AQUA    = "color_aqua"
	COLOR_LIME    = "color_lime"
	COLOR_GREEN   = "color_green"
	COLOR_MAGENTA = "color_magenta"
	COLOR_RED     = "color_red"
	COLOR_YELLOW  = "color_yellow"
	COLOR_WHITE   = "color_white"
	COLOR_NAVY    = "color_navy"
	COLOR_TEAL    = "color_teal"
	COLOR_PURPLE  = "color_purple"
	COLOR_MAROON  = "color_maroon"
	COLOR_OLIVE   = "color_olive"
	COLOR_GRAY    = "color_gray"
	COLOR_SILVER  = "color_silver"
)

func (color Color) encode() string {
	r, g, b, _ := color.color.RGBA()
	return fmt.Sprintf("\\red%d\\green%d\\blue%d;", r/256, g/256, b/256)
}

func (cTbl ColorTable) encode() string {
	var res string
	for i := range cTbl {
		res += cTbl[i].encode()
	}
	return res
}

// AddColor adds color to color table
func (cTbl *ColorTable) AddColor(c color.RGBA, name string) *ColorTable {
	*cTbl = append(*cTbl, Color{c, name})
	return cTbl
}
