package prose

import "strings"

func JoinWithCommas(phrases []string) string {
	if len(phrases) == 0 {
		return ""
	} else if len(phrases) == 1 {
		return phrases[0]
	} else if len(phrases) == 2 {
		return phrases[0] + " and " + phrases[1]
	} else {
		// вырезем из среза все кроме последнего элемента
		// и преобразуем в строку разделенную запятыми
		result := strings.Join(phrases[:len(phrases)-1], ", ")
		// приклеиваем к строке слово and
		result += ", and "
		// вырезаем из среза последний элемент
		// и приклеиваем его к строке
		result += phrases[len(phrases)-1]
		return result
	}
}
