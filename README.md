# Enigma-like Encryption

I recently watched a Computerphile video, [Cracking Enigma in 2021](https://www.youtube.com/watch?v=RzWB5jL5RX0).

I wanted to try to decrypt Engima-encrypted ciphertext via
the method of [CIPHERTEXT-ONLY CRYPTANALYSIS OF ENIGMA](https://web.archive.org/web/20060720040135/http://members.fortunecity.com/jpeschel/gillog1.htm)
by James J. Gillogly.
The first step is coming up with an emulated Enigma machine.
I had a lot of trouble writing an Engima emulator.
There's a lot of info floating around the web, but a lot of it is contradictory.
An example: the "ring setting". Everything you can find is super vague,
or is internally contradictory.

This is an Enigma-like encryption machine emulator.

Some of the specifics that web descriptions are weak on:

1. Advancing "the next" rotor. When does it happen? When 'A' comes around again?
After 26 keypresses? Or is it adjustable with "ring setting"?
2. What on earth does "ring setting" do?
3. How do you keep input contacts on a rotor differ from the content energized by a keypress?

I'm assuming:

1. Every time a rotor gets advanced 26 letters, the next rotor gets advanced 1 letter
2. I'm punting on understanding "ring setting".
I'm just going to assume that the ring setting is just a letter you dial a rotor to,
so it's like starting that many rotor advancing steps into an encryption.
3. I'm assuming that an 'A' keypress is a 0 input to the first, leftmost, rotor.

## How advancing a step effects a rotor
