package main

import "fmt"
import "rtf-doc"

func main() {

	// Создаем новый чистый, незамутнённый документ

	d := rtfdoc.New()
	// Настроить хедер
	// Таблица цветов
	ct := rtfdoc.ColorTable{}
	ct.AddColor(rtfdoc.Color{255, 0, 0, "Red"})
	ct.AddColor(rtfdoc.Color{0, 255, 0, "Green"})
	ct.AddColor(rtfdoc.Color{0, 0, 255, "Blue"})

	d.Header.ColorTBL = ct
	// ...

	fontTable := rtfdoc.NewFontTable()
	font1 := rtfdoc.NewFont("roman", 0, 0, "Times New Roman", "tnr")
	font2 := rtfdoc.NewFont("swiss", 0, 0, "Arial", "ari")
	font3 := rtfdoc.NewFont("swiss", 0, 0, "Comic Sans MS", "cs")
	fontTable.AddFont(font1)
	fontTable.AddFont(font2)
	fontTable.AddFont(font3)
	d.SetOrientation("landscape")
	d.SetFontTable(fontTable)

	txt := rtfdoc.NewText("TestTNR", 48, "tnr", fontTable, "Green", d.Header.ColorTBL)
	p := rtfdoc.NewParagraph()
	p.AddText(txt)
	d.AddContent(p)
	txt = rtfdoc.NewText("TestARI", 48, "ari", fontTable, "Blue", d.Header.ColorTBL)
	p = rtfdoc.NewParagraph()
	p.AddText(txt)
	d.AddContent(p)

	// Таблица
	t := rtfdoc.NewTable()
	// строка таблицы
	tr := rtfdoc.NewTableRow()
	// ячейка таблицы
	dc := rtfdoc.NewDataCell(d.GetDocumentWidth() / 3)
	// текст
	cell1Data := rtfdoc.NewText("Blue text In Cell", 14, "cs", fontTable, "Blue", d.Header.ColorTBL)
	p = rtfdoc.NewParagraph()
	p.AddText(cell1Data)
	p.SetAlignt("l")
	dc.SetContent(p)
	tr.AddCell(dc)

	dc = rtfdoc.NewDataCell(d.GetDocumentWidth() / 3)
	cell1Data = rtfdoc.NewText("Blue text In Left Cell", 14, "cs", fontTable, "Green", d.Header.ColorTBL)
	p = rtfdoc.NewParagraph()
	p.AddText(cell1Data)
	p.SetAlignt("r")
	dc.SetContent(p)
	tr.AddCell(dc)
	t.AddRow(tr)

	tr = rtfdoc.NewTableRow()
	dc = rtfdoc.NewDataCell(d.GetDocumentWidth() / 3)
	cell1Data = rtfdoc.NewText("Blue text In Top Cell", 14, "tnr", fontTable, "Red", d.Header.ColorTBL)
	p = rtfdoc.NewParagraph()
	p.AddText(cell1Data)
	p.SetAlignt("c")
	dc.SetContent(p)
	tr.AddCell(dc)

	dc = rtfdoc.NewDataCell(d.GetDocumentWidth() / 3)
	cell1Data = rtfdoc.NewText("Blue text In Left Top Cell", 14, "ari", fontTable, "Black", d.Header.ColorTBL)
	cell1Data.SetBold(true)
	p = rtfdoc.NewParagraph()
	p.AddText(cell1Data)
	p.SetAlignt("c")
	dc.SetContent(p)
	tr.AddCell(dc)
	t.AddRow(tr)

	d.AddContent(t)

	fmt.Println(d.Compose())
}
