package rtfdoc

import (
	"fmt"
	"image/color"
)

// NewDocument returns new rtf document instance
func NewDocument() *Document {
	doc := Document{
		orientation:  "portrait",
		header:       getDefaultHeader(),
		marginLeft:   720,
		marginRight:  720,
		marginTop:    720,
		marginBottom: 720,
		content:      nil,
	}
	doc.SetFormat("A4")

	// Default fonts
	ft := doc.NewFontTable()
	ft.AddFont("roman", 0, 2, "Times New Roman", FONT_TIMES_NEW_ROMAN)
	ft.AddFont("roman", 2, 2, "Symbol", FONT_SYMBOL)
	ft.AddFont("swiss", 0, 2, "Arial", FONT_ARIAL)
	ft.AddFont("swiss", 0, 2, "Comic Sans MS", FONT_COMIC_SANS_MS)

	// Default colortable
	ct := doc.NewColorTable()
	ct.AddColor(color.RGBA{R: 0, G: 0, B: 0, A: 255}, COLOR_BLACK)
	ct.AddColor(color.RGBA{R: 0, G: 0, B: 255, A: 255}, COLOR_BLUE)
	ct.AddColor(color.RGBA{R: 0, G: 255, B: 255, A: 255}, COLOR_AQUA)
	ct.AddColor(color.RGBA{R: 0, G: 255, B: 0, A: 255}, COLOR_LIME)
	ct.AddColor(color.RGBA{R: 0, G: 128, B: 0, A: 255}, COLOR_GREEN)
	ct.AddColor(color.RGBA{R: 255, G: 0, B: 255, A: 255}, COLOR_MAGENTA)
	ct.AddColor(color.RGBA{R: 255, G: 0, B: 0, A: 255}, COLOR_RED)
	ct.AddColor(color.RGBA{R: 255, G: 255, B: 0, A: 255}, COLOR_YELLOW)
	ct.AddColor(color.RGBA{R: 255, G: 255, B: 255, A: 255}, COLOR_WHITE)
	ct.AddColor(color.RGBA{R: 0, G: 0, B: 128, A: 255}, COLOR_NAVY)
	ct.AddColor(color.RGBA{R: 0, G: 128, B: 128, A: 255}, COLOR_TEAL)
	ct.AddColor(color.RGBA{R: 128, G: 0, B: 128, A: 255}, COLOR_PURPLE)
	ct.AddColor(color.RGBA{R: 128, G: 0, B: 0, A: 255}, COLOR_MAROON)
	ct.AddColor(color.RGBA{R: 128, G: 128, B: 0, A: 255}, COLOR_OLIVE)
	ct.AddColor(color.RGBA{R: 128, G: 128, B: 128, A: 255}, COLOR_GRAY)
	ct.AddColor(color.RGBA{R: 192, G: 192, B: 192, A: 255}, COLOR_SILVER)

	return &doc
}

func (doc *Document) getMargins() string {
	return fmt.Sprintf("\n\\margl%d\\margr%d\\margt%d\\margb%d",
		doc.marginLeft,
		doc.marginRight,
		doc.marginTop,
		doc.marginBottom)
}

func (doc *Document) compose() string {
	result := "{"
	result += doc.header.compose()
	if doc.orientation != "" {
		result += fmt.Sprintf("\n%s", doc.orientation)
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
func (doc *Document) SetFormat(format string) *Document {
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
func (doc *Document) SetOrientation(orientation string) *Document {

	if orientation == formatLandscape {
		doc.orientation = "\\landscape"
		if doc.pageFormat != "" {
			size, err := getSize(doc.pageFormat, formatLandscape)
			if err == nil {
				doc.pagesize = size
			}
		}
	} else {
		doc.orientation = ""
		if doc.pageFormat != "" {
			size, err := getSize(doc.pageFormat, formatPortrait)
			if err == nil {
				doc.pagesize = size
			}
		}
	}

	return doc
}

// GetDocumentWidth - returns document width
func (doc *Document) GetDocumentWidth() int {
	return doc.pagesize.width
}

func (doc *Document) SetMarginLeft(value int) *Document {
	doc.marginLeft = value
	return doc
}

func (doc *Document) SetMarginRight(value int) *Document {
	doc.marginRight = value
	return doc
}

func (doc *Document) SetMarginTop(value int) *Document {
	doc.marginTop = value
	return doc
}

func (doc *Document) SetMarginBottom(value int) *Document {
	doc.marginBottom = value
	return doc
}

// NewColorTable returns new color table
func (doc *Document) NewColorTable() *ColorTable {
	ct := ColorTable{}
	blackColor := color.RGBA{R: 0, G: 0, B: 0}
	ct.AddColor(blackColor, "Black")
	doc.header.ct = &ct
	return &ct
}

// NewFontTable returns new font table
func (doc *Document) NewFontTable() *FontTable {
	ft := FontTable{}
	doc.header.ft = &ft
	return &ft
}

// GetMaxContentWidth - returns maximum content width
func (doc *Document) GetMaxContentWidth() int {
	return doc.pagesize.width - doc.marginRight - doc.marginLeft
}

// GetTableCellWidthByRatio - returns slice of cells width from cells ratios
func (doc *Document) GetTableCellWidthByRatio(tableWidth int, ratio ...float64) []int {
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
func (doc *Document) Export() []byte {
	return []byte(doc.compose())
}
