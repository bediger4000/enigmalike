package main

import (
	"enigmalike/enigma"
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

	inputText := convertBuffer(buffer)

	rotorNames := []string{"I", "II", "III", "IV", "V"}

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == j {
				continue
			}
			for k := 0; k < 5; k++ {
				if i == k || j == k {
					continue
				}
				machine := enigma.NewMachine(rotorNames[i], rotorNames[j], rotorNames[k])

				for ring1 := 0; ring1 < 26; ring1++ {
					for ring2 := 0; ring2 < 26; ring2++ {
						for ring3 := 0; ring3 < 26; ring3++ {
							ringSettings := fmt.Sprintf("%c%c%c", ring1+'A', ring2+'A', ring3+'A')
							machine.SetRotors(ringSettings)
							outputText := machine.EncryptBuffer(inputText)
							N, ic := indexOfCoincidence(outputText)
							fmt.Printf("%.05f\t%d\t%s\t%s\t%s\t%s\n", ic, N, rotorNames[i], rotorNames[j], rotorNames[k], ringSettings)
						}
					}
				}
			}
		}
	}
}

func convertBuffer(buffer []byte) []rune {
	var upperCaseLetters []rune

	for _, letter := range string(buffer) {
		if !unicode.IsLetter(letter) {
			continue
		}
		upperCaseLetters = append(upperCaseLetters, unicode.ToUpper(letter))
	}

	return upperCaseLetters
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
