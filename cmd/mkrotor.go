package main

import "fmt"

// Create a struct Rotor from in/out letter correspondence

type RotorRep struct {
	Name       string
	OutLetters string
}

/*
Entry = ABCDEFGHIJKLMNOPQRSTUVWXYZ (rotor right side)
        ||||||||||||||||||||||||||
I     = EKMFLGDQVZNTOWYHXUSPAIBRCJ
II    = AJDKSIRUXBLHWTMCQGZNPYFVOE
III   = BDFHJLCPRTXVZNYEIWGAKMUSQO
IV    = ESOVPZJAYQUIRHXLNFTGKDCMWB
V     = VZBRGITYUPSDNHLXAWMJQOFECK
*/

var rotors = []RotorRep{
	RotorRep{
		Name:       "I",
		OutLetters: "EKMFLGDQVZNTOWYHXUSPAIBRCJ",
	},
	RotorRep{
		Name:       "II",
		OutLetters: "AJDKSIRUXBLHWTMCQGZNPYFVOE",
	},
	RotorRep{
		Name:       "III",
		OutLetters: "BDFHJLCPRTXVZNYEIWGAKMUSQO",
	},
	RotorRep{
		Name:       "IV",
		OutLetters: "ESOVPZJAYQUIRHXLNFTGKDCMWB",
	},
	RotorRep{
		Name:       "V",
		OutLetters: "VZBRGITYUPSDNHLXAWMJQOFECK",
	},
}

func main() {

	for _, rotor := range rotors {
		fmt.Printf("var Rotor%s = &Rotor{\n", rotor.Name)
		fmt.Printf("\tEncode: [26]int{\n\t\t")
		for i, letter := range rotor.OutLetters {
			fmt.Printf("'%c' - 'A', ", letter)
			if (i % 5) == 4 {
				fmt.Print("\n\t\t")
			}
		}
		fmt.Println("\n\t},")

		fmt.Printf("\tInverse: [26]int{\n\t\t")
		var inv [26]int
		for i, letter := range rotor.OutLetters {
			inv[letter-'A'] = i + 'A'
		}
		for i, letter := range inv {
			fmt.Printf("'%c' - 'A', ", letter)
			if (i % 5) == 4 {
				fmt.Print("\n\t\t")
			}
		}
		fmt.Println("\n\t},")

		fmt.Println("\n}")
	}
}
