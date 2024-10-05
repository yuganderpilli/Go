package main

import (
	"fmt"
	"math/rand"
	"time"
	"unicode/utf8"
)

var pl = fmt.Println

func main() {

	rStr := "abcdefg"

	pl("Rune count:", utf8.RuneCountInString(rStr))
	for i, runeVal := range rStr {
		fmt.Printf("%d:%#U:%c \n", i, runeVal, runeVal)
	}
	now := time.Now()
	pl(now)
	pl(now.Year())
	pl(now.Month())
	pl(now.Day())
	seedSecs := now.Unix() // gives second until 1 jan 1970
	pl(seedSecs)
	// generating random number using the seed seconds value as the seed
	randNum := rand.Intn(51)
	pl(randNum)

}
