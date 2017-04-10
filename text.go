package rtfdoc

import "fmt"

func (text Text) Compose() string {
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

func NewText(text string, fontSize int, fontCode string, ft FontTable, colorCode string, ct ColorTable) Text {
	// Выясняем, какой шрифт имеет код fontcode
	fn := 0
	for i := range ft {
		if ft[i].Code == fontCode {
			// Присваиваем тексту порядковый номер шрифта
			fn = i
		}
	}
	// то-же самое с цветом
	fc := 0
	for i := range ct {
		if ct[i].Code == colorCode {
			// Присваиваем тексту порядковый номер шрифта
			fc = i + 1
		}
	}
	return Text{
		fontSize:  fontSize,
		fontCode:  fn,
		colorCode: fc,
		text:      text,
	}
}

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

func (text *Text) SetColor(colorCode string, ct ColorTable) {
	fc := 0
	for i := range ct {
		if ct[i].Code == colorCode {
			// Присваиваем тексту порядковый номер шрифта
			fc = i + 1
		}
	}
	text.colorCode = fc
}
