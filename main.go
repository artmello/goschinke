package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/artmello/goschinke/schinke"
)

func main() {
	start := time.Now()
	var total, ok int
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		fields := strings.Fields(s.Text())
		if len(fields) != 3 {
			continue
		}

		total++
		word, noun, verb := fields[0], fields[1], fields[2]
		actualNoun, actualVerb := schinke.Stem(word)

		if actualNoun != noun {
			fmt.Printf("[%v] ERROR [noun]: expected %v, actual %v\n", word, noun, actualNoun)
			continue
		}
		if actualVerb != verb {
			fmt.Printf("[%v] ERROR [verb]: expected %v, actual %v\n", word, verb, actualVerb)
			continue
		}

		ok++
	}
	if err := s.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "goschinke: %v\n", os.Args[0], err)
		return
	}

	fmt.Printf("Accuracy: %.2f%% [%d/%d]\n", float64(100*ok)/float64(total), ok, total)
	fmt.Printf("goschinke took %gs\n", time.Since(start).Seconds())
}
