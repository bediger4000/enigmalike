package enigma

import (
	"enigmalike/rotor"
	"log"
	"unicode"
)

type Machine struct {
	rotor1    *rotor.Rotor
	rotor2    *rotor.Rotor
	rotor3    *rotor.Rotor
	reflector *rotor.Reflector
}

// NewMachine arranges 3 rotors ("first" is leftmost), but doesn't set them
func NewMachine(first, second, third string) *Machine {
	// set up rotors
	var rotor1, rotor2, rotor3 *rotor.Rotor

	if rotor1 = rotor.ChooseRotor(first); rotor1 == nil {
		log.Printf("no first rotor %q\n", first)
		return nil
	}
	if rotor2 = rotor.ChooseRotor(second); rotor2 == nil {
		log.Printf("no second rotor %q\n", second)
		return nil
	}
	if rotor3 = rotor.ChooseRotor(third); rotor3 == nil {
		log.Printf("no third rotor %q\n", third)
		return nil
	}

	return &Machine{
		rotor1:    rotor1,
		rotor2:    rotor2,
		rotor3:    rotor3,
		reflector: rotor.ReflectorB,
	}
}

func (m *Machine) EncryptBuffer(text []rune) []rune {
	var output []rune
	for _, letter := range text {
		if letter < 'A' || letter > 'Z' {
			continue
		}
		output = append(output, m.EncryptLetter(letter))
	}
	return output
}

func (m *Machine) EncryptLetter(inLetter rune) rune {

	// Give the input letter to the first rotor as a contact position,
	// which is 0 for 'A', 1 for 'B', 2 for 'C', etc etc
	outPos, carry := m.rotor1.CipherFwd(int(unicode.ToUpper(inLetter)-'A'), 1, false)
	outPos, carry = m.rotor2.CipherFwd(outPos, carry, false)
	outPos, carry = m.rotor3.CipherFwd(outPos, carry, false)

	outPos = m.reflector.Reflect(outPos)

	outPos = m.rotor3.CipherBkwd(outPos, false)
	outPos = m.rotor2.CipherBkwd(outPos, false)
	outPos = m.rotor1.CipherBkwd(outPos, false)

	return rune(outPos + 'A')
}

// SetRotors metaphorically turns the target enigma.Machine's
// rotor representations to the first three letters of the settings argument.
// The settings variable should be at least 3 letters long, [A-Z].
func (m *Machine) SetRotors(settings string) {

	// Reset rotors to 0 position, just in case settings formal argument
	// has a rune that doesn't fit.
	m.rotor1.Steps = 0
	m.rotor2.Steps = 0
	m.rotor3.Steps = 0

	for i, letter := range settings {
		setting := int(unicode.ToUpper(letter))
		if setting < 'A' || setting > 'Z' {
			log.Printf("Ignoring bad setting %c\n", setting)
			continue
		}
		setting -= 'A'
		switch i {
		case 0:
			m.rotor1.Steps = setting
		case 1:
			m.rotor2.Steps = setting
		case 2:
			m.rotor3.Steps = setting
		default:
			log.Printf("unused rotor %d  setting %c\n", i+1, setting+'A')
		}
	}
}
