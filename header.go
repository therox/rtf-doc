package rtfdoc

import "fmt"
import "image/color"

func getDefaultHeader() Header {
	ct := ColorTable{}
	blackColor := color.RGBA{R: 0, G: 0, B: 0}
	ct.SetColor(blackColor, "Black")
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
		res += fmt.Sprintf("\n{\\fonttbl;%s}", h.ft.encode())
	}
	if h.ct != nil {
		res += fmt.Sprintf("\n{\\colortbl;%s}", h.ct.encode())
	}
	return res
}
