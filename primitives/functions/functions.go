package main

import "fmt"

func main() {
	SimpleFunction()
	SimpleFunctionWithParameters(9, "nine")

	area(6, 7)
	newFuncton := higherOrderFunction(innerfunc)
	newFuncton(20)

	variadicFunctions("Favour", "Ayomide")
	variadicFunctionsWithDiffTypes(2, 2, 3, 4, 22, 3, 6.8990, "stuff", true)

	writingToFile()
}

var innerfunc = func(age int) int {
	return age * 30
}

//Simplest function in every go program every created
func SimpleFunction() {
	fmt.Println("Simplest function ever!")
}

//A function with parameters
func SimpleFunctionWithParameters(x int, y string) {
	fmt.Println("These are your parameters: x - ", x, "y - ", y)
}

//A function with a return value
func SimpleFunctionWithReturnValue(age int) int {
	ageIn30 := age + 30
	return ageIn30
}

//A function with a named return value. Yes you can name return values in go
///Here we named the return value @ageIn60: ageIn60
func SimpleFunctionWithNameReturnValue(age int) (ageIn60 int) {
	age += 60
	ageIn60 = age
	return
}

//A function with multiple values that can be returned, yeah another weird stuff
func SimpleFunctionWithMultipleReturnValues(money int) (isMillionaire bool, isPoor bool) {
	if money < 1000000 {
		isMillionaire = false
		isPoor = true
	} else {
		isMillionaire = true
		isPoor = false
	}

	return
}

//Anonymous function in go
var area = func(length int, breadth int) {
	area := length * breadth
	fmt.Println("The area is ", area)
}

//Higher-order functions:
//Higher-order functions are functions that take other functions as a parameter or returns other functions

//@higherOrderFunction takes in a function and spits out another one
func higherOrderFunction(innerfunc func(age int) int) func(age int) int {
	answer := innerfunc(60)
	fmt.Println(answer)
	return func(age int) int {
		return 0
	}
}

//Variadic functions, they can take multiple values of the same data type
func variadicFunctions(names ...string) {
	for name := range names {
		fmt.Println(name)
	}
}

//It also works with values of different types
func variadicFunctionsWithDiffTypes(values ...interface{}) {
	for value := range values {
		fmt.Println(value)
	}
}

//Deferred function calls
//Go has a special statement called defer that schedules a function call to be run after the function completes.
func writingToFile() {
	fmt.Println("writing....")
	defer fmt.Println("closing file..")

	fmt.Println("small calculations in the file and such")
}
