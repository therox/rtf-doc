package rtfdoc

// http://www.biblioscape.com/rtf15_spec.htm#Heading2

type DocumentItem interface {
	Compose() string
}

type CellItem interface {
	InCell()
}

type Header struct {
	Version string // Версия RTF, по-умолчанию, 1.5
	CharSet string // кодировка. Возможные варианты: ansi, mac, pc, pca
	Deff    string
	FontTBL FontTable
	//FileTBL    string
	ColorTBL ColorTable // Основные цветовые схемы. обращение в документе к ним с помощью управляющих слов \cfN, где N - порядковый номер цветовой схемы.
	//StyleSheet string
	//ListTables string
	//RevTBL     string
}
type Color struct {
	Red   int
	Green int
	Blue  int
	Code  string
}

type Document struct {
	Header
	orientation string
	DocumentSettings
	Content []DocumentItem
}

type DocumentSettings struct {
	PageSize Size
	Margins
}

type ColorTable []Color

type Font struct {
	Family  string // nil, roman, swiss, modern, script, decor, tech, bidi
	Charset int    // Specifies the character set of a font in the font table. Values for N are defined by Windows header files, and in the file RTFDEFS.H accompanying this document.
	Prq     int    // Specifies the pitch of a font in the font table.
	Name    string
	Code    string
}

type FontTable []Font

type Size struct {
	width  int
	height int
}

type Margins struct {
	left   int
	right  int
	top    int
	bottom int
}

//=================Table=======
type TableCell interface {
	cellCompose() string
	getCellWidth() int
	getBorders() string
	getVerticalMergedProperty() string
	getCellMargins() string
}

// Table - структура с таблицей
type Table struct {
	Data []TableRow
	TableProperties
}

type TableProperties struct {
	align   string
	margins string
}

type VerticalMerged struct {
	code string
}

type CellProperties struct {
	borders   string
	CellWidth int
	VerticalMerged
	margins string
}

// DataCell - структура ячейки таблицы с данными
type DataCell struct {
	Cell
}

type TableRow []TableCell

type Cell struct {
	CellProperties
	content Paragraph
}

// ============End of Table structs===========

type Paragraph struct {
	align   string
	indent  string
	content []Text
}

type Text struct {
	fontSize  int
	fontCode  int //code for font in font Table
	colorCode int
	italic    bool
	bold      bool
	underline bool
	super     bool
	sub       bool
	scaps     bool
	strike    bool
	text      string
}
