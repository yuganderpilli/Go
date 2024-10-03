package main

import (
	"fmt"
	"reflect"
)

var sample = "Hello world"
var pl = fmt.Println // ALiAS
// Above is a comment.

/* Multi
line

comment  */
// function execution strats in main function.
//

func main() {
	pl(sample)
	// reader := bufio.NewReader(os.Stdin)
	// name, _ := reader.ReadString('\n') // _ is called the blank identifier.
	// pl(name)
	// age, err := reader.ReadString('\n')
	// if err != nil {
	// 	pl(age)
	// } else {
	// 	log.Fatal(err)
	// }

	// Declaring Variables:

	var a = "hello world"
	pl(a)

	// Data types:
	pl(reflect.TypeOf(a))
	var aa = 23
	c := int(aa)
	pl(c)
}
