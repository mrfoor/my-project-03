package main

import (
	"fmt"
	"strings"
)

func ProcessCapitalize(s string) string {
	words := strings.Fields(s)
	for i := 0; i < len(words); i++ {
		if strings.HasPrefix(words[i], "(cap,") {
			var n int
			fmt.Sscanf(words[i], "(cap,%d)", &n)
			start := i - n
			if start < 0 {
				start = 0
			}
			for j := start; j < i; j++ {
				words[j] = strings.Title(strings.ToLower(words[j]))
			}
			words = append(words[:i], words[i+1:]...)
			i--
		} else if words[i] == "(cap)" && i-1 >= 0 {
			words[i-1] = strings.Title(strings.ToLower(words[i-1]))
			words = append(words[:i], words[i+1:]...)
			i--
		}
	}
	return strings.Join(words, " ")
}
// Capitalise les mots précédents après (cap)