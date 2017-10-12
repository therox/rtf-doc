// Package rtf-doc provides simple tools for creation and writing rtf documents.
// It is very early in development and has suck features as work with text
// (color, font, aligning), tables (merged cells, borders style, thickness and colors),
// and pictures (jpeg or png format)
package rtfdoc

import (
	"fmt"
	"image/color"
)

// NewDocument returns new rtf document instance
func NewDocument() *document {
	doc := document{
		orientation: OrientationPortrait,
		header:      getDefaultHeader(),
		content:     nil,
	}
	doc.marginLeft = 720
	doc.marginRight = 720
	doc.marginTop = 720
	doc.marginBottom = 720

	doc.SetFormat(FormatA4)
	doc.SetOrientation(OrientationPortrait)

	// Default fonts
	ft := doc.NewFontTable()
	ft.AddFont("roman", 0, 2, "Times New Roman", FontTimesNewRoman)
	ft.AddFont("roman", 2, 2, "Symbol", FontSymbol)
	ft.AddFont("swiss", 0, 2, "Arial", FontArial)
	ft.AddFont("swiss", 0, 2, "Comic Sans MS", FontComicSansMS)
	ft.AddFont("modern", 128, 1, "Curier New", FontCourierNew)

	// Default colortable
	ct := doc.NewColorTable()
	ct.AddColor(color.RGBA{R: 0, G: 0, B: 0, A: 255}, ColorBlack)
	ct.AddColor(color.RGBA{R: 0, G: 0, B: 255, A: 255}, ColorBlue)
	ct.AddColor(color.RGBA{R: 0, G: 255, B: 255, A: 255}, ColorAqua)
	ct.AddColor(color.RGBA{R: 0, G: 255, B: 0, A: 255}, ColorLime)
	ct.AddColor(color.RGBA{R: 0, G: 128, B: 0, A: 255}, ColorGreen)
	ct.AddColor(color.RGBA{R: 255, G: 0, B: 255, A: 255}, ColorMagenta)
	ct.AddColor(color.RGBA{R: 255, G: 0, B: 0, A: 255}, ColorRed)
	ct.AddColor(color.RGBA{R: 255, G: 255, B: 0, A: 255}, ColorYellow)
	ct.AddColor(color.RGBA{R: 255, G: 255, B: 255, A: 255}, ColorWhite)
	ct.AddColor(color.RGBA{R: 0, G: 0, B: 128, A: 255}, ColorNavy)
	ct.AddColor(color.RGBA{R: 0, G: 128, B: 128, A: 255}, ColorTeal)
	ct.AddColor(color.RGBA{R: 128, G: 0, B: 128, A: 255}, ColorPurple)
	ct.AddColor(color.RGBA{R: 128, G: 0, B: 0, A: 255}, ColorMaroon)
	ct.AddColor(color.RGBA{R: 128, G: 128, B: 0, A: 255}, ColorOlive)
	ct.AddColor(color.RGBA{R: 128, G: 128, B: 128, A: 255}, ColorGray)
	ct.AddColor(color.RGBA{R: 192, G: 192, B: 192, A: 255}, ColorSilver)

	return &doc
}

func (doc *document) getMargins() string {
	return fmt.Sprintf("\n\\margl%d\\margr%d\\margt%d\\margb%d",
		doc.marginLeft,
		doc.marginRight,
		doc.marginTop,
		doc.marginBottom)
}

func (doc *document) compose() string {
	result := "{"
	result += doc.header.compose()
	if doc.orientation == OrientationLandscape {
		result += fmt.Sprintf("\n\\landscape")
	}
	if doc.pagesize != (size{}) {
		result += fmt.Sprintf("\n\\paperw%d\\paperh%d", doc.pagesize.width, doc.pagesize.height)
	}

	result += doc.getMargins()

	for _, c := range doc.content {
		result += fmt.Sprintf("\n%s", c.compose())
	}
	result += "\n}"
	return result
}

// SetFormat sets page format (A2, A3, A4)
func (doc *document) SetFormat(format string) *document {
	doc.pageFormat = format
	if doc.orientation != "" {
		size, err := getSize(format, doc.orientation)
		if err == nil {
			doc.pagesize = size
		}
	}
	return doc
}

// SetOrientation - sets page orientation (portrait, landscape)
func (doc *document) SetOrientation(orientation string) *document {

	for _, i := range []string{OrientationLandscape, OrientationPortrait} {
		if orientation == i {
			doc.orientation = i
		}
	}
	size, err := getSize(doc.pageFormat, doc.orientation)
	if err == nil {
		doc.pagesize = size
	}

	return doc
}

// // GetDocumentWidth - returns document width
// func (doc *Document) GetDocumentWidth() int {
// 	return doc.pagesize.width
// }

// SetMarginLeft sets left margin for document work area
func (doc *document) SetMarginLeft(value int) *document {
	doc.marginLeft = value
	return doc
}

// SetMarginRight sets right margin for document work area
func (doc *document) SetMarginRight(value int) *document {
	doc.marginRight = value
	return doc
}

// SetMarginTop sets top margin for document work area
func (doc *document) SetMarginTop(value int) *document {
	doc.marginTop = value
	return doc
}

// SetMarginBottom sets bottom margin for document work area
func (doc *document) SetMarginBottom(value int) *document {
	doc.marginBottom = value
	return doc
}

// NewColorTable returns new color table for document
func (doc *document) NewColorTable() *colorTable {
	ct := colorTable{}
	doc.header.colorTable = &ct
	return &ct
}

// NewFontTable returns new font table for document
func (doc *document) NewFontTable() *fontTable {
	ft := fontTable{}
	doc.header.fontColor = &ft
	return &ft
}

// GetMaxContentWidth - returns maximum content width
func (doc *document) GetMaxContentWidth() int {
	return doc.pagesize.width - doc.marginRight - doc.marginLeft
}

// GetTableCellWidthByRatio - returns slice of cell widths from cells ratios
func (doc *document) GetTableCellWidthByRatio(tableWidth int, ratio ...float64) []int {
	tw := tableWidth
	if tw > doc.GetMaxContentWidth() {
		tw = doc.GetMaxContentWidth()
	}
	cellRatioSum := 0.0
	for _, cellRatio := range ratio {
		cellRatioSum += cellRatio
	}
	var cellWidth = make([]int, len(ratio))
	for i := range ratio {
		cellWidth[i] = int(ratio[i] * (float64(tw) / cellRatioSum))
	}
	return cellWidth
}

// Export exports document
func (doc *document) Export() []byte {
	return []byte(doc.compose())
}

// AddFont function adds font to document header
func (doc *document) AddFont(family string, charset int, prq int, name string, code string) *document {
	doc.fontColor.AddFont(family, charset, prq, name, code)
	return doc
}

// AddColor function adds colot to document color table
func (doc *document) AddColor(c color.RGBA, name string) *document {
	doc.colorTable.AddColor(c, name)
	return doc
}
