package main

import (
	"fmt"
	"reflect"
)

func main() {

	//Basic declararion of a slice
	var intSlice []int
	var strSlice []string

	fmt.Println(reflect.ValueOf(intSlice).Kind())
	fmt.Println(reflect.ValueOf(strSlice).Kind())

	//Declaring a slice with make, specifying just the length, which will in turn become the value
	//of it's capacity too
	var intSlice1 = make([]int, 10)
	//Declaring a slice with make, adding it's capacity and length
	var strSlice1 = make([]string, 20, 30)

	fmt.Println(reflect.ValueOf(intSlice1).Kind())
	fmt.Println(reflect.ValueOf(strSlice1).Kind())

	//giving a slice values using the basic declaration
	var intSlice2 = []int{10, 20, 30, 40}
	fmt.Println(len(intSlice2), cap(intSlice2))

	//declaring slice with the new keyword
	var newSlice = new([5]int)[0:5]
	fmt.Println(newSlice)

	//adding items to the slice
	newSlice[0] = 10
	newSlice[1] = 20

	newSlice = append(newSlice, 30, 30, 30, 30, 50, 90, 70, 90, 20, 30, 40)
	fmt.Println(newSlice)

	//removing item from slice
	newSlice = removeIndex(newSlice, 0)
	fmt.Println(newSlice)

	//copying a clice
	a := []int{5, 6, 7}
	b := make([]int, 5, 10)

	copy(b, a)
	fmt.Println(b)

	//Filtering a slice is basicaaly the same of filtering through an array.

	//Looping through a slice

	//Getting index and element for each loop
	for index, element := range newSlice {
		fmt.Println(index, element)
	}

	//choosing to use only the value in each loop
	for _, value := range newSlice {
		fmt.Println(value)
	}

	//simple looping through the slice, use an outside variable to get individual elements
	//in the slice
	for range newSlice {

	}

}

func removeIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}
