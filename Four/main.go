package main

import "fmt"

func sums(x int, y int) int {

	return x + y
}

func getQuotient(x float64, y float64) (ans float64, err error) {
	if y == 0 {
		return 0, fmt.Errorf("You can difived by zero")
	} else {
		return x / 2, nil
	}

}

func main() {
	pl := fmt.Println
	pl("hello world")
	// array declaration;
	var arr2 [5]int // empty array
	pl(arr2)
	arr := [5]int{1, 2, 3, 4, 5} // array with values
	pl(arr)

	//slice declaration : these are dynamic arrays

	s := make([]int, 5) // empty slice
	pl(s)
	s1 := []int{1, 2, 3} // slice with values
	pl(s1)
	s2 := []string{"hello", "world"}

	pl(s2)

	//  functions

	pl(sums(1, 2)) //
	sample, _ := getQuotient(1, 2)
	pl(sample)
}
