package rtfdoc

import "fmt"

// AddParagraph return new instance of Paragraph
func (doc *document) AddParagraph() *paragraph {
	p := paragraph{
		align:   AlignCenter,
		indent:  "\\fl360",
		content: nil,
		generalSettings: generalSettings{
			colorTable: doc.colorTable,
			fontColor:  doc.fontColor,
		},
	}
	doc.content = append(doc.content, &p)
	return &p
}

func (par paragraph) compose() string {
	indentStr := fmt.Sprintf(" \\fi%d \\li%d \\ri%d",
		par.indentFirstLine,
		par.indentLeftIndent,
		par.indentRightIndent)
	res := fmt.Sprintf("\n{\\pard %s \\q%s", indentStr, par.align)
	if par.isTable {
		res += "\\intbl"
	}

	for _, c := range par.content {
		res += c.compose()
	}
	res += "\n\\par}"
	return res
}

// SetIndentFirstLine function sets first line indent in twips
func (par *paragraph) SetIndentFirstLine(value int) *paragraph {
	par.indentFirstLine = value
	return par
}

// SetIndentRight function sets right indent in twips
func (par *paragraph) SetIndentRight(value int) *paragraph {
	par.indentRightIndent = value
	return par
}

// SetIndentLeft function sets left indent in twips
func (par *paragraph) SetIndentLeft(value int) *paragraph {
	par.indentLeftIndent = value
	return par
}

// SetAlignt sets paragraph align (c/center, l/left, r/right, j/justify)
func (par *paragraph) SetAlignt(align string) *paragraph {
	for _, i := range []string{
		AlignCenter,
		AlignLeft,
		AlignRight,
		AlignJustify,
		AlignDistribute,
	} {
		if i == align {
			par.align = i
		}
	}

	return par
}
