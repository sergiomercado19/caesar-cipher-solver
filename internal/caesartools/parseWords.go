package caesartools

import (
	"bufio"
	"log"
	"os"
)

// ParseWords puts the lists of words.txt into a map
func ParseWords(dictionary map[string]int, wordsPath string) {
	file, err := os.Open(wordsPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		dictionary[scanner.Text()] = 1
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
