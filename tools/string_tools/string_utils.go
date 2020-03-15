package string_tools

import "strings"

func CleanWhiteSpaces(str string) string {
	ss := strings.FieldsFunc(str, func(c rune) bool {
		return c == ' '
	})

	return strings.Join(ss, " ")
}
