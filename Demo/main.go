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
	t.SetTableMargins(50, 50, 50, 50)
	// строка таблицы
	tr := rtfdoc.NewTableRow()

	// Рассчет ячеек таблицы. Первый ряд
	c1 := rtfdoc.GetTableCellWidthByRatio(10000, 1, 3)

	// ячейка таблицы
	dc := rtfdoc.NewDataCell(c1[0])
	// текст
	cell1Data := rtfdoc.NewText("Чоткий текст на нескольких строчках\\line и еще строчка\\line и еще", 16, "cs", fontTable, "Blue", d.Header.ColorTBL)
	dc.SetVerticalMerged(true, false)
	//dc.SetCellMargins(200, 200, 200, 200)
	p = rtfdoc.NewParagraph()
	p.AddText(cell1Data)
	p.SetAlignt("j")
	dc.SetContent(p)
	tr.AddCell(dc)

	dc = rtfdoc.NewDataCell(c1[1])
	cell1Data = rtfdoc.NewText("Blue text In Left Cell", 16, "cs", fontTable, "Green", d.Header.ColorTBL)
	p = rtfdoc.NewParagraph()
	p.AddText(cell1Data)
	p.SetAlignt("r")
	dc.SetContent(p)
	tr.AddCell(dc)
	t.AddRow(tr)

	c2 := rtfdoc.GetTableCellWidthByRatio(10000, 1, 1.5, 1.5)
	// Это соединенная с верхней ячейка. Текст в ней возьмется из первой ячейки.
	tr = rtfdoc.NewTableRow()
	dc = rtfdoc.NewDataCell(c2[0])
	dc.SetVerticalMerged(false, true)
	//dc.SetContent(p)
	tr.AddCell(dc)

	dc = rtfdoc.NewDataCell(c2[1])
	cell1Data = rtfdoc.NewText("Blue text In Left Top Cell", 16, "ari", fontTable, "Black", d.Header.ColorTBL)
	cell1Data.SetEmphasis(true, false, false, false, false, false, false)
	p = rtfdoc.NewParagraph()
	p.AddText(cell1Data)
	p.SetAlignt("c")
	dc.SetContent(p)
	tr.AddCell(dc)

	dc = rtfdoc.NewDataCell(c2[2])
	cell1Data = rtfdoc.NewText("Third Cell", 16, "cs", fontTable, "Black", d.Header.ColorTBL)
	cell1Data.SetEmphasis(false, true, false, false, false, false, false)
	p = rtfdoc.NewParagraph()
	p.AddText(cell1Data)
	p.SetAlignt("c")
	dc.SetContent(p)
	tr.AddCell(dc)
	t.AddRow(tr)

	d.AddContent(t)

	fmt.Println(string(d.Export()))

}
