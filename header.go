package rtfdoc

import "fmt"

func getDefaultHeader() Header {
	ct := ColorTable{}
	ct.SetColor(Color{0, 0, 0, "Black"})
	return Header{
		Version:   "1",
		CharSet:   "ansi",
		Deff:      "0",
		FontTable: nil,
		//FileTBL:    "",
		ColorTable: ct,
		//StyleSheet: "",
		//ListTables: "",
		//RevTBL:     "",
	}
}

func (h Header) compose() string {
	res := fmt.Sprintf("\\rtf%s\\%s\\deff%s", h.Version, h.CharSet, h.Deff)

	if h.FontTable != nil {
		res += fmt.Sprintf("\n{\\fonttbl;%s}", h.FontTable.compose())
	}
	if h.ColorTable != nil {
		res += fmt.Sprintf("\n{\\colortbl;%s}", h.ColorTable.Compose())
	}
	return res
}
