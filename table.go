package rtfdoc

func getDefaultCellProperties() CellProperties {
	return CellProperties{
		WithBorder:   true,
		BorderLeft:   true,
		BorderRight:  true,
		BorderTop:    true,
		BorderBottom: true,
	}
}

func getDefaultEmptyCell() Cell {
	return Cell{
		CellProperties: getDefaultCellProperties(),
		content:        Paragraph{},
	}
}

func getCellWithProperties(cp CellProperties) Cell {
	return Cell{
		CellProperties: cp,
		content:        Paragraph{},
	}
}

func NewTable() Table {
	return Table{}
}
