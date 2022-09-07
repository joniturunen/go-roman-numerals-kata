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
	t.Run("Convert 5 to V", func(t *testing.T) {
		got := converToRomanNumeral(5)
		expected := "V"

		if got != expected {
			t.Errorf("got %q want %q", got, expected)
		}
	})
	t.Run("Convert 9 to IX", func(t *testing.T) {
		got := converToRomanNumeral(9)
		expected := "IX"

		if got != expected {
			t.Errorf("got %q want %q", got, expected)
		}
	})
	t.Run("Convert 13 to XII", func(t *testing.T) {
		got := converToRomanNumeral(13)
		expected := "XIII"

		if got != expected {
			t.Errorf("got %q want %q", got, expected)
		}
	})
	t.Run("Convert 3345 to MMMCCCXLV", func(t *testing.T) {
		got := converToRomanNumeral(3345)
		expected := "MMMCCCXLV"

		if got != expected {
			t.Errorf("got %q want %q", got, expected)
		}
	})
	t.Run("Convert 15000 to X̅V", func(t *testing.T) {
		got := converToRomanNumeral(15000)
		expected := "X̅V̅"

		if got != expected {
			t.Errorf("got %q want %q", got, expected)
		}
	})
}
