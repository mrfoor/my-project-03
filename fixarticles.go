package main

import (
	"strings"
)

func FixArticles(s string) string {
	words := strings.Fields(s)
	for i := 0; i < len(words)-1; i++ {
		if words[i] == "a" && strings.HasPrefix(strings.ToLower(words[i+1]), "a") {
			words[i] = "an"
		} else if words[i] == "an" && !strings.HasPrefix(strings.ToLower(words[i+1]), "a") {
			words[i] = "a"
		}
	}
	return strings.Join(words, " ")
}
