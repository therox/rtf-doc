package rtfdoc

import "fmt"

func getDefaultTableProperties() tableProperties {
	tp := tableProperties{
		align: "c",
	}
	tp.SetTableMargins(100, 100, 100, 100)
	return tp
}

func (tp *tableProperties) SetTableMargins(left, top, right, bottom int) *tableProperties {
	tp.margins.left = left
	tp.margins.right = right
	tp.margins.top = top
	tp.margins.bottom = bottom
	//margins := ""
	//if left != 0 {
	//	margins += fmt.Sprintf(" \\trpaddl%d", left)
	//}
	//if top != 0 {
	//	margins += fmt.Sprintf(" \\trpaddt%d", top)
	//}
	//if right != 0 {
	//	margins += fmt.Sprintf(" \\trpaddr%d", right)
	//}
	//if bottom != 0 {
	//	margins += fmt.Sprintf(" \\trpaddb%d", bottom)
	//}
	//margins += " "
	//tp.margins = margins

	return tp
}

func (tp *tableProperties) SetLeftMargin(value int) *tableProperties {
	tp.margins.left = value
	//tp.margins += fmt.Sprintf(" \\trpaddl%d", value)
	return tp
}

func (tp *tableProperties) SetRightMargin(value int) *tableProperties {
	tp.margins.right = value
	//tp.margins += fmt.Sprintf(" \\trpaddr%d", value)
	return tp
}

func (tp *tableProperties) SetTopMargin(value int) *tableProperties {
	tp.margins.top = value
	//tp.margins += fmt.Sprintf(" \\trpaddt%d", value)
	return tp
}

func (tp *tableProperties) SetBottomMargin(value int) *tableProperties {
	tp.margins.bottom = value
	//tp.margins += fmt.Sprintf(" \\trpaddb%d", value)
	return tp
}

//func (tp *tableProperties) getMargins() string {
//	return tp.margins
//}

// SetAlign sets table aligning (c/center, l/left, r/right)
func (tp *tableProperties) SetAlign(align string) *tableProperties {
	switch align {
	case "c", "center":
		tp.align = "c"
	case "l", "left":
		tp.align = "l"
	case "r", "right":
		tp.align = "r"
	default:
		tp.align = ""
	}
	return tp
}

func (tp *tableProperties) GetAlign() string {
	return tp.align
}

// AddTable returns Table instance
func (doc *Document) AddTable() *Table {
	t := Table{tableProperties: getDefaultTableProperties()}
	t.ct = doc.ct
	t.ft = doc.ft
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
		// Margins
		//if t.margins.left != 0 {
		//	res += fmt.Sprintf(" \\trpaddl%d", t.margins.left)
		//}
		//if t.margins.right != 0 {
		//	res += fmt.Sprintf(" \\trpaddr%d", t.margins.right)
		//}
		//if t.margins.top != 0 {
		//	res += fmt.Sprintf(" \\trpaddt%d", t.margins.top)
		//}
		//if t.margins.bottom != 0 {
		//	res += fmt.Sprintf(" \\trpaddb%d", t.margins.bottom)
		//}
		res += fmt.Sprintf(" \\trpaddl%d \\trpaddr%d \\trpaddt%d \\trpaddb%d\n", t.margins.left, t.margins.right, t.margins.top, t.margins.bottom)
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
	t.data = append(t.data, &tr)
	return &tr
}

func (tr *TableRow) encode() string {
	res := ""
	if len(tr.cells) != 0 {
		cBegin := 0
		for _, tc := range tr.cells {
			cBegin += tc.getCellWidth()
			res += fmt.Sprintf("\n%s%s%s%s\\cellx%v", tc.getVerticalMergedProperty(), tc.getCellMargins(), tc.getBorders(), tc.getCellTextVAlign(), cBegin)

		}
		for _, tc := range tr.cells {
			res += tc.cellCompose()
		}
	}
	return res
}

// AddDataCell returns new DataCell
func (tr *TableRow) AddDataCell(width int) *TableCell {
	cp := cellProperties{}
	cp.CellWidth = width
	cp.ft = tr.ft
	cp.ct = tr.ct
	dc := TableCell{
		cellProperties: cp,
	}
	dc.SetBorders(true, true, true, true)
	tr.cells = append(tr.cells, &dc)
	return &dc
}

// SetProperties sets cell properties
func (cp *cellProperties) SetProperties(cellWidth int, borders string) *cellProperties {
	cp.CellWidth = cellWidth
	cp.borders = borders
	return cp
}

// AddParagraph return cell's paragraph
func (dc *TableCell) AddParagraph() *Paragraph {
	p := Paragraph{
		align:  "l",
		indent: "\\fl360",
		generalSettings: generalSettings{
			ct: dc.ct,
			ft: dc.ft,
		},
	}
	dc.content = append(dc.content, &p)
	return &p
}
func (dc TableCell) cellCompose() string {
	res := ""
	for _, p := range dc.content {
		res += fmt.Sprintf("\n\\pard\\intbl %s \n", p.compose())
	}
	res += "\\cell"

	return res
}

func (dc TableCell) getCellWidth() int {
	return dc.CellWidth
}

// SetBorders sets borders to
// datacell
func (dc *TableCell) SetBorders(left, top, right, bottom bool) *TableCell {
	b := ""
	bTemplStr := "\\clbrdr%s\\brdrw15\\brdrs"
	if left {
		b += fmt.Sprintf(bTemplStr, "l")
	}
	if top {
		b += fmt.Sprintf(bTemplStr, "t")
	}
	if right {
		b += fmt.Sprintf(bTemplStr, "r")
	}
	if bottom {
		b += fmt.Sprintf(bTemplStr, "b")
	}
	dc.borders = b
	return dc
}

func (dc TableCell) getBorders() string {
	return dc.borders
}

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

// SetVerticalMerged verticalMergedCell
func (dc *TableCell) SetVerticalMerged(isFirst, isNext bool) *TableCell {
	if isFirst {
		dc.VerticalMerged = "\\clvmgf"
	}
	if isNext {
		dc.VerticalMerged = "\\clvmrg"
	}
	return dc
}

func (dc TableCell) getVerticalMergedProperty() string {
	return dc.VerticalMerged
}

// SetCellMargins sets cell margins
func (dc *TableCell) SetCellMargins(left, top, right, bottom int) *TableCell {
	m := ""
	if left != 0 {
		m += fmt.Sprintf("\\clpadl%d", left)
	}
	if top != 0 {
		m += fmt.Sprintf("\\clpadt%d", top)
	}
	if right != 0 {
		m += fmt.Sprintf("\\clpadr%d", right)
	}
	if bottom != 0 {
		m += fmt.Sprintf("\\clpadb%d", bottom)
	}
	dc.margins = m

	return dc
}

func (dc TableCell) getCellMargins() string {
	return dc.margins
}

// SetVAlign sets align (c/center, t/top, b/bottom)
func (dc *TableCell) SetVAlign(valign string) *TableCell {
	switch valign {
	case "c", "center":
		dc.vTextAlign = "\\clvertalc"
	case "t", "top":
		dc.vTextAlign = "\\clvertalt"
	case "b", "bottom":
		dc.vTextAlign = "\\clvertalb"
	default:
		dc.vTextAlign = ""
	}
	return dc
}

func (dc TableCell) getCellTextVAlign() string {
	return dc.vTextAlign
}

// SetWidth sets width of Table
func (t *Table) SetWidth(width int) *Table {
	t.width = width
	return t
}
