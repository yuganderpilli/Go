package main

import (
	"fmt"
	"reflect"
)

func main() {
	pl := fmt.Println
	var a string = " I am a sample variable"
	fmt.Println(a)
	var b = "sample"
	fmt.Println(b)
	yugander := "Hello world, I am yugander"
	fmt.Println(yugander)
	Cv1 := true
	Cv2 := bool(Cv1)
	fmt.Println(Cv2)
	Cv3 := "hello world"
	pl(reflect.TypeOf(Cv2))
	pl(reflect.TypeOf(Cv3))
	if reflect.TypeOf(Cv2).Kind() == reflect.Int {
		fmt.Println("I am  a int")
	} else if reflect.TypeOf(Cv2).Kind() == reflect.String {
		fmt.Println("I am a string")
	} else {
		fmt.Println(" I am somthing else")
	}
}
