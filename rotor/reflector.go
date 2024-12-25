package rotor

/*
Contacts    = ABCDEFGHIJKLMNOPQRSTUVWXYZ
              ||||||||||||||||||||||||||
Reflector B = YRUHQSLDPXNGOKMIEBFZCWVJAT

Reflectors are just pairs of *contacts*.
If contact for 'B' (position 1, 'A' -> 0) is energized,
so is contact for 'R', and this goes both ways, an 'R'
contact energized is a 'B' contact energized.
*/

type Reflector struct {
	wiring [26]int
}

// Reflect from in position to out position
func (r *Reflector) Reflect(inPos int) (outPos int) {
	outPos = r.wiring[inPos]
}

var ReflectorB = &Reflector{
	wiring: [26]int{
		'Y' - 'A', 'R' - 'A', 'U' - 'A', 'H' - 'A', 'Q' - 'A',
		'S' - 'A', 'L' - 'A', 'D' - 'A', 'P' - 'A', 'X' - 'A',
		'N' - 'A', 'G' - 'A', 'O' - 'A', 'K' - 'A', 'M' - 'A',
		'I' - 'A', 'E' - 'A', 'B' - 'A', 'F' - 'A', 'Z' - 'A',
		'C' - 'A', 'W' - 'A', 'V' - 'A', 'J' - 'A', 'A' - 'A',
		'T' - 'A',
	},
}
