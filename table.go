package rtfdoc

import "fmt"

// SetMarginLeft function sets table left margin
func (t *table) SetMarginLeft(value int) *table {
	t.marginLeft = value
	return t
}

// SetMarginRight function sets table right margin
func (t *table) SetMarginRight(value int) *table {
	t.marginRight = value
	return t
}

// SetMarginTop function sets table top margin
func (t *table) SetMarginTop(value int) *table {
	t.marginTop = value
	return t
}

// SetMarginBottom function sets table bottom margin
func (t *table) SetMarginBottom(value int) *table {
	t.marginBottom = value
	//tp.margins += fmt.Sprintf(" \\trpaddb%d", value)
	return t
}

//func (tp *tableProperties) getMargins() string {
//	return tp.margins
//}

// SetAlign sets table aligning (c/center, l/left, r/right)
func (t *table) SetAlign(align string) *table {
	for _, i := range []string{AlignCenter, AlignLeft, AlignRight} {
		if i == align {
			t.align = i
		}
	}
	return t
}

// AddTable returns Table instance
func (doc *document) AddTable() *table {
	t := table{
		align:    AlignCenter,
		docWidth: doc.maxWidth,
	}
	t.SetMarginLeft(100).SetMarginRight(100).SetMarginTop(100).SetMarginBottom(100)

	t.colorTable = doc.colorTable
	t.fontColor = doc.fontColor
	t.SetBorderLeft(true).
		SetBorderRight(true).
		SetBorderTop(true).
		SetBorderBottom(true).
		SetBorderStyle(BorderSingleThickness).
		SetBorderColor(ColorBlack).
		SetBorderWidth(15)
	t.updateMaxWidth()
	doc.content = append(doc.content, &t)
	return &t
}

func (t *table) updateMaxWidth() *table {
	t.maxWidth = t.docWidth - t.marginLeft - t.marginRight
	return t
}

func (t table) compose() string {
	res := ""
	var align = ""
	if t.align != "" {
		align = fmt.Sprintf("\\trq%s", t.align)
	}
	for _, tr := range t.data {
		res += fmt.Sprintf("\n{\\trowd %s", align)

		res += fmt.Sprintf("\n \\trpaddl%d \\trpaddr%d \\trpaddt%d \\trpaddb%d\n", t.marginLeft, t.marginRight, t.marginTop, t.marginBottom)
		//res += t.getMargins()
		res += tr.encode()
		res += "\\row}"
	}
	return res
}

// AddTableRow returns new table row instance
func (t *table) AddTableRow() *tableRow {
	tr := tableRow{
		generalSettings: generalSettings{
			fontColor:  t.fontColor,
			colorTable: t.colorTable,
		},
		tableWidth: t.maxWidth,
	}
	tr.SetBorderLeft(t.borderLeft).
		SetBorderRight(t.borderRight).
		SetBorderTop(t.borderTop).
		SetBorderBottom(t.borderBottom).
		SetBorderStyle(t.borderStyle).
		SetBorderColor(t.borderColor).
		SetBorderWidth(t.borderWidth)
	t.updateMaxWidth()
	t.data = append(t.data, &tr)
	return &tr
}

func (tr *tableRow) updateMaxWidth() *tableRow {
	tr.maxWidth = tr.tableWidth
	return tr
}

// SetBorderLeft function sets table left border presence
func (t *table) SetBorderLeft(isBorder bool) *table {
	t.borderLeft = isBorder
	return t
}

// SetBorderRight function sets table right border presence
func (t *table) SetBorderRight(isBorder bool) *table {
	t.borderRight = isBorder
	return t
}

// SetBorderTop function sets table top border presence
func (t *table) SetBorderTop(isBorder bool) *table {
	t.borderTop = isBorder
	return t
}

// SetBorderBottom function sets table bottom border presence
func (t *table) SetBorderBottom(isBorder bool) *table {
	t.borderBottom = isBorder
	return t
}

