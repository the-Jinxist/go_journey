package main

import (
	"fmt"
	"sort"
)

func main() {

	//initializing maps with empty values
	var emnployee = map[string]int{}

	//initializing maps with initial values
	var emnployeeWithSalaries map[string]string = map[string]string{"Favour": "$100,000"}
	fmt.Println(emnployee, emnployeeWithSalaries)

	//creating maps cia the `make` function
	var anotherEmployeeMap = make(map[string]int)
	//giving values to map
	anotherEmployeeMap["Favour"] = 1000
	fmt.Println(anotherEmployeeMap)

	//getting the length of a map
	fmt.Println(len(anotherEmployeeMap))

	//deleting items in a map
	delete(anotherEmployeeMap, "Favour")
	fmt.Println()

	//iterating over a map
	for key, value := range emnployeeWithSalaries {
		fmt.Println("Key: ", key, " Value: ", value)
	}

	//clearing all the items in a map
	emnployeeWithSalaries = map[string]string{}
	fmt.Println(emnployeeWithSalaries)

	//sorting the keys in a map
	unSortedMap := map[string]int{"India": 20, "Canada": 70, "Germany": 15}
	keys := []string{}

	for key := range unSortedMap {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	fmt.Println(keys)

	//sorting values in a map
	var values = make([]int, 0, len(unSortedMap))

	for _, value := range unSortedMap {
		values = append(values, value)
	}

	sort.Ints(values)
	fmt.Println(values)

}
