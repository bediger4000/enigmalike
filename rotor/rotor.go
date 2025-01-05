package rotor

import (
	"fmt"
	"os"
)

// Representation of a rotor in an electro-mechanical encryption device.
// Can shuffle 26 letters ('A' through 'Z' around).
// A "position" of a rotor consists of a count of "steps" from 0.
// The input to forward ciphering is the position (0 - 25) of the
// engergized contact on the rotor. The engergized contact is on the LHS of the rotor.
// The input to ciphering backward is the position (0 - 25) of the
// engergized contact on the rotor.
// Enciphering amounts to figuring out where rotationally the rotor is to the 0 ('A')
// contact, running the letter that contact corresponds to through the shuffling,
// the figuring out which output contact is energized. That's not the same as the letter
// output by shuffling because the rotor

type Rotor struct {
	Steps   int
	Encode  [26]int
	Inverse [26]int
}

func (r *Rotor) CipherFwd(inPos int, advance int, verbose bool) (outPos int, carry int) {
	r.Steps = ((r.Steps + advance) % 26)
	if r.Steps == 0 {
		// this rotor has been stepped 26 times, next rotor left should step
		carry = 1
	}

	// find index of this rotor that corresponds to inPos.
	// Since r.Steps is how far "ahead" this rotor is of
	// the 0 in position, the index calculated is which index
	// on this rotor corresponds to inPos
	internalPos := ((inPos + r.Steps) % 26)

	internalOutput := r.Encode[internalPos]

	outPos = internalOutput - r.Steps
	if outPos < 0 {
		outPos += 26
	}

	if verbose {
		fmt.Fprintf(os.Stderr, "CipherFwd, in pos %d, steps %d, internal pos %d, internal out %d, out pos %d\n",
			inPos, r.Steps, internalPos, internalOutput, outPos,
		)
	}

	return outPos, carry
}

