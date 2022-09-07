package main

import (
	"fmt"
	"io"
)

// Using reader interface allows us to test the function
func getUserInput(rdr io.Reader) int {
	var input int
	fmt.Print("Enter a number: ")
	fmt.Fscanln(rdr, &input)
	return input
}
