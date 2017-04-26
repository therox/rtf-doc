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

	ct := d.NewColorTable()

	ct.AddColor(color.RGBA{R: 255, G: 0, B: 0, A: 255}, "Red")
	ct.AddColor(color.RGBA{R: 0, G: 255, B: 0, A: 255}, "Green")
	ct.AddColor(color.RGBA{R: 0, G: 0, B: 255, A: 255}, "Blue")

and font table (with font codes, that we will use later "tnr", "ari", "cs")

	fontTable := d.NewFontTable()
	fontTable.AddFont("roman", 0, 0, "Times New Roman", "tnr")
	fontTable.AddFont("swiss", 0, 0, "Arial", "ari")
	fontTable.AddFont("swiss", 0, 0, "Comic Sans MS", "cs")

Set document orientation

	d.SetOrientation("landscape")

and document format

    d.SetFormat("A4")
	
Add first paragraph with string with Times New Roman font ("tnr" as we defined earlier)

    p := d.AddParagraph()
	p.AddText("Green first string (Times New Roman)", 48, "tnr", "Green")

Add more colourful text

	d.AddParagraph().AddText("Blue second string (Arial)", 48, "ari", "Blue")	
    d.AddParagraph().AddText("Red Third string (Comic Sans)", 48, "cs", "Red")

Add table

	t := d.AddTable()
	t.SetTableMargins(50, 50, 50, 50)
	t.SetWidth(10000)

Add first row to table

	tr := t.AddTableRow()

Get slice of cell widths for first row

	cWidth := t.GetTableCellWidthByRatio(1, 3)

First cell ...
	dc := tr.AddDataCell(cWidth[0])

will be vertical for 2 rows

	dc.SetVerticalMerged(true, false)

Add some text to it

	p = dc.AddParagraph()
	p.AddText("Blue text with cyrillic support with multiline", 16, "cs", "Blue")
	p.AddNewLine()

Add cyrillic (unicode) text

	p.AddText("Голубой кириллический текст с переносом строки внутри параграфа", 16, "cs", "Blue")

Set aligning for last paragraph

	p.SetAlignt("j")

And one more paragraph

	p = dc.AddParagraph()
	p.AddText("Another paragraph in vertical cell", 16, "cs", "Blue")

 with custom indent

	p.SetIndent(40, 0, 0)

 and central aligning

	p.SetAlignt("c")

Add last cell for current row

	dc = tr.AddDataCell(cWidth[1])
	p = dc.AddParagraph()
	p.AddText("Green text In top right cell with center align", 16, "cs", "Green")
	p.SetAlignt("c")

Second row    

	tr = t.AddTableRow()

 with 3 cells

	cWidth = t.GetTableCellWidthByRatio(1, 1.5, 1.5)

 first of which is merged with first cell of the first row

	dc = tr.AddDataCell(cWidth[0])
	dc.SetVerticalMerged(false, true)

Second cell

	dc = tr.AddDataCell(cWidth[1])
	p = dc.AddParagraph()
	p.SetAlignt("r")
	txt := p.AddText("Red text In bottom central cell with right align", 16, "ari", "Red")

 with bold emphasis

	txt.SetEmphasis(true, false, false, false, false, false, false)

Third cell

	dc = tr.AddDataCell(cWidth[2])
	p = dc.AddParagraph()
	p.SetAlignt("l")
	txt = p.AddText("Black text in bottom right cell with left align", 16, "cs", "Black")

 with italic emphasis
 
	txt.SetEmphasis(false, true, false, false, false, false, false)