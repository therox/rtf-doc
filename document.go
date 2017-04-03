package rtfdoc

import "fmt"

func New() Document {
	return Document{
		orientation: "portrait",
		Header:      getDefaultHeader(),
		DocumentSettings: DocumentSettings{
			PageSize: PageSize(11952, 16848),
			Margins:  Margins{720, 720, 720, 720},
		},
		Content: nil,
	}
}

func (doc *Document) Compose() string {
	result := "{"
	result += doc.Header.Compose()
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
		result += fmt.Sprintf("\n%s", c.Compose())
	}
	result += "\n}"
	return result
}

func (doc *Document) AddContent(content DocumentItem) {

	doc.Content = append(doc.Content, content)
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

func (doc *Document) SetFontTable(ft FontTable) {
	doc.Header.FontTBL = ft
}
