package caesartools

import (
	"fmt"
	"math"
)

// Shift takes in text and a shift and encrypts/decrypts the text using the Caesar Cipher algorthm
func Shift(text string, shift int) string {

	// Mod shift to ensure -25 <= shift <= 25
	var offset int = int(math.Mod(float64(shift), 26))
	var output string = ""

	for i := 0; i < len(text); i++ {

		var newVal rune = rune(text[i]) + rune(offset)
		if rune(text[i]) >= 'A' && rune(text[i]) <= 'Z' {
			// Uppercase
			if offset > 0 && newVal > 'Z' {
				newVal = 'A' - 1 + (newVal - 'Z')
			} else if offset < 0 && newVal < 'A' {
				newVal = 'Z' + 1 + (newVal - 'A')
			}
			output += fmt.Sprintf("%c", newVal)
		} else if rune(text[i]) >= 'a' && rune(text[i]) <= 'z' {
			// Lowercase
			var newVal rune = rune(text[i]) + rune(offset)
			if offset > 0 && newVal > 'z' {
				newVal = 'a' - 1 + (newVal - 'z')
			} else if offset < 0 && newVal < 'a' {
				newVal = 'z' + 1 + (newVal - 'a')
			}
			output += fmt.Sprintf("%c", newVal)
		} else {
			output += fmt.Sprintf("%c", text[i])
		}
	}

	return output
}
