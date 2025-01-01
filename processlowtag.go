package main

import (
	"strings"
)

func ProcessLowTag(s string) string {
	words := strings.Fields(s)
	for i := 1; i < len(words); i++ {
		if words[i] == "(low)" {
			words[i-1] = strings.ToLower(words[i-1])
			words = append(words[:i], words[i+1:]...)
			i--
		}
	}
	return strings.Join(words, " ")
}
// Met en minuscule le mot précédent après (low).