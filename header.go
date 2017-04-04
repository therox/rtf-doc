package rtfdoc

import "fmt"

func getDefaultHeader() Header {
	ct := ColorTable{}
	ct.SetColor(Color{0, 0, 0, "Black"})
	return Header{
		Version: "1",
		CharSet: "ansi",
		Deff:    "0",
		FontTBL: nil,
		//FileTBL:    "",
		ColorTBL: ct,
		//StyleSheet: "",
		//ListTables: "",
		//RevTBL:     "",
	}
}

func (h Header) Compose() string {
	res := fmt.Sprintf("\\rtf%s\\%s\\deff%s", h.Version, h.CharSet, h.Deff)

	if h.FontTBL != nil {
		res += fmt.Sprintf("\n{\\fonttbl;%s}", h.FontTBL.Compose())
	}
	if h.ColorTBL != nil {
		res += fmt.Sprintf("\n{\\colortbl;%s}", h.ColorTBL.Compose())
	}
	return res
}
