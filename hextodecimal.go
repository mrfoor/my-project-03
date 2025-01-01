package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Hextodecimal(s string) string {
	words := strings.Fields(s)
	for i := 0; i < len(words); i++ {
		if words[i] == "(hex)" {
			if i-1 >= 0 {
				hexStr := words[i-1]
				if num, err := strconv.ParseInt(hexStr, 16, 64); err == nil {
					words[i-1] = fmt.Sprint(num)
				}
			}
			words = append(words[:i], words[i+1:]...)
			i--
		}
	}
	return strings.Join(words, " ")
}
//Convertit les nombres hexadécimaux en décimaux.