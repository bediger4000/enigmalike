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
	advance := flag.Bool("a", false, "advance first rotor")
	first := flag.String("1", "I", "first rotor")
	second := flag.String("2", "II", "second rotor")
	third := flag.String("3", "III", "third rotor")
	flag.Parse()

	var cleartext string

	if *inFileName != "" {
		if buffer, err := os.ReadFile(*inFileName); err != nil {
			log.Fatal(err)
		} else {
			cleartext = string(buffer)
		}
	} else {
		if flag.NArg() > 0 {
			cleartext = flag.Arg(0)
		}
	}

	rotor1 := rotor.Rotors[*first]
	rotor2 := rotor.Rotors[*second]
	rotor3 := rotor.Rotors[*third]

	rotate := 0
	if *advance {
		rotate = 1
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
