package main

import (
	"enigmalike/enigma"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	inFileName := flag.String("i", "", "input file name")
	first := flag.String("1", "I", "first rotor")
	second := flag.String("2", "II", "second rotor")
	third := flag.String("3", "III", "third rotor")
	settings := flag.String("S", "AAA", "initial rotor settings")
	plugs := flag.String("P", "", "comma-separated plugboard settings")
	flag.Parse()

	var cleartext string

	if *inFileName != "" {
		if buffer, err := os.ReadFile(*inFileName); err != nil {
			log.Fatal(err)
		} else {
			cleartext = string(buffer)
		}
	} else {
		// no input file, read a string from the command line
		if flag.NArg() > 0 {
			cleartext = flag.Arg(0)
		}
	}

	machine := enigma.NewMachine(*first, *second, *third)
	machine.SetRotors(*settings)

	if len(*plugs) > 0 {
		swaps := strings.Split(*plugs, ",")
		machine.Plugboard(swaps...)
	}

	for _, letter := range cleartext {
		if letter < 'A' || letter > 'Z' {
			continue
		}
		cipherLetter := machine.EncryptLetter(letter)
		fmt.Printf("%c", cipherLetter)
	}
	fmt.Println()
}