// SetBorderStyle function sets table left border style
func (t *table) SetBorderStyle(bStyle string) *table {
	for _, i := range []string{
		BorderDashSmall,
		BorderDashed,
		BorderDotDash,
		BorderDotDotDash,
		BorderDotted,
		BorderDouble,
		BorderDoubleThickness,
		BorderWavyDouble,
		BorderEmboss,
		BorderEngrave,
		BorderHairline,
		BorderInset,
		BorderOutset,
		BorderShadowed,
		BorderSingleThickness,
		BorderStripped,
		BorderThickThinLarge,
		BorderThickThinMedium,
		BorderThickThinSmall,
		BorderThinThickLarge,
		BorderThinThickMedium,
		BorderThinThickSmall,
		BorderThinThickThinLarge,
		BorderThinThickThinMedium,
		BorderTriple,
		BorderWavy,
	} {
		if bStyle == i {
			t.borderStyle = i
			for tr := range t.data {
				t.data[tr].SetBorderStyle(i)
			}
			break
		}
	}
	return t
}

// SetBorderColor function sets color of the table's border and it's rows and cells
func (t *table) SetBorderColor(color string) *table {
	t.borderColor = color
	for tr := range t.data {
		t.data[tr].SetBorderColor(color)
	}
	return t
}

// SetBorderWidth function sets width of the table's border and it's rows and cells
func (t *table) SetBorderWidth(value int) *table {
	t.borderWidth = value
	for tr := range t.data {
		t.data[tr].SetBorderWidth(value)
	}
	return t
}

// SetWidth sets width of Table
func (t *table) SetWidth(width int) *table {
	t.width = width
	return t
}

// SetBorderLeft function sets left border presence
func (tr *tableRow) SetBorderLeft(isBorder bool) *tableRow {
	tr.borderLeft = isBorder
	return tr
}

// SetBorderRight function sets right border presence
func (tr *tableRow) SetBorderRight(isBorder bool) *tableRow {
	tr.borderRight = isBorder
	return tr
}

// SetBorderTop function sets top border presence
func (tr *tableRow) SetBorderTop(isBorder bool) *tableRow {
	tr.borderTop = isBorder
	return tr
}

// SetBorderBottom function sets bottom border presence
func (tr *tableRow) SetBorderBottom(isBorder bool) *tableRow {
	tr.borderBottom = isBorder
	return tr
}

// SetBorderStyle function sets border style
func (tr *tableRow) SetBorderStyle(bStyle string) *tableRow {
	for _, i := range []string{
		BorderDashSmall,
		BorderDashed,
		BorderDotDash,
		BorderDotDotDash,
		BorderDotted,
		BorderDouble,
		BorderDoubleThickness,
		BorderWavyDouble,
		BorderEmboss,
		BorderEngrave,
		BorderHairline,
		BorderInset,
		BorderOutset,
		BorderShadowed,
		BorderSingleThickness,
		BorderStripped,
		BorderThickThinLarge,
		BorderThickThinMedium,
		BorderThickThinSmall,
		BorderThinThickLarge,
		BorderThinThickMedium,
		BorderThinThickSmall,
		BorderThinThickThinLarge,
		BorderThinThickThinMedium,
		BorderTriple,
		BorderWavy,
	} {
		if bStyle == i {
			tr.borderStyle = i
			for c := range tr.cells {
				tr.cells[c].SetBorderStyle(i)
			}
			break
		}
	}
	return tr
}

// SetBorderColor sets border color of the row (and recursevely on its cells)
func (tr *tableRow) SetBorderColor(color string) *tableRow {
	tr.borderColor = color
	for c := range tr.cells {
		tr.cells[c].SetBorderColor(color)
	}
	return tr
}

// SetBorderWidth sets border width (and recursevely on its cells)
func (tr *tableRow) SetBorderWidth(value int) *tableRow {
	tr.borderWidth = value
	for c := range tr.cells {
		tr.cells[c].SetBorderWidth(value)
	}
	return tr
}

