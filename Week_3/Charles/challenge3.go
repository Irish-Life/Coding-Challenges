package main

import (
	"fmt"
)

func main() {
	htmlStr := "<h1> This is some <b> text </h1> incorrectly nested </br>"
	runes := strToRunes(htmlStr)
	tagRunes := findTags(runes)
	if testTagParity(tagRunes) {
		fmt.Println("Doing something")

	} else {
		fmt.Println("Unequal amount of opening and closing tags in string")
	}
}

func strToRunes(str string) []rune {
	return []rune(str)

}

func findTags(runes []rune) []rune {
	var tagRunes []rune
	for i := 0; i < len(runes); i++ {
		if string(runes[i]) == "<" || string(runes[i]) == ">" {
			tagRunes = append(tagRunes, runes[i])
		}
	}
	return tagRunes

}

// testTagParity takes in a slice containing opening and closing tags
// it tests if the number of opening tags is the same as closing
// if it returns True then the count is the same, if it returns false the count is not the same
// and the HTML is not correct.
func testTagParity(tagList []rune) bool {
	tagFrequency := make(map[rune]int)

	for _, item := range tagList {
		_, exist := tagFrequency[item]

		if exist {
			tagFrequency[item]++
		} else {
			tagFrequency[item] = 1
		}
	}

	var vals []int

	for k := range tagFrequency {
		vals = append(vals, tagFrequency[k])
	}

	sum := 0
	for _, v := range vals {
		sum += v
	}

	if float32(sum/len(vals)) == float32(vals[0]) {
		return true
	} else {
		return false
	}

}
