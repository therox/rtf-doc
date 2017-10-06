package main

import (
	"fmt"

	"github.com/therox/rtf-doc"
)

func main() {

	// Создаем новый чистый, незамутнённый документ

	d := rtfdoc.NewDocument()

	// Настроить хедер
	d.SetOrientation("landscape")
	d.SetFormat("A4")

	p := d.AddParagraph()
	p.AddText("Green first string (Times New Roman)", 48, rtfdoc.FONT_TIMES_NEW_ROMAN, rtfdoc.COLOR_GREEN)
	d.AddParagraph().AddText("Blue second string (Arial)", 48, rtfdoc.FONT_ARIAL, rtfdoc.COLOR_BLUE)
	d.AddParagraph().AddText("Red Third string (Comic Sans)", 48, rtfdoc.FONT_COMIC_SANS_MS, rtfdoc.COLOR_RED)

	// Таблица
	t := d.AddTable().SetWidth(10000)
	//t.SetLeftMargin(50).SetRightMargin(50).SetTopMargin(50).SetBottomMargin(50)
	t.SetMarginLeft(50).SetMarginRight(50).SetMarginTop(50).SetMarginBottom(50)

	// // строка таблицы
	tr := t.AddTableRow()

	// // Рассчет ячеек таблицы. Первый ряд
	cWidth := t.GetTableCellWidthByRatio(1, 3)

	// ячейка таблицы
	dc := tr.AddDataCell(cWidth[0])
	dc.SetVerticalMergedFirst()
	p = dc.AddParagraph()
	// текст
	p.AddText("Blue text with cyrillic support with multiline", 16, rtfdoc.FONT_COMIC_SANS_MS, rtfdoc.COLOR_BLUE)
	p.AddNewLine()
	p.AddText("Голубой кириллический текст с переносом строки внутри параграфа", 16, rtfdoc.FONT_COMIC_SANS_MS, rtfdoc.COLOR_BLUE)
	p.SetAlignt(rtfdoc.ALIGN_JUSTIFY)
	p = dc.AddParagraph().SetIndent(40, 0, 0).SetAlignt(rtfdoc.ALIGN_CENTER)
	p.AddText("Another paragraph in vertical cell", 16, rtfdoc.FONT_COMIC_SANS_MS, rtfdoc.COLOR_BLUE)

	dc = tr.AddDataCell(cWidth[1])
	p = dc.AddParagraph().SetAlignt(rtfdoc.ALIGN_CENTER)
	p.AddText("Green text In top right cell with center align", 16, rtfdoc.FONT_COMIC_SANS_MS, rtfdoc.COLOR_GREEN)
	tr = t.AddTableRow()

	cWidth = t.GetTableCellWidthByRatio(1, 1.5, 1.5)
	// // Это соединенная с верхней ячейка. Текст в ней возьмется из первой ячейки.
	dc = tr.AddDataCell(cWidth[0])
	dc.SetVerticalMergedNext()

	dc = tr.AddDataCell(cWidth[1])
	p = dc.AddParagraph()
	p.SetAlignt(rtfdoc.ALIGN_RIGHT)
	p.AddText("Red text In bottom central cell with right align", 16, rtfdoc.FONT_ARIAL, rtfdoc.COLOR_RED).SetBold()

	dc = tr.AddDataCell(cWidth[2])
	p = dc.AddParagraph()
	p.SetAlignt(rtfdoc.ALIGN_LEFT)
	p.AddText("Black text in bottom right cell with left align", 16, rtfdoc.FONT_COMIC_SANS_MS, rtfdoc.COLOR_BLACK).SetItalic()

	// f, err := os.Open("pic.jpg")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// pPic := d.AddParagraph()
	// pPic.AddPicture(f, rtfdoc.JPGFORMAT)
	// pPic.SetAlignt(rtfdoc.ALIGN_CENTER)
	// pic.SetWidth(200).SetHeight(150)

	fmt.Println(string(d.Export()))

}