// CipherBkwd takes the input *position* (which is fixed in space,
// figures out which rotor position that would match,
// runs the input position through the rotor shuffling backwards,
// and returns the output position
func (r *Rotor) CipherBkwd(inPos int, verbose bool) (outPos int) {
	// find index of this rotor that corresponds to inPos.
	// Since r.Steps is how far "ahead" this rotor is of
	// the 0 in position, the index calculated is which index
	// on this rotor corresponds to inPos
	internalPos := ((inPos + r.Steps) % 26) // LHS rotor contact

	internalOutput := r.Inverse[internalPos] // RHS rotor contact

	outPos = internalOutput - r.Steps
	if outPos < 0 {
		outPos += 26
	}

	return
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

var RotorI = &Rotor{
	Encode: [26]int{
		'E' - 'A', 'K' - 'A', 'M' - 'A', 'F' - 'A', 'L' - 'A',
		'G' - 'A', 'D' - 'A', 'Q' - 'A', 'V' - 'A', 'Z' - 'A',
		'N' - 'A', 'T' - 'A', 'O' - 'A', 'W' - 'A', 'Y' - 'A',
		'H' - 'A', 'X' - 'A', 'U' - 'A', 'S' - 'A', 'P' - 'A',
		'A' - 'A', 'I' - 'A', 'B' - 'A', 'R' - 'A', 'C' - 'A',
		'J' - 'A',
	},
	Inverse: [26]int{
		'U' - 'A', 'W' - 'A', 'Y' - 'A', 'G' - 'A', 'A' - 'A',
		'D' - 'A', 'F' - 'A', 'P' - 'A', 'V' - 'A', 'Z' - 'A',
		'B' - 'A', 'E' - 'A', 'C' - 'A', 'K' - 'A', 'M' - 'A',
		'T' - 'A', 'H' - 'A', 'X' - 'A', 'S' - 'A', 'L' - 'A',
		'R' - 'A', 'I' - 'A', 'N' - 'A', 'Q' - 'A', 'O' - 'A',
		'J' - 'A',
	},
}
var RotorII = &Rotor{
	Encode: [26]int{
		'A' - 'A', 'J' - 'A', 'D' - 'A', 'K' - 'A', 'S' - 'A',
		'I' - 'A', 'R' - 'A', 'U' - 'A', 'X' - 'A', 'B' - 'A',
		'L' - 'A', 'H' - 'A', 'W' - 'A', 'T' - 'A', 'M' - 'A',
		'C' - 'A', 'Q' - 'A', 'G' - 'A', 'Z' - 'A', 'N' - 'A',
		'P' - 'A', 'Y' - 'A', 'F' - 'A', 'V' - 'A', 'O' - 'A',
		'E' - 'A',
	},
	Inverse: [26]int{
		'A' - 'A', 'J' - 'A', 'P' - 'A', 'C' - 'A', 'Z' - 'A',
		'W' - 'A', 'R' - 'A', 'L' - 'A', 'F' - 'A', 'B' - 'A',
		'D' - 'A', 'K' - 'A', 'O' - 'A', 'T' - 'A', 'Y' - 'A',
		'U' - 'A', 'Q' - 'A', 'G' - 'A', 'E' - 'A', 'N' - 'A',
		'H' - 'A', 'X' - 'A', 'M' - 'A', 'I' - 'A', 'V' - 'A',
		'S' - 'A',
	},
}
var RotorIII = &Rotor{
	Encode: [26]int{
		'B' - 'A', 'D' - 'A', 'F' - 'A', 'H' - 'A', 'J' - 'A',
		'L' - 'A', 'C' - 'A', 'P' - 'A', 'R' - 'A', 'T' - 'A',
		'X' - 'A', 'V' - 'A', 'Z' - 'A', 'N' - 'A', 'Y' - 'A',
		'E' - 'A', 'I' - 'A', 'W' - 'A', 'G' - 'A', 'A' - 'A',
		'K' - 'A', 'M' - 'A', 'U' - 'A', 'S' - 'A', 'Q' - 'A',
		'O' - 'A',
	},
	Inverse: [26]int{
		'T' - 'A', 'A' - 'A', 'G' - 'A', 'B' - 'A', 'P' - 'A',
		'C' - 'A', 'S' - 'A', 'D' - 'A', 'Q' - 'A', 'E' - 'A',
		'U' - 'A', 'F' - 'A', 'V' - 'A', 'N' - 'A', 'Z' - 'A',
		'H' - 'A', 'Y' - 'A', 'I' - 'A', 'X' - 'A', 'J' - 'A',
		'W' - 'A', 'L' - 'A', 'R' - 'A', 'K' - 'A', 'O' - 'A',
		'M' - 'A',
	},
}
var RotorIV = &Rotor{
	Encode: [26]int{
		'E' - 'A', 'S' - 'A', 'O' - 'A', 'V' - 'A', 'P' - 'A',
		'Z' - 'A', 'J' - 'A', 'A' - 'A', 'Y' - 'A', 'Q' - 'A',
		'U' - 'A', 'I' - 'A', 'R' - 'A', 'H' - 'A', 'X' - 'A',
		'L' - 'A', 'N' - 'A', 'F' - 'A', 'T' - 'A', 'G' - 'A',
		'K' - 'A', 'D' - 'A', 'C' - 'A', 'M' - 'A', 'W' - 'A',
		'B' - 'A',
	},
	Inverse: [26]int{
		'H' - 'A', 'Z' - 'A', 'W' - 'A', 'V' - 'A', 'A' - 'A',
		'R' - 'A', 'T' - 'A', 'N' - 'A', 'L' - 'A', 'G' - 'A',
		'U' - 'A', 'P' - 'A', 'X' - 'A', 'Q' - 'A', 'C' - 'A',
		'E' - 'A', 'J' - 'A', 'M' - 'A', 'B' - 'A', 'S' - 'A',
		'K' - 'A', 'D' - 'A', 'Y' - 'A', 'O' - 'A', 'I' - 'A',
		'F' - 'A',
	},
}
var RotorV = &Rotor{
	Encode: [26]int{
		'V' - 'A', 'Z' - 'A', 'B' - 'A', 'R' - 'A', 'G' - 'A',
		'I' - 'A', 'T' - 'A', 'Y' - 'A', 'U' - 'A', 'P' - 'A',
		'S' - 'A', 'D' - 'A', 'N' - 'A', 'H' - 'A', 'L' - 'A',
		'X' - 'A', 'A' - 'A', 'W' - 'A', 'M' - 'A', 'J' - 'A',
		'Q' - 'A', 'O' - 'A', 'F' - 'A', 'E' - 'A', 'C' - 'A',
		'K' - 'A',
	},
	Inverse: [26]int{
		'Q' - 'A', 'C' - 'A', 'Y' - 'A', 'L' - 'A', 'X' - 'A',
		'W' - 'A', 'E' - 'A', 'N' - 'A', 'F' - 'A', 'T' - 'A',
		'Z' - 'A', 'O' - 'A', 'S' - 'A', 'M' - 'A', 'V' - 'A',
		'J' - 'A', 'U' - 'A', 'D' - 'A', 'K' - 'A', 'G' - 'A',
		'I' - 'A', 'A' - 'A', 'R' - 'A', 'P' - 'A', 'H' - 'A',
		'B' - 'A',
	},
}

var Rotors = map[string]*Rotor{
	"I":   RotorI,
	"II":  RotorII,
	"III": RotorIII,
	"IV":  RotorIV,
	"V":   RotorV,
}

// ChooseRotor returns a *copy* of a rotor it knows about,
// otherwise nil
func ChooseRotor(name string) *Rotor {
	if model, ok := Rotors[name]; ok {
		r := &Rotor{}
		_ = copy(r.Encode[:], model.Encode[:])
		_ = copy(r.Inverse[:], model.Inverse[:])
		return r
	}
	return nil
}
