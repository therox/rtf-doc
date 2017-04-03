package main

import "fmt"
import "rtf-doc"

func main() {
	//fmt.Println(string(composeHeader(getDefaultHeader())))
	// Создаем документ
	d := rtfdoc.New()
	// Настроить хедер
	d.Header.ColorTBL.SetColor(rtfdoc.Color{255, 0, 0})

	// ...
	t := rtfdoc.NewText("TestText", rtfdoc.Font{}, 48)
	p := rtfdoc.NewParagraph()
	p.AddText(t)
	d.AddContent(p)
	d.SetOrientation("landscape")
	font := rtfdoc.NewFont("roman", 0, 0, "Times New Roman", "tnr")
	fontTable := rtfdoc.NewFontTable()
	fontTable.AddFont(font)
	d.SetFontTable(fontTable)
	fmt.Println(d.Compose())
}
