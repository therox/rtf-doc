package rtfdoc

import "fmt"

// NewParagraph return new instance of Paragraph
// func NewParagraph() *Paragraph {
// 	return &Paragraph{
// 		align:   "l",
// 		indent:  "\\fl720",
// 		content: nil,
// 	}
// }

// AddParagraph return new instance of Paragraph
func (doc *Document) AddParagraph() *Paragraph {
	p := Paragraph{
		align:   "l",
		indent:  "\\fl360",
		content: nil,
		generalSettings: generalSettings{
			ct: doc.ct,
			ft: doc.ft,
		},
	}
	doc.content = append(doc.content, &p)
	return &p
}

// AddContent adds content
// func (par *Paragraph) AddContent(c Text) {
// 	par.content = append(par.content, c)
// }

func (par Paragraph) compose() string {
	res := fmt.Sprintf("\n{\\pard %s \\q%s", par.indent, par.align)

	for _, c := range par.content {
		res += c.compose()
	}
	res += "\n\\par}"
	return res
}

// func (par Paragraph) cellCompose() string {
// 	res := fmt.Sprintf("\n{\\pard %s \\q%s", par.indent, par.align)
// 	for _, c := range par.content {
// 		res += c.compose()
// 	}
// 	res += "\n}"

// 	return res
// }

// SetIndent sets indent to paragraph (fl - first line indent, li - left ident, ri - right indent in tweeps)
func (par *Paragraph) SetIndent(fl, li, ri int) {
	par.indent = ""

	if fl != 0 {
		par.indent += fmt.Sprintf(" \\fi%d", fl)
	}
	if li != 0 {
		par.indent += fmt.Sprintf(" \\li%d", fl)
	}
	if ri != 0 {
		par.indent += fmt.Sprintf(" \\ri%d", fl)
	}
}

// SetAlignt sets paragraph align (c/center, l/left, r/right, j/justify)
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
