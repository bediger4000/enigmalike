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

func NewMachine(first, second, third string, settings string) *Machine {
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

	for i, letter := range settings {
		setting := int(unicode.ToUpper(letter))
		if setting < 'A' || setting > 'Z' {
			log.Printf("Ignoring bad setting %c\n", setting)
			continue
		}
		setting -= 'A'
		switch i {
		case 0:
			rotor1.Steps = setting
		case 1:
			rotor2.Steps = setting
		case 2:
			rotor3.Steps = setting
		default:
			log.Printf("unused rotor %d  setting %c\n", i+1, setting+'A')
		}
	}

	return &Machine{
		rotor1:    rotor1,
		rotor2:    rotor2,
		rotor3:    rotor3,
		reflector: rotor.ReflectorB,
	}
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
