package main

import (
	"fmt"
)

func main() {
	//htmlStr := "<h1> This is some <b> text </h1> incorrectly nested </h1>"
	htmlStr := "<h1> This is some <b> text </b> correctly nested </h1>"
	//htmlStr := "<h1> This <div>is</div> some <b> text </b> correctly nested </h1>"


	runes := strToRunes(htmlStr)
	tagRunes, tagIndices := findTags(runes)
	if testTagParity(tagRunes) {
		testHTMLValidity(runes, tagIndices)

	} else {
		fmt.Println("Unequal amount of opening and closing tags in string")
	}
}

func strToRunes(str string) []rune {
	return []rune(str)
}

func findTags(runes []rune) ([]rune, []int) {
	var tagRunes []rune
	var tagIndices []int
	for i := 0; i < len(runes); i++ {
		if string(runes[i]) == "<" || string(runes[i]) == ">" {
			tagRunes = append(tagRunes, runes[i])
			tagIndices = append(tagIndices, i)
		}
	}
	return tagRunes, tagIndices

}

// testTagParity takes in a slice containing opening and closing tags
// it tests if the number of opening tags is the same as closing
// if it returns True then the count is the same, if it returns false the count is not the same
// and the HTML is not correct.
func testTagParity(tagList []rune) bool {
	tagFrequency := make(map[rune]int)

	// iterates through list of tag chars and creates
	// map/dict object of unique tags and their counts
	for _, item := range tagList {
		_, exist := tagFrequency[item]

		if exist {
			tagFrequency[item]++
		} else {
			tagFrequency[item] = 1
		}
	}

	// empty slice(array) of type int
	var vals []int

	for k := range tagFrequency {
		vals = append(vals, tagFrequency[k])
	}

	sum := 0
	for _, v := range vals {
		sum += v
	}
	// if the sum of the array is not equal to the first value
	// in the array, then we can say the numbers are not equal
	if float32(sum/len(vals)) == float32(vals[0]) {
		return true
	}

	return false

}

func testHTMLValidity(htmlRunes []rune, tagIndices []int) bool {
	if len(tagIndices) == 0 { //meaning there are only two tags left
		return true
	} else {
		reversedTagIndices := reverseIndex(tagIndices)
		fmt.Println(tagIndices)
		// fmt.Println(reversedTagIndices)
		startTag := getTagContents(htmlRunes, tagIndices[0], tagIndices[1])
		endTag := getTagContents(htmlRunes, reversedTagIndices[1], reversedTagIndices[0])

		if string(endTag[0]) == "/" { // if the end tag's first char is / then continue
			if len(endTag[1:]) != len(startTag) {
				fmt.Printf("Length of End tag %c (exluding '/') is not equal to length of start tag %c, HTML not valid\n", endTag, startTag)
				return false
			}
			for i := range startTag {
				if startTag[i] != endTag[i+1] {
					fmt.Printf("%c in start tag is not equal %c in end tag\n", startTag[i], endTag[i+1])
					return false
				} else {
					tagIndices = remove(tagIndices, 1)
					tagIndices = tagIndices[:len(tagIndices)-2]
					fmt.Println(tagIndices)
					testHTMLValidity(htmlRunes, tagIndices)
				}

			}

		} else {
			fmt.Printf("End tag %c is not a closing tag to %c, HTML not valid\n", endTag, startTag)
			return false
		}
	}

	return true
}

func reverseIndex(index []int) []int {
	reversed := []int{}
	for i := range index {
		n := index[len(index)-1-i]
		reversed = append(reversed, n)
	}
	return reversed
}

func getTagContents(htmlRunes []rune, start, end int) []rune {
	return htmlRunes[start+1 : end]
}

func remove(slice []int, i int) []int {
	return append(slice[:i-1], slice[i+1:]...) // i have no idea what ... does
}
