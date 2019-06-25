package rtfdoc

import (
	"fmt"
	"strings"
)

func (text Text) compose() string {
	var res strings.Builder

	var emphTextSlice []string
	if text.isBold {
		emphTextSlice = append(emphTextSlice, "\\b")
	}
	if text.isItalic {
		emphTextSlice = append(emphTextSlice, "\\i")
	}
	if text.isScaps {
		emphTextSlice = append(emphTextSlice, "\\scaps")
	}
	if text.isStrike {
		emphTextSlice = append(emphTextSlice, "\\strike")
	}
	if text.isSub {
		emphTextSlice = append(emphTextSlice, "\\sub")
	}
	if text.isSuper {
		emphTextSlice = append(emphTextSlice, "\\super")
	}
	if text.isUnderlining {
		emphTextSlice = append(emphTextSlice, "\\ul")
	}
	if text.rotated {
		emphTextSlice = append(emphTextSlice, "\\horzvert0")
	}

	PreparedText := convertNonASCIIToUTF16(text.content)

	res.WriteString(fmt.Sprintf("\n\\fs%d\\f%d \\cf%d %s{%s}\\f0", text.fontSize*2, text.fontCode, text.colorCode, strings.Join(emphTextSlice, " "), PreparedText))
	return res.String()
}

// AddText returns new text instance
func (p *Paragraph) AddText(textStr string, fontSize int, fontCode string, colorCode string) *Text {

	fn := 0
	for i, f := range *p.generalSettings.fontColor {
		if f.code == fontCode {

			fn = i
		}
	}

	fc := 0
	for i, c := range *p.generalSettings.colorTable {
		if c.name == colorCode {

			fc = i + 1
		}
	}
	txt := Text{
		fontSize:  fontSize,
		fontCode:  fn,
		colorCode: fc,
		content:   textStr,
		generalSettings: generalSettings{
			colorTable: p.colorTable,
			fontColor:  p.fontColor,
		},
	}
	p.content = append(p.content, &txt)
	return &txt
}

//AddNewLine adds new line into Paragraph text
func (p *Paragraph) AddNewLine() *Paragraph {
	txt := Text{
		content: "\\line",
	}
	p.content = append(p.content, &txt)
	return p
}

// SetBold function sets text to Bold
func (text *Text) SetBold() *Text {
	text.isBold = true
	return text
}

// SetItalic function sets text to Italic
func (text *Text) SetItalic() *Text {
	text.isItalic = true
	return text
}

// SetUnderlining function sets text to Underlining
func (text *Text) SetUnderlining() *Text {
	text.isUnderlining = true
	return text
}

// SetSuper function sets text to Super
func (text *Text) SetSuper() *Text {
	text.isSuper = true
	return text
}

// SetSub function sets text to Sub
func (text *Text) SetSub() *Text {
	text.isSub = true
	return text
}

// SetScaps function sets text to Scaps
func (text *Text) SetScaps() *Text {
	text.isScaps = true
	return text
}

// SetStrike function sets text to Strike
func (text *Text) SetStrike() *Text {
	text.isStrike = true
	return text
}

// SetRotate function rotates Text so it flows in a direction opposite to that of the main document (Horizontal in vertical and vertical in horizontal)
func (text *Text) SetRotate() *Text {
	text.rotated = true
	return text
}

func (text *Text) getEmphasis() string {
	return text.emphasis
}

// SetColor sets text color
func (text *Text) SetColor(colorCode string) *Text {
	for i := range *text.colorTable {
		if (*text.colorTable)[i].name == colorCode {
			// Присваиваем тексту порядковый номер шрифта
			text.colorCode = i + 1
		}
	}

	return text
}
