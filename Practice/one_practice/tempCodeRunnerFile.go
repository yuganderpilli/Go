package main

import (
	"bufio"
	"fmt"
	"os"
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
	reader := bufio.NewReader(os.Stdin)
	var hello;
	reader(hello)
}
