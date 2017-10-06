package rtfdoc

import "fmt"

func (t *Table) SetMarginLeft(value int) *Table {
	t.marginLeft = value
	return t
}

func (t *Table) SetMarginRight(value int) *Table {
	t.marginRight = value
	return t
}

func (t *Table) SetMarginTop(value int) *Table {
	t.marginTop = value
	return t
}

func (t *Table) SetMarginBottom(value int) *Table {
	t.marginBottom = value
	//tp.margins += fmt.Sprintf(" \\trpaddb%d", value)
	return t
}

//func (tp *tableProperties) getMargins() string {
//	return tp.margins
//}

// SetAlign sets table aligning (c/center, l/left, r/right)
func (t *Table) SetAlign(align string) *Table {
	for _, i := range []string{ALIGN_CENTER, ALIGN_LEFT, ALIGN_RIGHT} {
		if i == align {
			t.align = i
		}
	}
	return t
}

func (t *Table) GetAlign() string {
	return t.align
}

// AddTable returns Table instance
func (doc *Document) AddTable() *Table {
	t := Table{
		align: ALIGN_CENTER,
	}
	t.SetMarginLeft(100).SetMarginRight(100).SetMarginTop(100).SetMarginBottom(100)

	t.ct = doc.ct
	t.ft = doc.ft
	t.SetBorderLeft(true).
		SetBorderRight(true).
		SetBorderTop(true).
		SetBorderBottom(true).
		SetBorderStyle(BORDER_SINGLE_THICKNESS).
		SetBorderColor(COLOR_BLACK).
		SetBorderWidth(15)
	doc.content = append(doc.content, &t)
	return &t
}

