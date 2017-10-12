package rtfdoc

func getPixelsFromTwips(value int) int {
	return int(value / 15)
}

func getTwipsFromPixels(value int) int {
	return value * 15
}