func (tr *tableRow) encode() string {
	res := ""
	// Border settings
	bTempl := "\n \\trbrdr%s\\brdrw%d\\brdr%s"
	for c := range *tr.colorTable {
		if ((*tr.colorTable)[c]).name == tr.borderColor {
			bTempl += fmt.Sprintf("\\brdrcf%d", c+1)
		}

	}

	if tr.borderLeft {
		res += fmt.Sprintf(bTempl, "l", tr.borderWidth, tr.borderStyle)
	}
	if tr.borderRight {
		res += fmt.Sprintf(bTempl, "r", tr.borderWidth, tr.borderStyle)
	}
	if tr.borderTop {
		res += fmt.Sprintf(bTempl, "t", tr.borderWidth, tr.borderStyle)
	}
	if tr.borderBottom {
		res += fmt.Sprintf(bTempl, "b", tr.borderWidth, tr.borderStyle)
	}

	if len(tr.cells) != 0 {
		cellLengthPosition := 0
		for _, tc := range tr.cells {

			cellLengthPosition += tc.getCellWidth()
			res += tc.cellComposeProperties()
			res += fmt.Sprintf("\\cellx%d", cellLengthPosition)

		}
		res += "\n"
		for _, tc := range tr.cells {
			res += tc.cellComposeData()
		}
	}
	return res
}

// AddDataCell returns new DataCell for current table row
func (tr *tableRow) AddDataCell(width int) *tableCell {
	dc := tableCell{
		cellWidth: width,
		maxWidth:  width,
	}
	dc.fontColor = tr.fontColor
	dc.colorTable = tr.colorTable
	dc.SetBorderLeft(tr.borderLeft).
		SetBorderRight(tr.borderRight).
		SetBorderTop(tr.borderTop).
		SetBorderBottom(tr.borderBottom).
		SetBorderStyle(tr.borderStyle).
		SetBorderColor(tr.borderColor).
		SetBorderWidth(tr.borderWidth)
	dc.updateMaxWidth()
	tr.cells = append(tr.cells, &dc)
	return &dc
}

func (dc *tableCell) updateMaxWidth() *tableCell {
	dc.maxWidth = dc.cellWidth - dc.marginLeft - dc.marginRight
	return dc
}

// SetWidth sets width of the cell
func (dc *tableCell) SetWidth(cellWidth int) *tableCell {
	dc.cellWidth = cellWidth
	return dc
}

// AddParagraph creates cell's paragraph
func (dc *tableCell) AddParagraph() *paragraph {
	p := paragraph{
		isTable: true,
		align:   "l",
		indent:  "\\fl360",
		generalSettings: generalSettings{
			colorTable: dc.colorTable,
			fontColor:  dc.fontColor,
		},
		allowedWidth: dc.maxWidth,
	}
	p.updateMaxWidth()
	dc.content = append(dc.content, &p)
	return &p
}

func (dc tableCell) cellComposeProperties() string {
	res := ""
	// Тута свойства ячейки (границы, все дела...)
	bTempl := "\n \\clbrdr%s\\brdrw%d\\brdr%s"
	for c := range *dc.colorTable {
		if ((*dc.colorTable)[c]).name == dc.borderColor {
			bTempl += fmt.Sprintf("\\brdrcf%d", c+1)
		}

	}

	if dc.borderLeft {
		res += fmt.Sprintf(bTempl, "l", dc.borderWidth, dc.borderStyle)
	}
	if dc.borderRight {
		res += fmt.Sprintf(bTempl, "r", dc.borderWidth, dc.borderStyle)
	}
	if dc.borderTop {
		res += fmt.Sprintf(bTempl, "t", dc.borderWidth, dc.borderStyle)
	}
	if dc.borderBottom {
		res += fmt.Sprintf(bTempl, "b", dc.borderWidth, dc.borderStyle)
	}

	// Margins
	res += fmt.Sprintf("\n\\clpadl%d\\clpadr%d\\clpadt%d\\clpadb%d",
		dc.marginLeft, dc.marginRight, dc.marginTop, dc.marginBottom,
	)

	// Vertical Merged
	if dc.verticalMerged != "" {
		res += fmt.Sprintf("\\clvm%s", dc.verticalMerged)
	}

	// Aligning insite cell
	res += fmt.Sprintf("\\clvertal%s", dc.vTextAlign)

	return res
}

func (dc tableCell) cellComposeData() string {
	if len(dc.content) == 0 {
		dc.AddParagraph()
	}
	res := ""
	for _, p := range dc.content {
		res += fmt.Sprintf("%s \n", p.compose())
	}
	res += "\\cell"
	return res
}

