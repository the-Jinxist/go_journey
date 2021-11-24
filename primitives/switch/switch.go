package main

import (
	"fmt"
	"time"
)

func main() {

	today := time.Now()

	//Switch statement, with a little twist, as per na Go version.
	switch today.Day() {
	case 5:
		fmt.Println("Seems it's five mehn")
	case 10, 11, 12: //Here, we see that switch can have multiple values in one case branch.l
		fmt.Println("It's 10 get laid")
	case 14:
		fmt.Println("It's actually today")
	}

	//There's also a keywork known as `fallthrough` that moves the user flow to the next
	//successive case branch
	switch today.Day() {
	case 13:
		fmt.Println("13?")
	case 14:
		fmt.Println("It's today")
		fallthrough
	default:
		fmt.Println("fell through here")
	}

	//Go also allows case branches to have conditional statements, as well as, allowing the very first branch to initialize the switched value
	//but nobody asked for these features

	//Switch case statements with conditional statements
	/*
		switch {
		case today.Day() < 5:
			fmt.Println("Clean your house.")
		case today.Day() <= 10:
			fmt.Println("Buy some wine.")
		case today.Day() > 15:
			fmt.Println("Visit a doctor.")
		case today.Day() == 25:
			fmt.Println("Buy some food.")
		default:
			fmt.Println("No information available for that day.")
		}
	*/

	//Switch case statement with an initialize
	/*
		switch today := time.Now(); {
		case today.Day() < 5:
			fmt.Println("Clean your house.")
		case today.Day() <= 10:
			fmt.Println("Buy some wine.")
		case today.Day() > 15:
			fmt.Println("Visit a doctor.")
		case today.Day() == 25:
			fmt.Println("Buy some food.")
		default:
			fmt.Println("No information available for that day.")
		}
	*/

}
