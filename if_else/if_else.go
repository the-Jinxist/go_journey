package main

import "fmt"

func main() {

	value := 10

	//Simple if-else flow
	if value < 1 {
		fmt.Println("Value is less than 1 mehn")
	} else {
		fmt.Println("Shet!")
	}

	//If-Else flow control structures with extra else if branch:)
	if value < 1 {
		fmt.Println("Second time is the charm")
	} else if value >= 5 && value <= 9 {
		fmt.Println("somewhere between 5 and 9")
	} else {
		fmt.Println("omo e no catch am o")
	}

	//Go, if-else is very fance, it even allows for us to initialize a variable then test if for a condition
	//observe:

	if x := 100; x < 100 {
		fmt.Println("monkey dey climb")
	} else {
		fmt.Println("send the monkey back to the village")
	}

	//Also, i tested an saw that you can do the same thing with structs, seems it only works with primitive data types. shees

}
