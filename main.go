package main

import "os"

func main() {
	rdr := os.Stdin
	userInput := getUserInput(rdr)
	romanNumStr := converToRomanNumeral(userInput)
	println(romanNumStr)
}
