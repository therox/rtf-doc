package rtfdoc

import "fmt"

func NewParagraph() Paragraph {
	return Paragraph{
		align:           "l",
		indentFirstLine: 720,
		content:         nil,
	}
}

func (par *Paragraph) AddText(t Text) {
	par.content = append(par.content, t)
}

func (par Paragraph) Compose() string {
	indent := ""
	if par.indentFirstLine != 0 {
		indent += fmt.Sprintf(" \\fi%d", par.indentFirstLine)
	}
	if par.indentBlockLeft != 0 {
		indent += fmt.Sprintf(" \\li%d", par.indentBlockLeft)
	}
	if par.indentBlockRight != 0 {
		indent += fmt.Sprintf(" \\ri%d", par.indentBlockRight)
	}

	res := fmt.Sprintf("\n{\\pard %s \\q%s", indent, par.align)
	for _, c := range par.content {
		res += c.Compose()
	}
	res += "\n\\par}"
	return res
}

func (par Paragraph) CellCompose() string {
	res := fmt.Sprintf("\n{\\pard \\q%s", par.align)
	for _, c := range par.content {
		res += c.Compose()
	}
	res += "\n}"

	return res
}

func (par *Paragraph) SetIndent(fl, li, ri int) {
	par.indentFirstLine = fl
	par.indentBlockLeft = li
	par.indentBlockRight = ri
}

func (par *Paragraph) SetAlignt(align string) {
	al := "l"
	switch align {
	case "c", "center":
		al = "c"
	case "l", "left":
		al = "l"
	case "r", "right":
		al = "r"
	case "j", "justify":
		al = "j"
	}
	par.align = al
}
