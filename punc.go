package main

import (
	"strings"
)

func FixPunctuation(input string) string {
	var result strings.Builder
	length := len(input)
	spaceBeforePunctuation := false

	for i := 0; i < length; i++ {
		char := input[i]

		if char == ' ' {
			if i+1 < length && (input[i+1] == '.' || input[i+1] == ',' || input[i+1] == '!' || input[i+1] == '?') {
				spaceBeforePunctuation = true
				continue
			}
		}

		if spaceBeforePunctuation {
			result.WriteByte(input[i])
			spaceBeforePunctuation = false
			continue
		}

		if (char == '.' || char == ',' || char == '!' || char == '?') && i+1 < length && input[i+1] != ' ' {
			result.WriteByte(char)
			result.WriteByte(' ')
			continue
		}

		result.WriteByte(char)
	}

	return strings.TrimSpace(result.String())
}
//Corrige l'espacement autour de la ponctuation.