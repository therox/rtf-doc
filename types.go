package rtfdoc

import "image/color"

// http://www.biblioscape.com/rtf15_spec.htm#Heading2

// documentItem composing interface
type documentItem interface {
	compose() string
}

// cellItem cellizing interface
type cellItem interface {
	inCell()
}

type generalSettings struct {
	fontColor  *FontTable
	colorTable *ColorTable // Основные цветовые схемы. обращение в документе к ним с помощью управляющих слов \cfN, где N - порядковый номер цветовой схемы.
}

// Header - document header struct
type header struct {
	version string // RTF Version, default: 1.5
	charSet string // available options: ansi, mac, pc, pca
	deff    string
	generalSettings
	//FileTBL    string
	//StyleSheet string
	//ListTables string
	//RevTBL     string
}

// Color type for settings
type colorItem struct {
	rgbColor color.RGBA
	name     string
}

// Document - main document struct
type Document struct {
	header
	orientation string
	margins
	pageFormat string
	pagesize   size
	maxWidth   int
	content    []documentItem
}

// ColorTable defines color table
type ColorTable []colorItem

// font defines font struct
type font struct {
	family  string // nil, roman, swiss, modern, script, decor, tech, bidi
	charset int    // Specifies the character set of a font in the font table. Values for N are defined by Windows header files, and in the file RTFDEFS.H accompanying this document.
	prq     int    // Specifies the pitch of a font in the font table.
	name    string
	code    string
}

// FontTable defines font table
type FontTable []font

// Size struct
type size struct {
	width  int
	height int
}

// Table is a struct for Table.
type Table struct {
	width    int
	align    string
	docWidth int
	maxWidth int
	data     []*TableRow
	margins
	borders
	generalSettings
}

// TableCell defines cell properties
type TableCell struct {
	cellWidth      int
	verticalMerged string
	tableRowWidth  int
	maxWidth       int
	vTextAlign     string
	content        []*Paragraph
	borders
	margins
	generalSettings
	backgroundColor int
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
	cells      []*TableCell
	tableWidth int
	maxWidth   int
	borders
	generalSettings
}

// Main Picture struct
type Picture struct {
	format         string // EMF, PNG, JPEG
	paragraphWidth int
	maxWidth       int
	src            []byte
	scaleX         int
	scaleY         int
	cropL          int
	cropR          int
	cropT          int
	cropB          int
	height         int
	width          int
}

// ============End of Table structs===========

// Paragraph defines Paragraph instances
type Paragraph struct {
	isTable           bool
	align             string
	indent            string
	indentFirstLine   int
	indentLeftIndent  int
	indentRightIndent int
	content           []documentItem
	allowedWidth      int
	maxWidth          int
	generalSettings
}

// Text defines Text instances
type Text struct {
	fontSize      int
	fontCode      int //code for font in font Table
	colorCode     int
	isBold        bool
	isItalic      bool
	isUnderlining bool
	isScaps       bool
	isSuper       bool
	isSub         bool
	isStrike      bool
	emphasis      string
	content       string
	rotated       bool
	generalSettings
}

// Common paper orientation formats
const (
	OrientationPortrait  = "orientation_portrait"
	OrientationLandscape = "orientation_landscape"
)

// Commont paper formats
const (
	FormatA5 = "format_A5"
	FormatA4 = "format_A4"
	FormatA3 = "format_A3"
	FormatA2 = "format_A2"
)

// Aligning properties
const (
	AlignCenter     = "c"
	AlignLeft       = "l"
	AlignRight      = "r"
	AlignJustify    = "j"
	AlignDistribute = "d"

	VAlignTop     = "t"
	VAlignBottom  = "b"
	VAlignMiddle  = "c"
	VAlignJustify = "j"
)

// Common styles of border
const (
	BorderSingleThickness     = "s"
	BorderDoubleThickness     = "th"
	BorderShadowed            = "sh"
	BorderDouble              = "db"
	BorderDotted              = "dot"
	BorderDashed              = "dash"
	BorderHairline            = "hair"
	BorderInset               = "inset"
	BorderDashSmall           = "dashsm"
	BorderDotDash             = "dashd"
	BorderDotDotDash          = "dashdd"
	BorderOutset              = "outset"
	BorderTriple              = "triple"
	BorderThickThinSmall      = "tnthsg"
	BorderThinThickSmall      = "thtnsg"
	BorderThickThinMedium     = "tnthmg"
	BorderThinThickMedium     = "thtnmg"
	BorderThinThickThinMedium = "tnthtnmg"
	BorderThickThinLarge      = "tnthlg"
	BorderThinThickLarge      = "thtnlg"
	BorderThinThickThinLarge  = "tnthtnlg"
	BorderWavy                = "wavy"
	BorderWavyDouble          = "wavydb"
	BorderStripped            = "dashdotstr"
	BorderEmboss              = "emboss"
	BorderEngrave             = "engrave"
)

// Common image formats
const (
	ImageFormatJpeg = "jpeg"
	ImageFormatPng  = "png"
)

// List of common colors
const (
	ColorBlack   = "color_black"
	ColorBlue    = "color_blue"
	ColorAqua    = "color_aqua"
	ColorLime    = "color_lime"
	ColorGreen   = "color_green"
	ColorMagenta = "color_magenta"
	ColorRed     = "color_red"
	ColorYellow  = "color_yellow"
	ColorWhite   = "color_white"
	ColorNavy    = "color_navy"
	ColorTeal    = "color_teal"
	ColorPurple  = "color_purple"
	ColorMaroon  = "color_maroon"
	ColorOlive   = "color_olive"
	ColorGray    = "color_gray"
	ColorSilver  = "color_silver"
)

// Common fonts
const (
	FontTimesNewRoman = "font_times_new_roman"
	FontSymbol        = "font_symbol"
	FontArial         = "font_arial"
	FontComicSansMS   = "font_comic_sans_ms"
	FontCourierNew    = "font_courier_new"
)
