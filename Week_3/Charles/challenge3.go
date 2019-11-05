package main

import (
	"fmt"
)

func main() {
	htmlStr := "<h1> This is some <b> text </h1> incorrectly nested </br>"
	runes := strToRunes(htmlStr)
	for i := 0; i < len(runes); i++ {
		if string(runes[i]) == "<" || string(runes[i]) == ">" {
			fmt.Printf("Rune %v is '%c'\n", i, runes[i])
		}
	}

}

func strToRunes(str string) []rune {
	return []rune(str)

}

func findTagIndices(runes []rune) {

}
