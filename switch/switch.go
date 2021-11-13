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
	case 10:
		fmt.Println("It's 10 get laid")
	case 13:
		fmt.Println("It's actually today")
	}

}
