package rtfdoc

import "fmt"

func getDefaultHeader() Header {
	return Header{
		Version: "1",
		CharSet: "ansi",
		Deff:    "0",
		FontTBL: nil,
		//FileTBL:    "",
		ColorTBL: ColorTable{NewColor(0, 0, 0)},
		//StyleSheet: "",
		//ListTables: "",
		//RevTBL:     "",
	}
}

func composeHeader(header Header) string {
	h := fmt.Sprintf("\\rtf%s\\%s\\deff%s", header.Version, header.CharSet, header.Deff)

	if header.FontTBL != nil {
		h += fmt.Sprintf("\n{\\fonttbl;%s}", header.FontTBL)
	}
	if header.ColorTBL != nil {
		h += fmt.Sprintf("\n{\\colortbl;%s}", header.ColorTBL)
	}
	return h
}
