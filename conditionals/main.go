package main

import (
	"fmt"
	"os"
)

// func main() {
// 	if len(os.Args) > 1 {
// 		fmt.Println(os.Args[1])
// 	} else {
// 		fmt.Println("Hello, I am Gopher!")
// 	}
// }


// reqrite the above code with variable
func main() {
	args := os.Args
	if len(args) > 1 {
		fmt.Println(args[1])
	} else {
		fmt.Println("Hello, Gopher!")
	}
}


