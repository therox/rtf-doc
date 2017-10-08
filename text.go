package rtfdoc

import "fmt"

func (text Text) compose() string {
	var res string
	var emphasisBegin string
	var EmphasisEnd string
	if text.emphasis != "" {
		emphasisBegin = "{" + text.getEmphasis() + " "
		EmphasisEnd = "}"
	}
	PreparedText := convertCyrillicToUTF16(text.text)

	res += fmt.Sprintf("\n\\fs%d\\f%d \\cf%d %s%s %s\\f0", text.fontSize*2, text.fontCode, text.colorCode, emphasisBegin, PreparedText, EmphasisEnd)
	return res
}

// AddText returns new text instance
func (p *Paragraph) AddText(text string, fontSize int, fontCode string, colorCode string) *Text {

	fn := 0
	for i, f := range *p.generalSettings.ft {
		if f.code == fontCode {

			fn = i
		}
	}

	fc := 0
	for i, c := range *p.generalSettings.ct {
		if c.name == colorCode {

			fc = i + 1
		}
	}
	txt := Text{
		fontSize:  fontSize,
		fontCode:  fn,
		colorCode: fc,
		text:      text,
		generalSettings: generalSettings{
			ct: p.ct,
			ft: p.ft,
		},
	}
	p.content = append(p.content, &txt)
	return &txt
}

//AddNewLine adds new line into paragraph text
func (p *Paragraph) AddNewLine() *Paragraph {
	txt := Text{
		text: "\\line",
	}
	p.content = append(p.content, &txt)
	return p
}

// SetBold function sets text to Bold
func (text *Text) SetBold() *Text {
	text.emphasis += " \\b"
	return text
}

// SetItalic function sets text to Italic
func (text *Text) SetItalic() *Text {
	text.emphasis += " \\i"
	return text
}

// SetUnderlining function sets text to Underlining
func (text *Text) SetUnderlining() *Text {
	text.emphasis += " \\ul"
	return text
}

// SetSuper function sets text to Super
func (text *Text) SetSuper() *Text {
	text.emphasis += " \\super"
	return text
}

// SetSub function sets text to Sub
func (text *Text) SetSub() *Text {
	text.emphasis += " \\sub"
	return text
}

// SetScaps function sets text to Scaps
func (text *Text) SetScaps() *Text {
	text.emphasis += " \\scaps"
	return text
}

// SetStrike function sets text to Strike
func (text *Text) SetStrike() *Text {
	text.emphasis += " \\strike"
	return text
}

func (text *Text) getEmphasis() string {
	return text.emphasis
}

// SetColor sets text color
func (text *Text) SetColor(colorCode string) *Text {
	for i := range *text.ct {
		if (*text.ct)[i].name == colorCode {
			// Присваиваем тексту порядковый номер шрифта
			text.colorCode = i + 1
		}
	}

	return text
}
