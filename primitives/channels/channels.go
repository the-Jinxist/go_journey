package main

import "fmt"

/* Go provides a mechanism called a channel that is used to share data between goroutines.
When you execute a concurrent activity as a goroutine a resource or data needs to be shared between goroutines,
channels act as a conduit(pipe) between the goroutines
and provide a mechanism that guarantees a synchronous exchange.
*/

func main() {

	//Code to create a channel
	unBuffered := make(chan string) // Unbuffered channel of integer type
	buffered := make(chan int, 10)  // Buffered channel of integer type

	//Syntax to send a value into the channel
	buffered <- 500
	unBuffered <- "String"

	//Code to recieve value from the channel
	value := unBuffered
	fmt.Println(value)

}
