package caesartools

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

type counter struct {
	Value int
	Shift int
}

// Solve attempts to brute force all shift variations until it finds English words
func Solve(text string, dictionary map[string]int) int {

	// Create regex to filter out all non letters
	reg, err := regexp.Compile("[^a-zA-Z]+")
	if err != nil {
		log.Fatal(err)
	}

	// Initialize counters
	var maxMatches = counter{0, 0}

	// Loop through shifts
	for shift := 0; shift < 26; shift++ {

		var matches int = 0

		// Split input by words and loop through them
		words := strings.Split(text, " ")
		for _, word := range words {

			// Process word: remove non alpha chars and turn lower case
			var processedWord string = reg.ReplaceAllString(word, "")
			processedWord = strings.ToLower(processedWord)

			// Shift word
			var shiftedWord string = ""
			for i := 0; i < len(processedWord); i++ {
				var newVal rune = rune(processedWord[i]) - rune(shift)
				// Loop around
				if newVal < 'a' {
					newVal += rune(26)
				}
				shiftedWord += fmt.Sprintf("%c", newVal)
			}

			if _, ok := dictionary[shiftedWord]; ok {
				// Word found in dictionary
				matches++
			}
		}

		if matches > maxMatches.Value {
			maxMatches.Value = matches
			maxMatches.Shift = shift
		}

	}

	return maxMatches.Shift
}
