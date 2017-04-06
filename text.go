package rtfdoc

import "fmt"

func (text Text) Compose() string {
	var res string

	decor := ""
	if text.bold {
		decor += "\\b"
	}
	if text.italic {
		decor += "\\i"
	}
	if text.scaps {
		decor += "\\scaps"
	}
	if text.strike {
		decor += "\\strike"
	}
	if text.super {
		decor += "\\super"
	}
	if text.sub {
		decor += "\\sub"
	}

	var decorEnd string
	if decor != "" {
		decor = "{ " + decor
		decorEnd = "}"
	}
	PreparedText := convertCyrillicToUTF16(text.text)

	res += fmt.Sprintf("\n\\fs%d\\f%d \\cf%d %s %s %s\\f0", text.fontSize*2, text.fontCode, text.colorCode, decor, PreparedText, decorEnd)
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

func (text *Text) SetBold(bold bool) {
	text.bold = bold
}

func (text *Text) SetItalic(italic bool) {
	text.italic = italic
}

func (text *Text) SetUnderlining(ul bool) {
	text.underline = ul
}
func (text *Text) SetSuper(sup bool) {
	text.super = sup
}

func (text *Text) SetSub(sub bool) {
	text.sub = sub
}

func (text *Text) SetScaps(scaps bool) {
	text.scaps = scaps
}

func (text *Text) SetStrike(strike bool) {
	text.strike = strike
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
