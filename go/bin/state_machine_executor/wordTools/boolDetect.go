package wordTools

import (
	"strings"
	"unicode/utf8"
)

func BoolDetect(message string) (bool, bool) {
	message = strings.ToLower(message)

	if utf8.RuneCountInString(message) < 6 {
		if strings.Contains(message, "да") {
			return true, true
		} else if strings.Contains(message, "нет") {
			return false, true
		}
	}

	return false, false
}