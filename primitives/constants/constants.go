package main

import (
	"fmt"
	"strconv"
)

func main() {
	//Defining a constant: the naming convention for constanst is erm all caps snake case
	const VALUE int64 = 30
	fmt.Printf(strconv.FormatInt(VALUE, 2))

	//Constant declaration block
	const (
		PI             = 3.142
		RONALDO_NUMBER = 7
	)

	fmt.Println(PI, RONALDO_NUMBER)
	/*
		Naming Conventions for Golang Constants
			Name of constants must follow the same rules as variable names, which means a valid constant name must starts with a letter or underscore, followed by any number of letters, numbers or underscores.
			By convention, constant names are usually written in uppercase letters. This is for their easy identification and differentiation from variables in the source code.
	*/

}
