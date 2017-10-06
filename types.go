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
	pageFormat  string
	pagesize    size
	margins
	content []DocumentItem
}

// ColorTable defines color table
type ColorTable []Color

// Font defines font struct
type Font struct {
	family  string // nil, roman, swiss, modern, script, decor, tech, bidi
	charset int    // Specifies the character set of a font in the font table. Values for N are defined by Windows header files, and in the file RTFDEFS.H accompanying this document.
	prq     int    // Specifies the pitch of a font in the font table.
	name    string
	code    string
}

// FontTable defines font table
type FontTable []Font

// Size struct
type size struct {
	width  int
	height int
}

// Table is a struct for table.
type Table struct {
	width int
	align string
	margins
	borders
	generalSettings
	data []*TableRow
}

// TableCell defines cell properties
type TableCell struct {
	borders
	cellWidth      int
	verticalMerged string
	margins
	vTextAlign string
	generalSettings
	content []*Paragraph
}

type borders struct {
	borderLeft   bool
	borderRight  bool
	borderTop    bool
	borderBottom bool
	borderStyle  string
	borderWidth  int
	borderColor  string
}

type margins struct {
	marginLeft   int
	marginRight  int
	marginTop    int
	marginBottom int
}

// TableRow definces Table Row struct
type TableRow struct {
	cells []*TableCell
	borders
	generalSettings
}

// ============End of Table structs===========

// Paragraph defines paragraph instances
type Paragraph struct {
	isTable bool
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

const (
	ORIENTATION_PORTRAIT  = "orientation_portrait"
	ORIENTATION_LANDSCAPE = "orientation_landscape"

	FORMAT_A5 = "format_A5"
	FORMAT_A4 = "format_A4"
	FORMAT_A3 = "format_A3"
	FORMAT_A2 = "format_A2"

	ALIGN_CENTER  = "c"
	ALIGN_LEFT    = "l"
	ALIGN_RIGHT   = "r"
	ALIGN_JUSTIFY = "j"

	VALIGN_TOP    = "t"
	VALIGN_BOTTOM = "b"
	VALIGN_MIDDLE = "m"

	BORDER_SINGLE_THICKNESS       = "s"
	BORDER_DOUBLE_THICKNESS       = "th"
	BORDER_SHADOWED               = "sh"
	BORDER_DOUBLE                 = "db"
	BORDER_DOTTED                 = "dot"
	BORDER_DASHED                 = "dash"
	BORDER_HAIRLINE               = "hair"
	BORDER_INSET                  = "inset"
	BORDER_DASH_SMALL             = "dashsm"
	BORDER_DOT_DASH               = "dashd"
	BORDER_DOT_DOT_DASH           = "dashdd"
	BORDER_OUTSET                 = "outset"
	BORDER_TRIPLE                 = "triple"
	BORDER_THICK_THIN_SMALL       = "tnthsg"
	BORDER_THIN_THICK_SMALL       = "thtnsg"
	BORDER_THICK_THIN_MEDIUM      = "tnthmg"
	BORDER_THIN_THICK_MEDIUM      = "thtnmg"
	BORDER_THIN_THICK_THIN_MEDIUM = "tnthtnmg"
	BORDER_THICK_THIN_LARGE       = "tnthlg"
	BORDER_THIN_THICK_LARGE       = "thtnlg"
	BORDER_THIN_THICK_THIN_LARGE  = "tnthtnlg"
	BORDER_WAVY                   = "wavy"
	BORDER_WAVY_DOUBLE            = "wavydb"
	BORDER_STRIPPED               = "dashdotstr"
	BORDER_EMBOSS                 = "emboss"
	BORDER_ENGRAVE                = "engrave"
)
