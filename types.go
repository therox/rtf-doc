package rtfdoc

import "image/color"

// http://www.biblioscape.com/rtf15_spec.htm#Heading2

// DocumentItem composing interface
type DocumentItem interface {
	compose() string
}

// CellItem cellizing interface
type CellItem interface {
	InCell()
}

type generalSettings struct {
	ft *FontTable
	ct *ColorTable // Основные цветовые схемы. обращение в документе к ним с помощью управляющих слов \cfN, где N - порядковый номер цветовой схемы.
}

// Header - document header struct
type header struct {
	version string // RTF Version, default: 1.5
	charSet string // available options: ansi, mac, pc, pca
	Deff    string
	generalSettings
	//FileTBL    string
	//StyleSheet string
	//ListTables string
	//RevTBL     string
}

// Color type for settings
type Color struct {
	color color.RGBA
	name  string
}

// Document - main document struct
type Document struct {
	header
	orientation string
	documentSettings
	content []DocumentItem
}

// DocumentSettings - struct for document settings
type documentSettings struct {
	pageFormat string
	pagesize   size
	margins
}

// ColorTable defines color table
type ColorTable []Color

// Font defines font struct
type Font struct {
	Family  string // nil, roman, swiss, modern, script, decor, tech, bidi
	Charset int    // Specifies the character set of a font in the font table. Values for N are defined by Windows header files, and in the file RTFDEFS.H accompanying this document.
	Prq     int    // Specifies the pitch of a font in the font table.
	Name    string
	Code    string
}

// FontTable defines font table
type FontTable []Font

// Size struct
type size struct {
	width  int
	height int
}

type margins struct {
	left   int
	right  int
	top    int
	bottom int
}

// Table is a struct for table.
type Table struct {
	data []*TableRow
	tableProperties
}

type tableProperties struct {
	width int
	align string
	margins
	generalSettings
}

// CellProperties define cell properties struct
type cellProperties struct {
	borders        string
	CellWidth      int
	VerticalMerged string
	margins        string
	vTextAlign     string
	generalSettings
}

// TableCell defines cell properties
type TableCell struct {
	cellProperties
	content []*Paragraph
}

// TableRow definces Table Row struct
type TableRow struct {
	cells []*TableCell
	generalSettings
}

// ============End of Table structs===========

// Paragraph defines paragraph instances
type Paragraph struct {
	align   string
	indent  string
	content []DocumentItem
	generalSettings
}

// Text defines Text instances
type Text struct {
	fontSize  int
	fontCode  int //code for font in font Table
	colorCode int
	emphasis  string
	text      string
	generalSettings
}
