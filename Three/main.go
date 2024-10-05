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
	// below is for loop

	for x := 'a'; x <= 'z'; x++ {
		pl(x)
	}

	// for loop is also used to create while loops

	fx := 0

	for fx < 5 {
		pl(fx)
		fx++
	}
	// for true { INFINITE Loop
	// 	pl(rand.Intn(2))

	// }

	var arr1 [5]int

	arr2 := [5]int{1, 2, 3, 4, 5}

	for _, value := range arr1 {
		pl(value)
	}
	for _, value := range arr2 {
		pl(value)
	}
}
