package main

import "fmt"
import "rtf-doc"

func main() {
	//fmt.Println(string(composeHeader(getDefaultHeader())))
	// Создаем документ
	d := rtfdoc.New()
	// Настроить хедер
	d.Header.ColorTBL.Set(rtfdoc.Color{255, 0, 0})

	// ...
	d.AddHeader("Test123", "center", 48)
	d.SetOrientation("landscape")

	fmt.Println(d.String())
}
