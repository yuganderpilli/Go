package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var rating string

	fmt.Println("Please enter your rating")
	fmt.Scanf("%s", &rating)
	numRation, err := strconv.ParseFloat(strings.TrimSpace(rating), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your rating", numRation+1)
	}
}
