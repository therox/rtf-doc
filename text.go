package rtfdoc

import "fmt"

func (text Text) compose() string {
	var res string
	var emphasisBegin string
	var EmphasisEnd string
	if text.emphasis != "" {
		emphasisBegin = "{" + text.getEmphasis()
		EmphasisEnd = "}"
	}
	PreparedText := convertCyrillicToUTF16(text.text)

	res += fmt.Sprintf("\n\\fs%d\\f%d \\cf%d %s%s%s\\f0", text.fontSize*2, text.fontCode, text.colorCode, emphasisBegin, PreparedText, EmphasisEnd)
	return res
}

// AddText returns new text instance
func (p *Paragraph) AddText(text string, fontSize int, fontCode string, colorCode string) *Text {
	// Выясняем, какой шрифт имеет код fontcode
	fn := 0
	for i, f := range *p.generalSettings.ft {
		if f.Code == fontCode {
			// Присваиваем тексту порядковый номер шрифта
			fn = i
		}
	}
	// то-же самое с цветом
	fc := 0
	for i, c := range *p.generalSettings.ct {
		if c.name == colorCode {
			// Присваиваем тексту порядковый номер шрифта
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
func (p *Paragraph) AddNewLine() {
	txt := Text{
		text: "\\line",
	}
	p.content = append(p.content, &txt)
}

// SetEmphasis - sets text emphasis
func (text *Text) SetEmphasis(bold, italic, underlining, super, sub, scaps, strike bool) {
	text.emphasis = ""
	if bold {
		text.emphasis += " \\b"
	}
	if italic {
		text.emphasis += " \\i"
	}
	if underlining {
		text.emphasis += " \\ul"
	}
	if super {
		text.emphasis += " \\super"
	}
	if sub {
		text.emphasis += " \\sub"
	}
	if scaps {
		text.emphasis += " \\scaps"
	}
	if strike {
		text.emphasis += " \\strike"
	}
}

func (text *Text) getEmphasis() string {
	return text.emphasis
}

// SetColor sets text color
func (text *Text) SetColor(colorCode string, ct ColorTable) {
	fc := 0
	for i := range ct {
		if ct[i].name == colorCode {
			// Присваиваем тексту порядковый номер шрифта
			fc = i + 1
		}
	}
	text.colorCode = fc
}
