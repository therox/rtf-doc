package rtfdoc

func PageSize(width, height int) Size {
	return Size{width: width, height: height}
}

func PageMargins(left, right, top, bottom int) margins {
	return margins{left: left, right: right, top: top, bottom: bottom}
}
