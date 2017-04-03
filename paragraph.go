package rtfdoc

import "fmt"

func NewParagraph() Paragraph {
	return Paragraph{
		align:   "l",
		content: nil,
	}
}

func (par *Paragraph) AddText(t Text) {
	par.content = append(par.content, t)
}

func (par Paragraph) Compose() string {
	res := fmt.Sprintf("\n{\\pard\\q%s", par.align)
	for _, c := range par.content {
		res += c.Compose()
	}
	res += "\n\\par}"
	return res
}
