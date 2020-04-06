package caesartools

import (
	"fmt"
	"math"
)

// Encrypt takes in plaintext and a shift and encrypts the text using the Caesar Cipher algorthm
func Encrypt(text string, shift int) string {

	// Mod shift to ensure -25 <= shift <= 25
	var offset int = int(math.Mod(float64(shift), 26))

	var ciphertext string = ""

	for i := 0; i < len(text); i++ {

		var newVal rune = rune(text[i]) + rune(offset)
		if rune(text[i]) >= 'A' && rune(text[i]) <= 'Z' {
			// Uppercase
			if offset > 0 && newVal > 'Z' {
				newVal = 'A' - 1 + (newVal - 'Z')
			} else if offset < 0 && newVal < 'A' {
				newVal = 'Z' + 1 + (newVal - 'A')
			}
			ciphertext += fmt.Sprintf("%c", newVal)
		} else if rune(text[i]) >= 'a' && rune(text[i]) <= 'z' {
			// Lowercase
			var newVal rune = rune(text[i]) + rune(offset)
			if offset > 0 && newVal > 'z' {
				newVal = 'a' - 1 + (newVal - 'z')
			} else if offset < 0 && newVal < 'a' {
				newVal = 'z' + 1 + (newVal - 'a')
			}
			ciphertext += fmt.Sprintf("%c", newVal)
		} else {
			ciphertext += fmt.Sprintf("%c", text[i])
		}
	}

	return ciphertext
}
