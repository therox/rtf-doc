package rtfdoc

import "fmt"

func getDefaultTableProperties() TableProperties {
	tp := TableProperties{
		align: "c",
	}
	tp.SetTableMargins(100, 100, 100, 100)
	return tp
}

func (doc *Document) NewTable() *Table {
	t := Table{TableProperties: getDefaultTableProperties()}
	doc.AddContent(t)
	return &t
}

func (t *Table) AddRow(row TableRow) {
	t.Data = append(t.Data, row)
}

func (t *TableProperties) SetTableMargins(left, top, right, bottom int) {
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
	t.margins = margins
}

func (t *TableProperties) getMargins() string {
	return t.margins
}

func (t Table) compose() string {
	res := ""
	var align = ""
	if t.align != "" {
		align = fmt.Sprintf("\\trq%s", t.align)
	}
	for _, tr := range t.Data {
		res += fmt.Sprintf("\n{\\trowd %s", align)
		res += t.getMargins()
		res += tr.compose()
		res += "\n\\row}"
	}
	return res
}

func (table *Table) AddTableRow() *TableRow {
	tr := TableRow{}
	table.AddRow(tr)
	return &tr
}
func (tr *TableRow) addCell(cell TableCell) {
	tr.cells = append(tr.cells, cell)
}

func (tr TableRow) compose() string {
	res := ""
	if len(tr.cells) != 0 {
		cBegin := 0
		for _, dc := range tr.cells {
			cBegin += dc.getCellWidth()
			res += fmt.Sprintf("\n%s%s%s%s\\cellx%v", dc.getVerticalMergedProperty(), dc.getCellMargins(), dc.getBorders(), dc.getCellTextVAlign(), cBegin)

		}
		for _, dc := range tr.cells {
			res += dc.cellCompose()
		}
	}
	return res
}

func (tr *TableRow) NewDataCell(width int) *DataCell {
	cp := CellProperties{}
	cp.CellWidth = width
	dc := DataCell{
		Cell{
			content:        Paragraph{},
			CellProperties: cp,
		},
	}
	dc.SetBorders(true, true, true, true)
	tr.addCell(dc)
	return &dc
}
func NewDataCellWithProperties(cp CellProperties) DataCell {
	return DataCell{Cell{
		content:        Paragraph{},
		CellProperties: cp,
	}}
}

func (cp *CellProperties) SetProperties(cellWidth int, borders string) {
	cp.CellWidth = cellWidth
	cp.borders = borders
	return
}

func (dc *DataCell) SetContent(c Paragraph) {
	dc.content = c
}

func (dc DataCell) cellCompose() string {
	res := fmt.Sprintf("\n\\pard\\intbl %s \\cell", dc.Cell.content.cellCompose())

	return res
}

func (dc DataCell) getCellWidth() int {
	return dc.CellWidth
}

func (dc *DataCell) SetBorders(left, top, right, bottom bool) {
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

func (dc DataCell) getBorders() string {
	return dc.borders
}

func (tp *TableProperties) SetAlign(align string) {
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

func (tp *TableProperties) GetAlign() string {
	return tp.align
}

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

func (dc *DataCell) SetVerticalMerged(isFirst, isNext bool) {
	if isFirst {
		dc.VerticalMerged.code = "\\clvmgf"
	}
	if isNext {
		dc.VerticalMerged.code = "\\clvmrg"
	}
}

func (dc DataCell) getVerticalMergedProperty() string {
	return dc.VerticalMerged.code
}

func (dc *DataCell) SetCellMargins(left, top, right, bottom int) {
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

func (dc DataCell) getCellMargins() string {
	return dc.margins
}

func (dc *DataCell) SetVAlign(valign string) {
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

func (dc DataCell) getCellTextVAlign() string {
	return dc.vTextAlign
}

func (t *Table) SetWidth(width int) {
	t.width = width
}
