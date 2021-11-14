package main

import "fmt"

func main() {
	panicMode(1)
	recoverPanicMode(1)
}

//Function that shows panic at work. Too much panic no good mehn
//Panic is basically like throwing exceptions
func panicMode(mode int) {
	if mode == 0 {
		fmt.Println("Everybody calm down, don't panic!")
	} else {
		panic("bROOOOOO")
	}
}

//So the next thing is like catching the execption thrown by panic, we use the `recover` keyword for this mehn
func recoverPanicMode(mode int) {
	var action int
	fmt.Println("Enter 1 for Student and 2 for Professional")
	fmt.Scanln(&action)
	/*  Use of Switch Case in Golang */
	switch action {
	case 1:
		fmt.Printf("I am a  Student")
	case 2:
		fmt.Printf("I am a  Professional")
	default:
		panic(fmt.Sprintf("I am a  %d", action))
	}

	defer func() {
		fmt.Println("ran")
		action := recover()
		fmt.Println("no panicking here: ", action)
	}()

}
