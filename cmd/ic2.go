package main

/*
 * Calculate index of coincidence for the text in a file named on the
 * command line: ./ic2 <file name>
 *
 * Converts all input letters into upper-case. A "letter" here
 * is a Unicode code point.
 */

import (
	"fmt"
	"log"
	"os"
	"unicode"
)

func main() {
	buffer, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	var upperCaseLetters []rune

	for _, letter := range string(buffer) {
		if !unicode.IsLetter(letter) {
			continue
		}
		upperCaseLetters = append(upperCaseLetters, unicode.ToUpper(letter))
	}

	N, ic := indexOfCoincidence(upperCaseLetters)

	fmt.Printf("%d\t%.05f\n", N, ic)
}

// indexOfCoincidence examines the formal argument string
// (which should be uppercase unicode letters, 'A' - 'Z')
// counts the letters in the string, and returns a count of
// letters and the index of coincidence of that string
func indexOfCoincidence(buffer []rune) (int, float64) {
	N := 0
	var frequencies [26]int

	for _, r := range string(buffer) {
		frequencies[int(r-'A')]++
		N++
	}

	sum := 0
	for _, freq := range frequencies {
		sum += freq * (freq - 1)
	}

	ic := float64(sum) / float64(N*(N-1))

	return N, ic
}
