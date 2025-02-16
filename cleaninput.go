package main

import (
	"strings"
)

func cleanInput(text string) []string {
	strSlc := strings.Split(strings.TrimSpace(text), " ")
	for i := range strSlc {
		strSlc[i] = strings.ToLower(strSlc[i])
	}
	return strSlc
}
