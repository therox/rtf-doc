package main

// http://www.biblioscape.com/rtf15_spec.htm#Heading2

type RTFHeader struct {
	Version    string // Версия RTF, по-умолчанию, 1.5
	CharSet    string // кодировка. Возможные варианты: ansi, mac, pc, pca
	Deff       string
	FontTBL    string
	FileTBL    string
	ColorTBL   RTFColorTBL // Основные цветовые схемы. обращение в документе к ним с помощью управляющих слов \cfN, где N - порядковый номер цветовой схемы.
	StyleSheet string
	ListTables string
	RevTBL     string
}
type RTFRGBColor struct {
	Red   int
	Green int
	Blue  int
}

type RTFDocument struct {
	RTFHeader
}
type RTFCode string

type RTFColorTBL []RTFRGBColor
