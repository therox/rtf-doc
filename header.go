package main

import "fmt"

func getDefaultHeader() RTFHeader {
	return RTFHeader{
		Version:  "1",
		CharSet:  "ansi",
		Deff:     "0",
		FontTBL:  "",
		FileTBL:  "",
		ColorTBL: RTFColorTBL{RGBColor(0, 0, 0), RGBColor(255, 0, 0)},
		//ColorTBL:   fmt.Sprintf("%s%s", RGBColor(0, 0, 0).String(), RGBColor(255, 0, 0)),
		StyleSheet: "",
		ListTables: "",
		RevTBL:     "",
	}
}

func composeHeader(header RTFHeader) []byte {
	return []byte(fmt.Sprintf("\\rtf%s\\%s\\deff%s\n"+
		"{\\colortbl;%s}", header.Version, header.CharSet, header.Deff, header.ColorTBL.String()))
}
