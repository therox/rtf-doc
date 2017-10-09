package rtfdoc

import "fmt"

func getDefaultHeader() header {
	return header{
		version: "1",
		charSet: "ansi",
		deff:    "0",
	}
}

func (h header) compose() string {
	res := fmt.Sprintf("\\rtf%s\\%s\\deff%s", h.version, h.charSet, h.deff)

	if h.fontColor != nil {
		res += fmt.Sprintf("\n{\\fonttbl;%s}", h.fontColor.encode())
	}
	if h.colorTable != nil {
		res += fmt.Sprintf("\n{\\colortbl;%s}", h.colorTable.encode())
	}
	return res
}
