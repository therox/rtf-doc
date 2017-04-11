package rtfdoc

import "fmt"

func getDefaultHeader() Header {
	ct := ColorTable{}
	ct.SetColor(Color{0, 0, 0, "Black"})
	return Header{
		Version: "1",
		CharSet: "ansi",
		Deff:    "0",
		ft:      nil,
		ct:      ct,
		//FileTBL:    "",
		//StyleSheet: "",
		//ListTables: "",
		//RevTBL:     "",
	}
}

func (h Header) compose() string {
	res := fmt.Sprintf("\\rtf%s\\%s\\deff%s", h.Version, h.CharSet, h.Deff)

	if h.ft != nil {
		res += fmt.Sprintf("\n{\\fonttbl;%s}", h.ft.compose())
	}
	if h.ct != nil {
		res += fmt.Sprintf("\n{\\colortbl;%s}", h.ct.Compose())
	}
	return res
}
