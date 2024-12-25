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
	// step := flag.Bool("s", false, "show intermediate runes")
	verbose := flag.Bool("v", false, "verbose output")
	advance := flag.Bool("a", false, "advance first rotor")
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
		outPos, carry := rotor.Rotor1.CipherFwd(int(unicode.ToUpper(r)-'A'), rotate, *verbose)

		fout.AddLetter(rune(outPos + 'A'))

		if *verbose {
			fmt.Fprintf(os.Stderr, "output letter %c (%d), carry %d\n",
				(outPos + 'A'), outPos, carry)
		}
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
