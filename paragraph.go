package rtfdoc

import "fmt"

// AddParagraph return new instance of Paragraph
func (doc *Document) AddParagraph() *Paragraph {
	p := Paragraph{
		align:   AlignCenter,
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

func (par Paragraph) compose() string {
	res := fmt.Sprintf("\n{\\pard %s \\q%s", par.indent, par.align)
	if par.isTable {
		res += "\\intbl"
	}

	for _, c := range par.content {
		res += c.compose()
	}
	res += "\n\\par}"
	return res
}

// SetIndent sets indent to paragraph (fl - first line indent, li - left ident, ri - right indent in tweeps)
func (par *Paragraph) SetIndent(fl, li, ri int) *Paragraph {
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
	return par
}

// SetAlignt sets paragraph align (c/center, l/left, r/right, j/justify)
func (par *Paragraph) SetAlignt(align string) *Paragraph {
	for _, i := range []string{AlignCenter, AlignLeft, AlignRight} {
		if i == align {
			par.align = i
		}
	}

	return par
}
