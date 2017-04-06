package rtfdoc

import "fmt"

func New() Document {
	return Document{
		orientation: "portrait",
		Header:      getDefaultHeader(),
		DocumentSettings: DocumentSettings{
			PageSize: PageSize(11952, 16848),
			Margins:  Margins{720, 720, 720, 720},
		},
		Content: nil,
	}
}

func (doc *Document) compose() string {
	result := "{"
	result += doc.Header.compose()
	if doc.orientation != "" {
		result += fmt.Sprintf("\n\\%s", doc.orientation)
	}
	if doc.PageSize != (Size{}) {
		result += fmt.Sprintf("\n\\paperw%d\\paperh%d", doc.PageSize.width, doc.PageSize.height)
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

func (doc *Document) SetOrientation(orientation string) {
	if orientation == "landscape" {
		doc.orientation = "landscape"
		doc.PageSize = PageSize(16848, 11952)
	} else {
		doc.orientation = ""
		doc.PageSize = PageSize(11952, 16848)
	}
}

func (doc *Document) SetFontTable(ft FontTable) {
	doc.Header.FontTBL = ft
}

func (doc *Document) GetDocumentWidth() int {
	return doc.PageSize.width
}

func (doc *Document) SetMargins(lm, tm, rm, bm int) {
	doc.Margins = Margins{
		lm,
		rm,
		tm,
		bm,
	}
}
func (doc *Document) getLeftMargin() int {
	return doc.Margins.left
}
func (doc *Document) getRightMargin() int {
	return doc.Margins.right
}
func (doc *Document) getTopMargin() int {
	return doc.Margins.top
}
func (doc *Document) getBottomMargin() int {
	return doc.Margins.bottom
}

func (doc *Document) getMaxContentWidth() int {
	return doc.PageSize.width - doc.Margins.right - doc.Margins.left
}
func (doc *Document) GetTableCellWidthByRatio(tableWidth int, ratio ...float64) []int {
	tw := tableWidth
	if tw > doc.getMaxContentWidth() {
		tw = doc.getMaxContentWidth()
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
