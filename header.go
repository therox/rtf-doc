package rtfdoc

import "fmt"

func getDefaultHeader() header {
	return header{
		version: "1",
		charSet: "ansi",
		Deff:    "0",
		//FileTBL:    "",
		//StyleSheet: "",
		//ListTables: "",
		//RevTBL:     "",
	}
}

func (h header) compose() string {
	res := fmt.Sprintf("\\rtf%s\\%s\\deff%s", h.version, h.charSet, h.Deff)

	if h.ft != nil {
		res += fmt.Sprintf("\n{\\fonttbl;%s}", h.ft.encode())
	}
	if h.ct != nil {
		res += fmt.Sprintf("\n{\\colortbl;%s}", h.ct.encode())
	}
	return res
}
