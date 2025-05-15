package main

import (
	"fmt"
	"log"
	"os"
	"unicode"
)

func main() {
	buffer, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	N := 0
	frequencies := make(map[rune]int)
	for _, r := range string(buffer) {
		if !unicode.IsLetter(r) {
			continue
		}
		frequencies[unicode.ToUpper(r)]++
		N++
	}

	sum := 0
	for _, freq := range frequencies {
		sum += freq * (freq - 1)
	}

	ic := float64(sum) / float64(N*(N-1))

	fmt.Printf("%d\t%0.5f\n", N, ic)
}
