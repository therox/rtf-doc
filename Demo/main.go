package main

import (
	"fmt"
	"image/color"

	"github.com/therox/rtf-doc"
)

func main() {

	// Создаем новый чистый, незамутнённый документ

	d := rtfdoc.NewDocument()

	d.SetOrientation("landscape")
	d.SetFormat("A4")

	// Настроить хедер
	// Таблица цветов
	d.NewColorTable().AddColor(color.RGBA{R: 255, G: 0, B: 0, A: 255}, "Red").
		AddColor(color.RGBA{R: 0, G: 255, B: 0, A: 255}, "Green").
		AddColor(color.RGBA{R: 0, G: 0, B: 255, A: 255}, "Blue")

	d.NewFontTable().AddFont("roman", 0, 0, "Times New Roman", "tnr").
		AddFont("swiss", 0, 0, "Arial", "ari").
		AddFont("swiss", 0, 0, "Comic Sans MS", "cs")

	p := d.AddParagraph()
	p.AddText("Green first string (Times New Roman)", 48, "tnr", "Green")
	d.AddParagraph().AddText("Blue second string (Arial)", 48, "ari", "Blue")
	d.AddParagraph().AddText("Red Third string (Comic Sans)", 48, "cs", "Red")

	// Таблица
	t := d.AddTable().SetWidth(10000)
	t.SetTableMargins(50, 50, 50, 50)

	// // строка таблицы
	tr := t.AddTableRow()

	// // Рассчет ячеек таблицы. Первый ряд
	cWidth := t.GetTableCellWidthByRatio(1, 3)

	// ячейка таблицы
	dc := tr.AddDataCell(cWidth[0])
	dc.SetVerticalMerged(true, false)
	p = dc.AddParagraph()
	// текст
	p.AddText("Blue text with cyrillic support with multiline", 16, "cs", "Blue")
	p.AddNewLine()
	p.AddText("Голубой кириллический текст с переносом строки внутри параграфа", 16, "cs", "Blue")
	p.SetAlignt("j")
	p = dc.AddParagraph().SetIndent(40, 0, 0).SetAlignt("c")
	p.AddText("Another paragraph in vertical cell", 16, "cs", "Blue")

	dc = tr.AddDataCell(cWidth[1])
	p = dc.AddParagraph().SetAlignt("c")
	p.AddText("Green text In top right cell with center align", 16, "cs", "Green")
	tr = t.AddTableRow()

	cWidth = t.GetTableCellWidthByRatio(1, 1.5, 1.5)
	// // Это соединенная с верхней ячейка. Текст в ней возьмется из первой ячейки.
	dc = tr.AddDataCell(cWidth[0])
	dc.SetVerticalMerged(false, true)

	dc = tr.AddDataCell(cWidth[1])
	p = dc.AddParagraph()
	p.SetAlignt("r")
	p.AddText("Red text In bottom central cell with right align", 16, "ari", "Red").SetBold()

	dc = tr.AddDataCell(cWidth[2])
	p = dc.AddParagraph()
	p.SetAlignt("l")
	p.AddText("Black text in bottom right cell with left align", 16, "cs", "Black").SetItalic()

	fmt.Println(string(d.Export()))

}
