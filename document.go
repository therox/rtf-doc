package rtfdoc

import "fmt"

// NewDocument returns new rtf document instance
func NewDocument() *Document {
	doc := Document{
		orientation: "portrait",
		Header:      getDefaultHeader(),
		documentSettings: documentSettings{
			margins: margins{720, 720, 720, 720},
		},
		content: nil,
	}
	doc.SetFormat("A4")
	return &doc
}

// SetColorTable - set color table in document instance
func (doc *Document) SetColorTable(ct ColorTable) {
	doc.Header.ct = ct
}

func (doc *Document) getMargins() string {
	if doc.margins != (margins{}) {
		return fmt.Sprintf("\n\\margl%d\\margr%d\\margt%d\\margb%d",
			doc.margins.left,
			doc.margins.right,
			doc.margins.top,
			doc.margins.bottom)
	}
	return ""
}

func (doc *Document) compose() string {
	result := "{"
	result += doc.Header.compose()
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

// AddContent adds new content to document instance
// func (doc *Document) AddContent(content *DocumentItem) {
// 	doc.content = append(doc.content, content)
// }

// SetFormat sets page format (A2, A3, A4)
func (doc *Document) SetFormat(format string) {
	doc.pageFormat = format
	if doc.orientation != "" {
		size, err := getSize(format, doc.orientation)
		if err == nil {
			doc.pagesize = size
		}
	}
}

// SetOrientation - sets page orientation (portrait, landscape)
func (doc *Document) SetOrientation(orientation string) {

	if orientation == "landscape" {
		doc.orientation = "\\landscape"
		if doc.pageFormat != "" {
			size, err := getSize(doc.pageFormat, "landscape")
			if err == nil {
				doc.pagesize = size
			}
		}
	} else {
		doc.orientation = ""
		if doc.pageFormat != "" {
			size, err := getSize(doc.pageFormat, "portrait")
			if err == nil {
				doc.pagesize = size
			}
		}
	}
}

// SetFontTable - sets font table in document instance
func (doc *Document) SetFontTable(ft *FontTable) {
	doc.Header.ft = *ft
}

// GetDocumentWidth - returns document width
func (doc *Document) GetDocumentWidth() int {
	return doc.pagesize.width
}

// SetMargins - sets document margins
func (doc *Document) SetMargins(left, top, right, bottom int) {
	doc.margins = margins{
		left,
		right,
		top,
		bottom,
	}
}

// NewColorTable returns new color table
func (doc *Document) NewColorTable() *ColorTable {
	ct := ColorTable{}
	doc.Header.ct = ct
	return &ct
}

// NewFontTable returns new font table
func (doc *Document) NewFontTable() *FontTable {
	ft := FontTable{}
	doc.Header.ft = ft
	return &ft
}

// GetMaxContentWidth - returns maximum content width
func (doc *Document) GetMaxContentWidth() int {
	return doc.pagesize.width - doc.margins.right - doc.margins.left
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
