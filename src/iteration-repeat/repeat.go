package iteration

import "strings"

const repeatCount = 5

// Repeat retrieve a character and return the same repeated five times
func Repeat(character string, repeat int) string {
	var repeated string
	if repeat == 0 {
		for i := 0; i < repeatCount; i++ {
			repeated += character
		}
	} else {
		repeated = strings.Repeat(character, repeat)
	}
	return repeated
}