func (dc tableCell) getCellWidth() int {
	return dc.cellWidth
}

// SetBorders sets borders to
// datacell

// SetBorderLeft function set left border to be visible
func (dc *tableCell) SetBorderLeft(value bool) *tableCell {
	dc.borderLeft = value
	return dc
}

// SetBorderRight function sets right border to be visible
func (dc *tableCell) SetBorderRight(value bool) *tableCell {
	dc.borderRight = value
	return dc
}

// SetBorderTop function sets top border to be visible
func (dc *tableCell) SetBorderTop(value bool) *tableCell {
	dc.borderTop = value
	return dc
}

// SetBorderBottom function sets bottom border to be visible
func (dc *tableCell) SetBorderBottom(value bool) *tableCell {
	dc.borderBottom = value
	return dc
}

// SetBorderWidth function sets cell's border width px
func (dc *tableCell) SetBorderWidth(value int) *tableCell {
	dc.borderWidth = value
	return dc
}

// SetBorderStyle function sets cell's border style
func (dc *tableCell) SetBorderStyle(bStyle string) *tableCell {
	bStyle = BorderSingleThickness
	for _, i := range []string{
		BorderDashSmall,
		BorderDashed,
		BorderDotDash,
		BorderDotDotDash,
		BorderDotted,
		BorderDouble,
		BorderDoubleThickness,
		BorderWavyDouble,
		BorderEmboss,
		BorderEngrave,
		BorderHairline,
		BorderInset,
		BorderOutset,
		BorderShadowed,
		BorderSingleThickness,
		BorderStripped,
		BorderThickThinLarge,
		BorderThickThinMedium,
		BorderThickThinSmall,
		BorderThinThickLarge,
		BorderThinThickMedium,
		BorderThinThickSmall,
		BorderThinThickThinLarge,
		BorderThinThickThinMedium,
		BorderTriple,
		BorderWavy,
	} {
		if bStyle == i {
			dc.borderStyle = i
			break
		}
	}
	return dc
}

// GetTableCellWidthByRatio returns slice of cell widths
func (t *table) GetTableCellWidthByRatio(ratio ...float64) []int {

	cellRatioSum := 0.0
	for _, cellRatio := range ratio {
		cellRatioSum += cellRatio
	}
	var cellWidth = make([]int, len(ratio))
	for i := range ratio {
		cellWidth[i] = int(ratio[i] * (float64(t.width) / cellRatioSum))
	}
	return cellWidth
}

// SetVerticalMergedFirst sets this cell to be first in vertical merging.
func (dc *tableCell) SetVerticalMergedFirst() *tableCell {
	dc.verticalMerged = "gf"
	return dc
}

// SetVerticalMergedNext sets this cell to be not first cell in vertical merging.
func (dc *tableCell) SetVerticalMergedNext() *tableCell {
	dc.verticalMerged = "rg"
	return dc
}

// func (dc TableCell) getVerticalMergedProperty() string {
// 	return dc.verticalMerged
// }

// SetMarginLeft function sets this cell's left margin
func (dc *tableCell) SetMarginLeft(value int) *tableCell {
	dc.marginLeft = value
	return dc
}

// SetMarginRight function sets this cell's right margin
func (dc *tableCell) SetMarginRight(value int) *tableCell {
	dc.marginRight = value
	return dc
}

// SetMarginTop function sets this cell's top margin
func (dc *tableCell) SetMarginTop(value int) *tableCell {
	dc.marginTop = value
	return dc
}

// SetMarginBottom function sets this cell's bottom margin
func (dc *tableCell) SetMarginBottom(value int) *tableCell {
	dc.marginBottom = value
	return dc
}

// SetVAlign sets align
func (dc *tableCell) SetVAlign(valign string) *tableCell {
	for _, i := range []string{VAlignBottom, VAlignMiddle, VAlignTop} {
		if valign == i {
			dc.vTextAlign = i
		}
	}
	return dc
}

// SetBorderColor function sets cell's border color
func (dc *tableCell) SetBorderColor(color string) *tableCell {
	dc.borderColor = color
	return dc
}
