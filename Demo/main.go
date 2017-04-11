package main

import "fmt"
import "rtf-doc"

func main() {

	// Создаем новый чистый, незамутнённый документ

	d := rtfdoc.NewDocument()
	// Настроить хедер
	// Таблица цветов
	ct := rtfdoc.ColorTable{}
	ct.AddColor(rtfdoc.Color{255, 0, 0, "Red"})
	ct.AddColor(rtfdoc.Color{0, 255, 0, "Green"})
	ct.AddColor(rtfdoc.Color{0, 0, 255, "Blue"})

	d.SetColorTable(ct)
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

	p := d.NewParagraph()
	txt := p.NewText("Green first string (Times New Roman)", 48, "tnr", "Green")
	p.AddContent(txt)
	d.AddContent(p)

	p = d.NewParagraph()
	txt = p.NewText("Blue second string (Arial)", 48, "ari", "Blue")
	p.AddContent(txt)
	d.AddContent(p)

	p = d.NewParagraph()
	txt = p.NewText("Red Third string (Comic Sans)", 48, "cs", "Red")
	p.AddContent(txt)
	d.AddContent(p)

	// Таблица
	t := d.NewTable()
	t.SetTableMargins(50, 50, 50, 50)
	t.SetWidth(10000)
	// строка таблицы
	tr := t.NewTableRow()

	// Рассчет ячеек таблицы. Первый ряд
	c1 := t.GetTableCellWidthByRatio(1, 3)

	// ячейка таблицы
	dc := tr.NewDataCell(c1[0])
	// текст
	dc.SetVerticalMerged(true, false)
	p = dc.NewParagraph()
	txt = p.NewText("Кириллический текст в нескольких ячейках на нескольких строчках\\line и еще строчка\\line и еще", 16, "cs", "Blue")
	p.SetAlignt("j")
	p.AddContent(txt)
	dc.SetContent(p)
	tr.AddCell(dc)

	dc = tr.NewDataCell(c1[1])
	p = dc.NewParagraph()
	txt = p.NewText("Blue text In Left Cell", 16, "cs", "Green")
	p.SetAlignt("r")
	p.AddContent(txt)
	dc.SetContent(p)
	tr.AddCell(dc)
	t.AddRow(tr)

	c2 := t.GetTableCellWidthByRatio(1, 1.5, 1.5)
	// Это соединенная с верхней ячейка. Текст в ней возьмется из первой ячейки.
	tr = t.NewTableRow()

	dc = tr.NewDataCell(c2[0])
	dc.SetVerticalMerged(false, true)
	tr.AddCell(dc)

	dc = tr.NewDataCell(c2[1])
	p = dc.NewParagraph()
	p.SetAlignt("c")
	txt = p.NewText("Blue text In Left Top Cell", 16, "ari", "Black")
	txt.SetEmphasis(true, false, false, false, false, false, false)
	p.AddContent(txt)
	dc.SetContent(p)
	tr.AddCell(dc)

	dc = tr.NewDataCell(c2[2])
	p = dc.NewParagraph()
	p.SetAlignt("c")
	txt = p.NewText("Third Cell", 16, "cs", "Black")
	txt.SetEmphasis(false, true, false, false, false, false, false)
	p.AddContent(txt)
	dc.SetContent(p)
	tr.AddCell(dc)
	t.AddRow(tr)
	d.AddContent(t)

	fmt.Println(string(d.Export()))

}
