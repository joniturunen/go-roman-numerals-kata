package main

import (
	"fmt"
	"io"
)

func getUserInput(rdr io.Reader) int {
	var input int
	fmt.Print("Enter a number: ")
	fmt.Fscanln(rdr, &input)
	return input
}
