package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"unicode"

	"enigmalike/rotor"
)

func main() {
	inFileName := flag.String("i", "", "input file name")
	verbose := flag.Bool("v", false, "verbose output")
	advance := flag.Bool("a", false, "don't advance first rotor")
	first := flag.String("1", "I", "first rotor")
	second := flag.String("2", "II", "second rotor")
	third := flag.String("3", "III", "third rotor")
	settings := flag.String("S", "AAA", "initial rotor settings")
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

	// set up rotors
	var rotor1, rotor2, rotor3 *rotor.Rotor
	var ok bool

	if rotor1 = rotor.ChooseRotor(*first); rotor1 == nil {
		log.Fatalf("no first rotor %q\n", *first)
	}
	if rotor2 = rotor.ChooseRotor(*second); rotor2 == nil {
		log.Fatalf("no second rotor %q\n", *second)
	}
	if rotor3 = rotor.ChooseRotor(*third); rotor3 == nil {
		log.Fatalf("no third rotor %q\n", *third)
	}

	for i, letter := range *settings {
		setting := int(unicode.ToUpper(letter))
		if setting < 'A' || setting > 'Z' {
			fmt.Fprintf(os.Stderr, "Ignoring bad setting %c\n", setting)
			continue
		}
		setting -= 'A'
		switch i {
		case 0:
			rotor1.Steps = setting
			if *verbose {
				fmt.Fprintf(os.Stderr, "rotor 1 (fast) setting %c\n", setting+'A')
			}
		case 1:
			rotor2.Steps = setting
			if *verbose {
				fmt.Fprintf(os.Stderr, "rotor 2 (medium) setting %c\n", setting+'A')
			}
		case 2:
			rotor3.Steps = setting
			if *verbose {
				fmt.Fprintf(os.Stderr, "rotor 3 (slow) setting %c\n", setting+'A')
			}
		default:
			fmt.Fprintf(os.Stderr, "unused rotor %d  setting %c\n", i+1, setting+'A')
		}
	}

	rotate := 1
	if *advance {
		rotate = 0
	}

	fout := &TraditionalOutput{}

	for _, r := range cleartext {
		if !unicode.IsLetter(r) {
			continue
		}

		if *verbose {
			fmt.Fprintf(os.Stderr, "input letter %c (%d)\n",
				r, (unicode.ToUpper(r) - 'A'))
		}

		// Give the input letter to the first rotor as a contact position,
		// which is 0 for 'A', 1 for 'B', 2 for 'C', etc etc
		outPos, carry := rotor1.CipherFwd(int(unicode.ToUpper(r)-'A'), rotate, *verbose)

		if *verbose {
			fmt.Fprintf(os.Stderr, "first rotor output letter %c (%d), carry %d\n",
				(outPos + 'A'), outPos, carry)
		}

		outPos, carry = rotor2.CipherFwd(outPos, carry, *verbose)

		if *verbose {
			fmt.Fprintf(os.Stderr, "second rotor output letter %c (%d), carry %d\n",
				(outPos + 'A'), outPos, carry)
		}

		outPos, carry = rotor3.CipherFwd(outPos, carry, *verbose)

		if *verbose {
			fmt.Fprintf(os.Stderr, "third rotor output letter %c (%d), carry %d\n",
				(outPos + 'A'), outPos, carry)
		}

		outPos = rotor.ReflectorB.Reflect(outPos)

		if *verbose {
			fmt.Fprintf(os.Stderr, "reflector B output letter %c (%d)\n",
				(outPos + 'A'), outPos)
		}

		outPos = rotor3.CipherBkwd(outPos, *verbose)

		if *verbose {
			fmt.Fprintf(os.Stderr, "backward through third rotor output letter %c (%d)\n",
				(outPos + 'A'), outPos)
		}

		outPos = rotor2.CipherBkwd(outPos, *verbose)

		if *verbose {
			fmt.Fprintf(os.Stderr, "backward through second rotor output letter %c (%d)\n",
				(outPos + 'A'), outPos)
		}

		outPos = rotor1.CipherBkwd(outPos, *verbose)

		if *verbose {
			fmt.Fprintf(os.Stderr, "backward through first rotor output letter %c (%d)\n",
				(outPos + 'A'), outPos)
		}

		fout.AddLetter(rune(outPos + 'A'))
	}

	fout.Output(os.Stdout)
}

type TraditionalOutput struct {
	letters []rune
}

func (t *TraditionalOutput) AddLetter(in rune) {
	t.letters = append(t.letters, in)
}

func (t *TraditionalOutput) Output(w io.Writer) {
	for i, l := range t.letters {
		fmt.Fprintf(w, "%c", l)
		if (i % 5) == 4 {
			fmt.Print(" ")
		}
		if (i % 25) == 24 {
			fmt.Println()
		}
	}
	fmt.Fprintf(w, "\n")
}
