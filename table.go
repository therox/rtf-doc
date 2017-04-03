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
		TextProperties: TextProperties{},
		text:           "",
	}
}

func getCellWithProperties(cp CellProperties, tp TextProperties) Cell {
	return Cell{
		cp,
		tp,
		"",
	}
}

func getDefaultTable() Table {
	return Table{}
}
