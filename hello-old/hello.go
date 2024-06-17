// //write hello world program in go
// package main
// import "fmt"

// import "rsc.io/quote"
// func main() {
// 	fmt.Println("Hello, World!")
// 	fmt.Println(quote.Hello())
// }
// 	//run the program using go run filename.go

package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
