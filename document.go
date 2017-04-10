package rtfdoc

import "fmt"

func NewDocument() Document {
	doc := Document{
		orientation: "portrait",
		Header:      getDefaultHeader(),
		DocumentSettings: DocumentSettings{
			Margins: Margins{720, 720, 720, 720},
		},
		Content: nil,
	}
	doc.SetFormat("A4")
	return doc
}

func (doc *Document) compose() string {
	result := "{"
	result += doc.Header.compose()
	if doc.orientation != "" {
		result += fmt.Sprintf("\n%s", doc.orientation)
	}
	if doc.pagesize != (Size{}) {
		result += fmt.Sprintf("\n\\paperw%d\\paperh%d", doc.pagesize.width, doc.pagesize.height)
	}
	if doc.Margins != (Margins{}) {
		result += fmt.Sprintf("\n\\margl%d\\margr%d\\margt%d\\margb%d",
			doc.Margins.left,
			doc.Margins.right,
			doc.Margins.top,
			doc.Margins.bottom)
	}
	for _, c := range doc.Content {
		result += fmt.Sprintf("\n%s", c.compose())
	}
	result += "\n}"
	return result
}

func (doc *Document) AddContent(content DocumentItem) {
	doc.Content = append(doc.Content, content)
}

func (doc *Document) SetFormat(format string) {
	doc.pageFormat = format
	if doc.orientation != "" {
		size, err := getSize(format, doc.orientation)
		if err == nil {
			doc.pagesize = size
		}
	}
}

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

func (doc *Document) SetFontTable(ft FontTable) {
	doc.Header.FontTBL = ft
}

func (doc *Document) GetDocumentWidth() int {
	return doc.pagesize.width
}

func (doc *Document) SetMargins(left, top, right, bottom int) {
	doc.Margins = Margins{
		left,
		right,
		top,
		bottom,
	}
}

//func (doc *Document) getLeftMargin() int {
//	return doc.Margins.left
//}
//func (doc *Document) getRightMargin() int {
//	return doc.Margins.right
//}
//func (doc *Document) getTopMargin() int {
//	return doc.Margins.top
//}
//func (doc *Document) getBottomMargin() int {
//	return doc.Margins.bottom
//}

func (doc *Document) GetMaxContentWidth() int {
	return doc.pagesize.width - doc.Margins.right - doc.Margins.left
}
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

func (doc *Document) Export() []byte {
	return []byte(doc.compose())
}
