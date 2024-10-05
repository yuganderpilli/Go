package main

import (
	"fmt"
	"unicode/utf8"
)

var pl = fmt.Println

func main() {

	rStr := "abcdefg"

	pl("Rune count:", utf8.RuneCountInString(rStr))
	for i, runeVal := range rStr {
		fmt.Printf("%d  %c \n", i, runeVal)
	}
}
