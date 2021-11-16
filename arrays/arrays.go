package main

import (
	"fmt"
	"reflect"
)

func main() {

	//declaring an array
	var intArray [5]int
	var stringArray [5]string

	fmt.Println(reflect.TypeOf(intArray))
	fmt.Println(reflect.TypeOf(stringArray))

	//assigning values to the array
	stringArray[0] = "Favour"
	stringArray[2] = "Olukayode"

	fmt.Println(stringArray)

	//initializing an array with an array literal
	x := [5]int{}
	var y = [5]int{10, 20, 30}
	fmt.Println(x)
	fmt.Println(y)

	//initializing an array with elipsis.
	//the compiler will infer the length of the array for the amount of values put in the array on init
	arrayElipsis := [...]int{5, 7, 8}
	fmt.Println(arrayElipsis)
	fmt.Println(len(arrayElipsis))

	//intitializing an array with sepecific values in specific positions
	specificInit := [5]int{0: 20, 3: 30}
	fmt.Println(specificInit)

	//Looping through an array
	for x := 0; x < len(specificInit); x++ {
		fmt.Println(x)
	}

	//copying an array
	array1 := [3]int{4, 5, 6}

	//copying the array by it's original value
	array2 := array1

	//copying by referencing. any change to the referenced value will be copied
	array3 := &array1

	array1[0] = 10

	fmt.Println(array1)
	fmt.Println(array2)
	fmt.Println(array3)

	//checking if an item exist
	arrayChecking := [3]string{"hey", "are", "you free tonight?"}
	fmt.Println(itemExists(arrayChecking, "hey"))

	//filtering arrays. in [n:m], the filtered array will start from n ganganm, but it will end in index m-1
	//if n or m is missing, it'll either start from the beginning or stop at the end respectively
	countries := [...]string{"India", "Canada", "Japan", "Germany", "Italy"}

	fmt.Printf("Countries: %v\n", countries)
	fmt.Printf(":2 %v\n", countries[:2])
	fmt.Printf("1:3 %v\n", countries[1:3])
	fmt.Printf("2: %v\n", countries[2:])
	fmt.Printf("2:5 %v\n", countries[2:5])

	fmt.Printf("0:3 %v\n", countries[0:3])

	fmt.Printf("Last element: %v\n", countries[len(countries)-1])

	fmt.Printf("All elements: %v\n", countries[0:len(countries)])
	fmt.Println(countries[:])
	fmt.Println(countries[0:])
	fmt.Println(countries[0:len(countries)])

	fmt.Printf("Last two elements: %v\n", countries[len(countries)-2:len(countries)])

}

//a function to check if an item in items exists
func itemExists(arrayType interface{}, item interface{}) bool {
	arr := reflect.ValueOf(arrayType)

	if arr.Kind() != reflect.Array {
		panic("Invalid data-type")
	}

	for i := 0; i < arr.Len(); i++ {
		if arr.Index(i).Interface() == item {
			return true
		}
	}

	return false
}
