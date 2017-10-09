# RTF Generator

Generates rtf documents

## Installation

``go get -u github.com/therox/rtf-doc``

## Usage

Import package

	import rtfdoc "github.com/therox/rtf-doc"

Create new document instance

    d := rtfdoc.NewDocument()

Setting up header information. First set up color table
Set document orientation

	d.SetOrientation(rtfdoc.OrientationPortrait)

and document format

    d.SetFormat(rtfdoc.FormatA4)
	
Add first paragraph with string with Times New Roman font ("tnr" as we defined earlier)

    p := d.AddParagraph()
	p.AddText("Green first string (Times New Roman)", 48, rtfdoc.FontTimesNewRoman, rtfdoc.ColorGreen)

Add more colourful text

	d.AddParagraph().AddText("Blue second string (Arial)", 48, rtfdoc.FontArial, rtfdoc.ColorBlue)	
    d.AddParagraph().AddText("Red Third string (Comic Sans)", 48, rtfdoc.ComicSansMS, rtfdoc.ColorRed)

Add table

	t := d.AddTable()
	t.SetMarginLeft(50).SetMarginRight(50).SetMarginTop(50).SetMarginBottom(50)

	t.SetWidth(10000)

Add first row to table

	tr := t.AddTableRow()

Get slice of cell widths for current row

	cWidth := t.GetTableCellWidthByRatio(1, 3)

First cell

	dc := tr.AddDataCell(cWidth[0])

will be vertical for 2 rows

	dc.SetVerticalMergedFirst()

Add some text to it

	p = dc.AddParagraph()
	p.AddText("Blue text with cyrillic support with multiline", 16, rtfdoc.FontComicSansMS, rtfdoc.ColorBlue)
	p.AddNewLine()

Add cyrillic (unicode) text

	p.AddText("Голубой кириллический текст с переносом строки внутри параграфа", 16, rtfdoc.FontComicSansMS, rtfdoc.ColorBlue)

Set aligning for last paragraph

	p.SetAlignt(rtfdoc.AlignJustify)

And one more paragraph

	p = dc.AddParagraph()
	p.AddText("Another paragraph in vertical cell", 16, rtfdoc.FontComicSansMS, ColorBlue)

 with custom indent

	p.SetIndentFirstLine(40)

 and central aligning

	p.SetAlignt(rtfdoc.AlignCenter)

Add last cell for current row

	dc = tr.AddDataCell(cWidth[1])
	p = dc.AddParagraph().SetAlignt(rtfdoc.AlignCenter)
	p.AddText("Green text In top right cell with center align", 16, rtfdoc.FontComicSansMS, rtfdoc.ColorGreen)


Second row    

	tr = t.AddTableRow()

 with 3 cells

	cWidth = t.GetTableCellWidthByRatio(1, 1.5, 1.5)

 first of which is merged with first cell of the first row

	dc = tr.AddDataCell(cWidth[0])
	dc.SetVerticalMergedNext()

Second cell

	dc = tr.AddDataCell(cWidth[1])
	p = dc.AddParagraph().SetAlignt(rtfdoc.AlignRight)
	txt := p.AddText("Red text In bottom central cell with right align", 16, rtfdoc.FontArial, rtfdoc.ColorRed)

 with bold emphasis

	txt.SetBold()

Third cell

	dc = tr.AddDataCell(cWidth[2])
	p = dc.AddParagraph().SetAlignt(rtfdoc.AlignLeft)
	txt = p.AddText("Black text in bottom right cell with left align", 16, rtfdoc.FontComicSansMS, rtfdoc.ColorBlack)

 with italic emphasis
 
	txt.SetItalic()