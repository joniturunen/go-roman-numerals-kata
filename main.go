package main

import "os"

func main() {
	rdr := os.Stdin
	inputStr := getUserInput(rdr)
	romanNumStr := converToRomanNumeral(inputStr)
	println(romanNumStr)
}
