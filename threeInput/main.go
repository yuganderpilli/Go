package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Enter your city name")
	reader := bufio.NewReader(os.Stdin)
	city, _ := reader.ReadString('\n')
	var state string
	fmt.Printf("Enter your state name")
	fmt.Scanf("%s", &state)

	fmt.Println(city, state)
}
