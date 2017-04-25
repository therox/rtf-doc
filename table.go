package rtfdoc

import "fmt"

func getDefaultTableProperties() tableProperties {
	tp := tableProperties{
		align: "c",
	}
	tp.SetTableMargins(100, 100, 100, 100)
	return tp
}

func (tp *tableProperties) SetTableMargins(left, top, right, bottom int) {
	margins := ""
	if left != 0 {
		margins += fmt.Sprintf(" \\trpaddl%d", left)
	}
	if top != 0 {
		margins += fmt.Sprintf(" \\trpaddt%d", top)
	}
	if right != 0 {
		margins += fmt.Sprintf(" \\trpaddr%d", right)
	}
	if bottom != 0 {
		margins += fmt.Sprintf(" \\trpaddb%d", bottom)
	}
	margins += " "
	tp.margins = margins
}

func (tp *tableProperties) getMargins() string {
	return tp.margins
}

// SetAlign sets table aligning (c/center, l/left, r/right)
func (tp *tableProperties) SetAlign(align string) {
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
		res += t.getMargins()
		res += tr.encode()
		res += "\n\\row}"
	}
	return res
}

// AddTableRow returns new table row instance
func (t *Table) AddTableRow() *TableRow {
	tr := TableRow{
		ft: t.ft,
		ct: t.ct,
	}
	t.data = append(t.data, &tr)
	return &tr
}

// AddCell add cell to TableRow
// func (tr *TableRow) AddCell(cell *TableCell) {
// 	tr.cells = append(tr.cells, cell)
// }

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
func (cp *cellProperties) SetProperties(cellWidth int, borders string) {
	cp.CellWidth = cellWidth
	cp.borders = borders
	return
}

// SetContent sets paragraph to datacell
// func (dc *DataCell) SetContent(c Paragraph) {
// 	dc.content = c
// }

// AddParagraph return cell's paragraph
func (dc *TableCell) AddParagraph() *Paragraph {
	p := Paragraph{
		align:  "l",
		indent: "\\fl360",
		ct:     dc.ct,
		ft:     dc.ft,
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

// SetBorders sets borders to datacell
func (dc *TableCell) SetBorders(left, top, right, bottom bool) {
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
func (dc *TableCell) SetVerticalMerged(isFirst, isNext bool) {
	if isFirst {
		dc.VerticalMerged = "\\clvmgf"
	}
	if isNext {
		dc.VerticalMerged = "\\clvmrg"
	}
}

func (dc TableCell) getVerticalMergedProperty() string {
	return dc.VerticalMerged
}

// SetCellMargins sets cell margins
func (dc *TableCell) SetCellMargins(left, top, right, bottom int) {
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
}

func (dc TableCell) getCellMargins() string {
	return dc.margins
}

// SetVAlign sets align (c/center, t/top, b/bottom)
func (dc *TableCell) SetVAlign(valign string) {
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
}

func (dc TableCell) getCellTextVAlign() string {
	return dc.vTextAlign
}

// SetWidth sets width of Table
func (t *Table) SetWidth(width int) {
	t.width = width
}
