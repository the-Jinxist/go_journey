package main

import "fmt"

func main() {

	//Traditional for loop
	//let's see how many statments we can leave out of the for loop initialization and the variations
	k := 1
	for ; k <= 10; k++ {
		fmt.Println("k: ", k)
	}

	//Okay, we can't just initialize and leave it like that lol
	// for x := 1 {
	// 	fmt.Println(x)
	//}

	//Infinite loop, seems we can leave everything out. that's white there's no need for a while loop in go
	x := 0
	for {
		if x == 10 {
			break
		}
		x++
	}

	//We can remove all but the incremental count
	z := 1
	for ; ; z++ {
		if z == 10 {
			fmt.Println("z bro!")
			break
		}
	}

	//For with special `range` keyword
	//range allows us to cycle through an iterable
	for x := range "hello" {
		fmt.Println(x)
	}

}
