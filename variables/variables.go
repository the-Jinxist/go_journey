package main

import (
	"fmt"
	"reflect"
)

func main() {

	//Simple variable creation with explicit typing
	var explicitTypedName string = "Favour"
	fmt.Print(explicitTypedName)

	//Simple variable creation with type inference
	var name = "Favour"
	fmt.Print(name)

	//In go, it's possible to declare variables without initializing them but you'll have to do them with
	//their data type
	//e,g monkey will be given a defult value of an empty string
	var monkey string
	fmt.Print(monkey)

	//Declaration of multiple varibales. Multiple variables with the same data type, and multiple values for each of them
	//lmao seems giving variables default values doesn't extend to multiple variable declaration
	var brother, sister, aunt = "brother", "sister", "aunt"
	fmt.Print(brother, sister, aunt)

	//Short variable declaration. //Side note: the "reflect" can be used to check the data type of objects
	hey := "wassup"
	fmt.Print(reflect.TypeOf(hey))

	//Variable declaration block
	var (
		stockName                 = "SNY"
		stockPrice                = "$1000"
		marketCap                 = "$10B"
		lastMarketActivity        = "down"
		lastMarketActivityPercent = "2.5%"
	)

	fmt.Println(stockName, stockPrice, marketCap, lastMarketActivity, lastMarketActivity, lastMarketActivityPercent)

	/*
		Naming Conventions for Golang Variables
			These are the following rules for naming a Golang variable:

			A name must begin with a letter, and can have any number of additional letters and numbers.
			A variable name cannot start with a number.
			A variable name cannot contain spaces.
			If the name of a variable begins with a lower-case letter, it can only be accessed within the current package this is considered as unexported variables.
			If the name of a variable begins with a capital letter, it can be accessed from packages outside the current package one this is considered as exported variables.
			If a name consists of multiple words, each word after the first should be capitalized like this: empName, EmpAddress, etc.
			Variable names are case-sensitive (car, Car and CAR are three different variables).

	*/

}
