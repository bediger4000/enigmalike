# Enigma-like Encryption

I recently watched a Computerphile video, [Cracking Enigma in 2021](https://www.youtube.com/watch?v=RzWB5jL5RX0).

I wanted to try to decrypt Engima-encrypted ciphertext via
the method of [CIPHERTEXT-ONLY CRYPTANALYSIS OF ENIGMA](https://web.archive.org/web/20060720040135/http://members.fortunecity.com/jpeschel/gillog1.htm)
by James J. Gillogly.
The first step is coming up with an emulated Enigma machine.
I had a lot of trouble writing an Engima emulator.
There's a lot of info floating around the web, but a lot of it is contradictory.
An example: the "ring setting". Everything you can find is super vague,
or isn't consistent from sentence to sentence.

This is an Enigma-like encryption machine emulator.

Some of the specifics that web descriptions are weak on:

1. The "fast rotor" advances once per keypress.
Advancing "medium rotor". When does it happen? When 'A' comes around again?
After 26 keypresses? Or is it adjustable with "ring setting"?
2. What on earth does "ring setting" do?
3. How do input contacts on a rotor differ from the stationary (right hand side)
contact energized by a keypress?

I'm assuming:

1. Every time a rotor gets advanced 26 letters, the next rotor gets advanced 1 letter.
That is, I'm ignoring all the online verbiage about "notches" in the rotors.
There's also a double step oddity when advancing the physical Enigma machine's middle
rotor I'm not going to bother trying to emulate.
2. I'm punting on understanding "ring setting".
I'm just going to assume that the ring setting is just a letter you dial a rotor to,
so it's like starting that many steps or key presses into an encryption.
3. I'm assuming that an 'A' keypress is a 0 input to the first, rightmost, rotor.

## How advancing a step effects my rotor emulation

I consider the neutral setting when the 'A' contact on the rightmost (fast) rotor
aligns with the 'A' input contact.
Likewise, the 'A' contact of the medium rotor aligns with the 'A' contact of the fast rotor,
and the 'A' contact of the slowest rotor aligns with 'A' on the medium rotor.
The neutral setting is the 0 (zero) position for 'A', 1 position for 'B',
2 position for 'C' and so on.
A rotor displaces angularly, the 0, 1, 2... positions stay the same in space.

I imagine this as looking at a rotor down the axis from the left side.
Using a clock analogy, 0 position is where "12" would appear on a clock face.
The position of 13 is where "6" would appear.
The 26 position is immediately counterclockwise of 0,
and 1 is immediately clockwise.
These numbered positions stay constant in space,
while the rotors change with each emulated keypress.
I have an emulated keypress stepping the fast rotor one letter further in
the alphabet.
If The 'A' contact on a rotor appears at the 0 position,
an emulated keypress moves the rotor so that its 'B' contact is at the 0 position,
and the 'A' contact is at 25 position.
The code for my emulated rotors assume position input,
an integer between 0 and 25 inclusive, and provides a positional output.

The famous rotor wiring shuffles 26 inputs to 26 outputs,
and is usually given by a letter-to-letter correspondence:

```
Entry = ABCDEFGHIJKLMNOPQRSTUVWXYZ (rotor right side)
        ||||||||||||||||||||||||||
I     = EKMFLGDQVZNTOWYHXUSPAIBRCJ
II    = AJDKSIRUXBLHWTMCQGZNPYFVOE
III   = BDFHJLCPRTXVZNYEIWGAKMUSQO
IV    = ESOVPZJAYQUIRHXLNFTGKDCMWB
V     = VZBRGITYUPSDNHLXAWMJQOFECK
```

When a rotor steps, the wiring steps with it.
Say Rotor I is in the "fast", leftmost, position.
An 'A' keypress steps Rotor I so that the 'B' (or maybe 'Z')
contact on the right hand side of the rotor gets energized.
That means the 'K' (or maybe 'J') contact on the left hand side
of the rotor carries current through to the middle rotor.

Keeping track of this is difficult to keep in your head.
My rotor emulation code does everything in terms of numerical positions.
It figures out which letter is in the input position,
which letter the emulated wiring would energize,
then determines the output position of that energized letter.
