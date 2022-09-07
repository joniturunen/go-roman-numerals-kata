package main

import (
	"fmt"
	"strings"
)

func converToRomanNumeral(num int) string {
	maxTrustedValue := 10000
	if num > maxTrustedValue {
		fmt.Printf("Numbers bigger than %v are not converted using large roman numerals.\n", maxTrustedValue)
	}
	numerals := []struct {
		arabic int
		roman  string
	}{
		{10000, "X̅"},
		{5000, "V̅"},
		{4000, "MV̅"},
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}
	var n strings.Builder
	for _, numeral := range numerals {
		for num >= numeral.arabic {
			n.WriteString(numeral.roman)
			num -= numeral.arabic
		}
	}
	return n.String()
}
