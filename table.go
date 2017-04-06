package rtfdoc

import "fmt"

func getDefaultCellProperties() CellProperties {
	return CellProperties{
		borders:   []string{"t", "b", "l", "r"},
		CellWidth: 1440,
	}
}

func getDefaultTableProperties() TableProperties {
	return TableProperties{
		align: "c",
	}
}

func NewTable() Table {
	return Table{TableProperties: getDefaultTableProperties()}
}

func (t *Table) AddRow(row TableRow) {
	t.Data = append(t.Data, row)
}

func (t Table) Compose() string {
	res := ""
	var align = ""
	if t.align != "" {
		align = fmt.Sprintf("\\trq%s", t.align)
	}
	for _, tr := range t.Data {
		res += fmt.Sprintf("\n{\\trowd %s", align)
		res += tr.Compose()
		res += "\n\\row}"
	}
	return res
}

func NewTableRow() TableRow {
	return TableRow{}
}
func (tr *TableRow) AddCell(cell TableCell) {
	*tr = append(*tr, cell)
}

func (tr TableRow) Compose() string {
	res := ""
	if len(tr) != 0 {
		cBegin := 0
		for _, dc := range tr {
			borders := ""

			if len(dc.getBorders()) > 0 {
				bTemplStr := "\\clbrdr%s\\brdrw15\\brdrs"
				for _, b := range dc.getBorders() {
					borders += fmt.Sprintf(bTemplStr, b)
				}
			}
			cBegin += dc.getCellWidth()
			res += fmt.Sprintf("\n%s%s\\cellx%v", dc.getVerticalMergedProperty(), borders, cBegin)

		}
		for _, dc := range tr {
			res += dc.CellCompose()
		}
	}
	return res
}

func NewDataCell(width int) DataCell {
	cp := getDefaultCellProperties()
	cp.CellWidth = width
	return DataCell{Cell{
		content:        Paragraph{},
		CellProperties: cp,
	}}
}
func NewDataCellWithProperties(cp CellProperties) DataCell {
	return DataCell{Cell{
		content:        Paragraph{},
		CellProperties: cp,
	}}
}

func (cp *CellProperties) SetProperties(cellWidth int, borders []string) {
	cp.CellWidth = cellWidth
	cp.borders = borders

	return
}

func (dc *DataCell) SetContent(c Paragraph) {
	dc.content = c
}

func (dc DataCell) CellCompose() string {
	res := fmt.Sprintf("\n\\pard\\intbl %s \\cell", dc.Cell.content.CellCompose())

	return res
}

func (dc DataCell) getCellWidth() int {
	return dc.CellWidth
}

func (dc DataCell) getBorders() []string {
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

func GetTableCellWidthByRatio(tableWidth int, ratio ...float64) []int {

	cellRatioSum := 0.0
	for _, cellRatio := range ratio {
		cellRatioSum += cellRatio
	}
	var cellWidth = make([]int, len(ratio))
	for i := range ratio {
		cellWidth[i] = int(ratio[i] * (float64(tableWidth) / cellRatioSum))
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
