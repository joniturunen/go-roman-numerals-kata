package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestUserInput(t *testing.T) {
	t.Run("Get user input", func(t *testing.T) {
		input := fmt.Sprintln(10)
		rdr := strings.NewReader(input)
		got := getUserInput(rdr)
		expected := 10

		if got != expected {
			t.Errorf("got %d want %d", got, expected)
		}
	})
}

func TestArabicNumeralsToRomanNumerals(t *testing.T) {
	t.Run("Convert 1 to I", func(t *testing.T) {
		got := converToRomanNumeral(1)
		expected := "I"

		if got != expected {
			t.Errorf("got %q want %q", got, expected)
		}
	})
	t.Run("Convert 2 to II", func(t *testing.T) {
		got := converToRomanNumeral(2)
		expected := "II"

		if got != expected {
			t.Errorf("got %q want %q", got, expected)
		}
	})
}
