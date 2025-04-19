package iteration

import "strings"

func Repeat(character string, count uint8) string {
	var repeated strings.Builder
	for range count {
		repeated.WriteString(character)
	}
	return repeated.String()
}