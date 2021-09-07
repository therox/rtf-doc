package rtfdoc_test

import (
	"fmt"

	rtfdoc "github.com/therox/rtf-doc"
)

func ExampleDocument() {
	doc := rtfdoc.NewDocument()
	doc.AddParagraph().SetAlign(rtfdoc.AlignCenter).AddText("Hello, world!", 14, rtfdoc.FontTimesNewRoman, rtfdoc.ColorAqua)
	bs := doc.Export()
	fmt.Println(string(bs))
}
