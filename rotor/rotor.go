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
		// this rotor has been stepped 26 times, so carry to the next rotor left
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
Entry = ABCDEFGHIJKLMNOPQRSTUVWXYZ
        ||||||||||||||||||||||||||
I     = EKMFLGDQVZNTOWYHXUSPAIBRCJ
*/
var Rotor1 = &Rotor{
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
