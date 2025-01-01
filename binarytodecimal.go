package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Binarytodecimal(s string) string {
	words := strings.Fields(s)
	for i := 0; i < len(words); i++ {
		if words[i] == "(bin)" {
			if i-1 >= 0 {
				binaryStr := words[i-1]
				if num, err := strconv.ParseInt(binaryStr, 2, 64); err == nil {
					words[i-1] = fmt.Sprint(num)
				}
			}
			words = append(words[:i], words[i+1:]...)
			i--
		}
	}
	return strings.Join(words, " ")
}
// Convertit les nombres binaires en décimaux.