func (t Table) compose() string {
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
func (t *Table) AddTableRow() *TableRow {
	tr := TableRow{
		generalSettings: generalSettings{
			ft: t.ft,
			ct: t.ct,
		},
	}
	tr.SetBorderLeft(t.borderLeft).
		SetBorderRight(t.borderRight).
		SetBorderTop(t.borderTop).
		SetBorderBottom(t.borderBottom).
		SetBorderStyle(t.borderStyle).
		SetBorderColor(t.borderColor).
		SetBorderWidth(t.borderWidth)
	t.data = append(t.data, &tr)
	return &tr
}

func (t *Table) SetBorderLeft(isBorder bool) *Table {
	t.borderLeft = isBorder
	return t
}
func (t *Table) SetBorderRight(isBorder bool) *Table {
	t.borderRight = isBorder
	return t
}
func (t *Table) SetBorderTop(isBorder bool) *Table {
	t.borderTop = isBorder
	return t
}
func (t *Table) SetBorderBottom(isBorder bool) *Table {
	t.borderBottom = isBorder
	return t
}

func (t *Table) SetBorderStyle(bStyle string) *Table {
	for _, i := range []string{
		BORDER_DASH_SMALL,
		BORDER_DASHED,
		BORDER_DOT_DASH,
		BORDER_DOT_DOT_DASH,
		BORDER_DOTTED,
		BORDER_DOUBLE,
		BORDER_DOUBLE_THICKNESS,
		BORDER_WAVY_DOUBLE,
		BORDER_EMBOSS,
		BORDER_ENGRAVE,
		BORDER_HAIRLINE,
		BORDER_INSET,
		BORDER_OUTSET,
		BORDER_SHADOWED,
		BORDER_SINGLE_THICKNESS,
		BORDER_STRIPPED,
		BORDER_THICK_THIN_LARGE,
		BORDER_THICK_THIN_MEDIUM,
		BORDER_THICK_THIN_SMALL,
		BORDER_THIN_THICK_LARGE,
		BORDER_THIN_THICK_MEDIUM,
		BORDER_THIN_THICK_SMALL,
		BORDER_THIN_THICK_THIN_LARGE,
		BORDER_THIN_THICK_THIN_MEDIUM,
		BORDER_TRIPLE,
		BORDER_WAVY,
	} {
		if bStyle == i {
			t.borderStyle = i
			break
		}
	}
	return t
}

func (t *Table) SetBorderColor(color string) *Table {
	t.borderColor = color
	return t
}
func (t *Table) SetBorderWidth(value int) *Table {
	t.borderWidth = value
	return t
}

func (tr *TableRow) SetBorderLeft(isBorder bool) *TableRow {
	tr.borderLeft = isBorder
	return tr
}
func (tr *TableRow) SetBorderRight(isBorder bool) *TableRow {
	tr.borderRight = isBorder
	return tr
}
func (tr *TableRow) SetBorderTop(isBorder bool) *TableRow {
	tr.borderTop = isBorder
	return tr
}
func (tr *TableRow) SetBorderBottom(isBorder bool) *TableRow {
	tr.borderBottom = isBorder
	return tr
}

func (tr *TableRow) SetBorderStyle(bStyle string) *TableRow {
	for _, i := range []string{
		BORDER_DASH_SMALL,
		BORDER_DASHED,
		BORDER_DOT_DASH,
		BORDER_DOT_DOT_DASH,
		BORDER_DOTTED,
		BORDER_DOUBLE,
		BORDER_DOUBLE_THICKNESS,
		BORDER_WAVY_DOUBLE,
		BORDER_EMBOSS,
		BORDER_ENGRAVE,
		BORDER_HAIRLINE,
		BORDER_INSET,
		BORDER_OUTSET,
		BORDER_SHADOWED,
		BORDER_SINGLE_THICKNESS,
		BORDER_STRIPPED,
		BORDER_THICK_THIN_LARGE,
		BORDER_THICK_THIN_MEDIUM,
		BORDER_THICK_THIN_SMALL,
		BORDER_THIN_THICK_LARGE,
		BORDER_THIN_THICK_MEDIUM,
		BORDER_THIN_THICK_SMALL,
		BORDER_THIN_THICK_THIN_LARGE,
		BORDER_THIN_THICK_THIN_MEDIUM,
		BORDER_TRIPLE,
		BORDER_WAVY,
	} {
		if bStyle == i {
			tr.borderStyle = i
			break
		}
	}
	return tr
}

func (tr *TableRow) SetBorderColor(color string) *TableRow {
	tr.borderColor = color
	return tr
}
func (tr *TableRow) SetBorderWidth(value int) *TableRow {
	tr.borderWidth = value
	return tr
}

func (tr *TableRow) encode() string {
	res := ""
	// Border settings
	bTempl := "\n \\trbrdr%s\\brdrw%d\\brdr%s"
	for c := range *tr.ct {
		if ((*tr.ct)[c]).name == tr.borderColor {
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

// AddDataCell returns new DataCell
func (tr *TableRow) AddDataCell(width int) *TableCell {
	dc := TableCell{
		cellWidth: width,
	}
	dc.ft = tr.ft
	dc.ct = tr.ct
	dc.SetBorderLeft(tr.borderLeft).
		SetBorderRight(tr.borderRight).
		SetBorderTop(tr.borderTop).
		SetBorderBottom(tr.borderBottom).
		SetBorderStyle(tr.borderStyle).
		SetBorderColor(tr.borderColor).
		SetBorderWidth(tr.borderWidth)

	tr.cells = append(tr.cells, &dc)
	return &dc
}

func (dc *TableCell) SetWidth(cellWidth int) *TableCell {
	dc.cellWidth = cellWidth
	return dc
}

// AddParagraph return cell's paragraph
func (dc *TableCell) AddParagraph() *Paragraph {
	p := Paragraph{
		isTable: true,
		align:   "l",
		indent:  "\\fl360",
		generalSettings: generalSettings{
			ct: dc.ct,
			ft: dc.ft,
		},
	}
	dc.content = append(dc.content, &p)
	return &p
}
func (dc TableCell) cellComposeProperties() string {
	res := ""
	// Тута свойства ячейки (границы, все дела...)
	bTempl := "\n \\clbrdr%s\\brdrw%d\\brdr%s"
	for c := range *dc.ct {
		if ((*dc.ct)[c]).name == dc.borderColor {
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

func (dc TableCell) cellComposeData() string {
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

func (dc TableCell) getCellWidth() int {
	return dc.cellWidth
}

// SetBorders sets borders to
// datacell

func (dc *TableCell) SetBorderLeft(value bool) *TableCell {
	dc.borderLeft = value
	return dc
}
func (dc *TableCell) SetBorderRight(value bool) *TableCell {
	dc.borderRight = value
	return dc
}
func (dc *TableCell) SetBorderTop(value bool) *TableCell {
	dc.borderTop = value
	return dc
}
func (dc *TableCell) SetBorderBottom(value bool) *TableCell {
	dc.borderBottom = value
	return dc
}

func (dc *TableCell) SetBorderWidth(value int) *TableCell {
	dc.borderWidth = value
	return dc
}

func (dc *TableCell) SetBorderStyle(bStyle string) *TableCell {
	for _, i := range []string{
		BORDER_DASH_SMALL,
		BORDER_DASHED,
		BORDER_DOT_DASH,
		BORDER_DOT_DOT_DASH,
		BORDER_DOTTED,
		BORDER_DOUBLE,
		BORDER_DOUBLE_THICKNESS,
		BORDER_WAVY_DOUBLE,
		BORDER_EMBOSS,
		BORDER_ENGRAVE,
		BORDER_HAIRLINE,
		BORDER_INSET,
		BORDER_OUTSET,
		BORDER_SHADOWED,
		BORDER_SINGLE_THICKNESS,
		BORDER_STRIPPED,
		BORDER_THICK_THIN_LARGE,
		BORDER_THICK_THIN_MEDIUM,
		BORDER_THICK_THIN_SMALL,
		BORDER_THIN_THICK_LARGE,
		BORDER_THIN_THICK_MEDIUM,
		BORDER_THIN_THICK_SMALL,
		BORDER_THIN_THICK_THIN_LARGE,
		BORDER_THIN_THICK_THIN_MEDIUM,
		BORDER_TRIPLE,
		BORDER_WAVY,
	} {
		if bStyle == i {
			dc.borderStyle = i
			break
		}
	}
	return dc
}

// func (dc *TableCell) SetBorders(left, top, right, bottom bool) *TableCell {
// 	dc.borderLeft = left
// 	dc.borderRight = right
// 	dc.borderTop = top
// 	dc.borderBottom = bottom
// 	// b := ""
// 	// bTemplStr := "\\clbrdr%s\\brdrw15\\brdrs"
// 	// if left {
// 	// 	b += fmt.Sprintf(bTemplStr, "l")
// 	// }
// 	// if top {
// 	// 	b += fmt.Sprintf(bTemplStr, "t")
// 	// }
// 	// if right {
// 	// 	b += fmt.Sprintf(bTemplStr, "r")
// 	// }
// 	// if bottom {
// 	// 	b += fmt.Sprintf(bTemplStr, "b")
// 	// }
// 	// dc.borders = b
// 	return dc
// }

// func (dc TableCell) getBorders() string {
// 	return dc.borders
// }

// GetTableCellWidthByRatio returns slice of cells width
func (t *Table) GetTableCellWidthByRatio(ratio ...float64) []int {

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
func (dc *TableCell) SetVerticalMergedFirst() *TableCell {
	dc.verticalMerged = "gf"
	return dc
}

// SetVerticalMergedFirst sets this cell to be not first cell in vertical merging.
func (dc *TableCell) SetVerticalMergedNext() *TableCell {
	dc.verticalMerged = "rg"
	return dc
}

func (dc TableCell) getVerticalMergedProperty() string {
	return dc.verticalMerged
}

func (dc *TableCell) SetMarginLeft(value int) *TableCell {
	dc.marginLeft = value
	return dc
}
func (dc *TableCell) SetMarginRight(value int) *TableCell {
	dc.marginRight = value
	return dc
}
func (dc *TableCell) SetMarginTop(value int) *TableCell {
	dc.marginTop = value
	return dc
}
func (dc *TableCell) SetMarginBottom(value int) *TableCell {
	dc.marginBottom = value
	return dc
}

// SetVAlign sets align (c/center, t/top, b/bottom)
func (dc *TableCell) SetVAlign(valign string) *TableCell {
	for _, i := range []string{VALIGN_BOTTOM, VALIGN_MIDDLE, VALIGN_TOP} {
		if valign == i {
			dc.vTextAlign = i
		}
	}
	return dc
}

func (dc *TableCell) SetBorderColor(color string) *TableCell {
	dc.borderColor = color
	return dc
}

// SetWidth sets width of Table
func (t *Table) SetWidth(width int) *Table {
	t.width = width
	return t
}
