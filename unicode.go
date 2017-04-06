package rtfdoc

import (
	"fmt"
	"unicode/utf16"
)

func isCyrillicLetter(l rune) bool {

	for _, rl := range "абвгдеёжзийклмнопрстуфхцчшщъыьэюяАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ№" {
		if l == rl {
			return true
		}
	}

	return false
}

// convertTextToUTF16 функция берет на вход строку и возвращает обратно строку,
// в которой русские буквы заменяются на эквиваленты в кодировке UTF-16
func convertCyrillicToUTF16(text string) string {
	res := ""
	for _, r := range text {
		if isCyrillicLetter(r) {
			res += fmt.Sprintf("\\u%d\\'3f", utf16.Encode([]rune{r})[0])
		} else {
			res += string(r)
		}
	}
	return res
}
