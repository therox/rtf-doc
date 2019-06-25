package rtfdoc

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf16"
)

// convertTextToUTF16 функция берет на вход строку и возвращает обратно строку,
// в которой русские буквы заменяются на эквиваленты в кодировке UTF-16
func convertNonASCIIToUTF16(text string) string {
	var res strings.Builder
	for _, r := range text {
		// if isCyrillicLetter(r) {
		if unicode.Is(unicode.Cyrillic, r) || r == '№' {
			res.WriteString(fmt.Sprintf("\\u%d\\'3f", utf16.Encode([]rune{r})[0]))
		} else {
			res.WriteString(string(r))
		}
	}
	return res.String()
}
