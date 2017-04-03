package rtfdoc

import "fmt"

func (text Text) Compose() string {
	var res string
	res += fmt.Sprintf("\n\\fs%d %s", text.fontSize*2, text.text)
	return res
}

func NewText(text string, font Font, fontSize int) Text {
	return Text{
		fontSize: fontSize,
		text:     text,
		font:     font,
	}
}
