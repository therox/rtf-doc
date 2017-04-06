package rtfdoc

import "fmt"

func NewParagraph() Paragraph {
	return Paragraph{
		align:   "l",
		indent:  "\\fl720",
		content: nil,
	}
}

func (par *Paragraph) AddText(t Text) {
	par.content = append(par.content, t)
}

func (par Paragraph) Compose() string {
	res := fmt.Sprintf("\n{\\pard %s \\q%s", par.indent, par.align)
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
	par.indent = ""

	if fl != 0 {
		par.indent += fmt.Sprintf(" \\fl%d", fl)
	}
	if li != 0 {
		par.indent += fmt.Sprintf(" \\li%d", fl)
	}
	if ri != 0 {
		par.indent += fmt.Sprintf(" \\ri%d", fl)
	}
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
