package main

func converToRomanNumeral(num int) string {
	switch num {
	case 1:
		return "I"
	case 2:
		return "II"
	default:
		return ""
	}
}
