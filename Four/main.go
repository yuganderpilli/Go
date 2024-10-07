package main

import "fmt"

func main() {
	pl := fmt.Println
	pl("hello world")
	// array declaration;
	var arr2 [5]int
	pl(arr2)
	arr := [5]int{1, 2, 3, 4, 5}
	pl(arr)

	//slice declaration : these are dynamic arrays

	s := make([]int, 5)
	pl(s)
	s1 := []int{1, 2, 3}
	pl(s1)
	s2 := []string{"hello", "world"}

	pl(s2)
}
