package rtfdoc

import "fmt"

func New() Document {
	return Document{
		orientation: "portrait",
		Header:      getDefaultHeader(),
		//PageSize:    PageSize(16848, 11952),
		PageSize: PageSize(11952, 16848),
		Margins:  Margins{720, 720, 720, 720},
		Content:  []string{},
	}
}

func (doc *Document) String() string {
	result := "{"
	result += composeHeader(doc.Header)
	if doc.orientation != "" {
		result += fmt.Sprintf("\n\\%s", doc.orientation)
	}
	if doc.PageSize != (Size{}) {
		result += fmt.Sprintf("\n\\paperw%d\\paperh%d", doc.PageSize.width, doc.PageSize.height)
	}
	if doc.Margins != (Margins{}) {
		result += fmt.Sprintf("\n\\margl%d\\margr%d\\margt%d\\margb%d",
			doc.Margins.left,
			doc.Margins.right,
			doc.Margins.top,
			doc.Margins.bottom)
	}
	for _, c := range doc.Content {
		result += fmt.Sprintf("\n%s", c)
	}
	result += "\n}"
	return result
}

func (doc *Document) AddHeader(txt string, hAlign string, fontsize int) {
	alignSymbol := "c"
	switch hAlign {
	case "left":
		alignSymbol = "l"
	case "right":
		alignSymbol = "r"
	case "justified":
		alignSymbol = "j"
	}

	doc.Content = append(doc.Content, fmt.Sprintf("\n\\q%s\\fs%d %s", alignSymbol, fontsize*2, txt))
}

func (doc *Document) SetOrientation(orientation string) {
	if orientation == "landscape" {
		doc.orientation = "landscape"
		doc.PageSize = PageSize(16848, 11952)
	} else {
		doc.orientation = ""
		doc.PageSize = PageSize(11952, 16848)
	}
}
