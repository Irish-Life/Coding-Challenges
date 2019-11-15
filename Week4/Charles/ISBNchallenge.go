package main


import (
	"fmt"
	"strconv"

)

func main() {

	var valid int64 = 8535902775
	//var inValid int64 = 1843369283
	reversed := reverseInts(valid)
	fmt.Println(testValidISBN(reversed))

}

func reverseInts(input int64) []int {
	strInt := intToString(input)
	ints := stringToIntArr(strInt)

	for i, j := 0, len(ints)-1; i < j; i, j = i+1, j-1 {
		ints[i], ints[j] = ints[j], ints[i]

	}

	return ints

}

func intToString(input int64) string {
	strInt := strconv.FormatInt(input, 10)
	return strInt

}

func stringToIntArr(input string) []int {
	var ints []int
	for _, c := range input {
		charInt := int(c - '0')
		ints = append(ints, charInt)

	}

	return ints
}

func testValidISBN(input []int) bool {
	fmt.Println(input)
	sum := 0
	for i, num := range input {
		sum += (i+1) * num
	}
	
	sumDiv := float64(sum) / 11.0

	if sumDiv == float64(int64(sumDiv)){
		return true
	} else {
		return false
	}


}
