package rtfdoc

// http://www.biblioscape.com/rtf15_spec.htm#Heading2

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
}

type Document struct {
	Header
	orientation string
	PageSize    Size
	Margins

	Content []string
}
type ColorTable []Color

type Font struct {
	Family  string // nil, roman, swiss, modern, script, decor, tech, bidi
	Charset int    // Specifies the character set of a font in the font table. Values for N are defined by Windows header files, and in the file RTFDEFS.H accompanying this document.
	Prq     int    // Specifies the pitch of a font in the font table.
	Name    string
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

// Table - структура с таблицей
type Table struct {
	HeaderRow []HeaderCell
	Data      []DataRow
}

type CellProperties struct {
	WithBorder   bool
	BorderLeft   bool
	BorderRight  bool
	BorderTop    bool
	BorderBottom bool
}

type TextProperties struct {
	Font *Font
}

// TableCell - структура ячейки таблицы с заголовком
type HeaderCell struct {
	Cell
}

// DataCell - структура ячейки таблицы с данными
type DataCell struct {
	Cell
}

type DataRow []DataCell

type Cell struct {
	CellProperties
	TextProperties
	text string
}

// ============End of Table structs===========
