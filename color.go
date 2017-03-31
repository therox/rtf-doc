package main

import "fmt"

func RGBColor(red, green, blue int) RTFRGBColor {
	return RTFRGBColor{
		Red:   red,
		Green: green,
		Blue:  blue,
	}
}

func (color RTFRGBColor) String() string {
	return fmt.Sprintf("\\red%d\\green%d\\blue%d;", color.Red, color.Green, color.Blue)
}

func (cTbl RTFColorTBL) String() string {
	var res string
	for i := range cTbl {
		res += cTbl[i].String()
	}
	return res
}
