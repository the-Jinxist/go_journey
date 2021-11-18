package main

import (
	"encoding/json"
	"fmt"
)

/*
A struct (short for "structure") is a collection of data fields with declared data types. Golang has the ability to declare and create own data types by combining one or more types, including both built-in and user-defined types.
Each data field in a struct is declared with a known type,
which could be a built-in type or another user-defined type.
Structs are the only way to create concrete user-defined types in Golang.
Struct types are declared by composing a fixed set of unique fields. Structs can improve modularity and allow to create and pass complex data structures around the system.
You can also consider Structs as a template for creating a data record, like an employee record or an e-commerce product.
The declaration starts with the keyword type, then a name for the new struct, and finally the keyword struct.
Within the curly brackets, a series of data fields are specified with a name and a type.
*/

type rectangle struct {
	length  float64
	breadth float64
	color   string
}

type shape struct {
	length  int
	breadth int
	color   string

	geometry struct {
		area      int
		perimeter int
	}
}

//nested structs
type Salary struct {
	Basic, HRA, TA float32
}

type Employee struct {
	FirstName, LastName, Email string
	Age                        int
	MonthlySalary              []Salary
}

//while working with json, you can assign optional strings to your struct to match field to json fields
type EmployeeJson struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	City      string `json:"city"`
}

//adding methods to structs
func (e Employee) EmployeeInfo(values string) string {
	fmt.Println(e.FirstName, e.LastName)
	fmt.Println(e.Age)
	fmt.Println(e.Email)
	for _, info := range e.MonthlySalary {
		fmt.Println("===================")
		fmt.Println(info.Basic)
		fmt.Println(info.HRA)
		fmt.Println(info.TA)
	}
	return "----------------------"
}

func main() {

	var bigRectangle = rectangle{10.5, 25.10, "red"}
	fmt.Println(bigRectangle)

	var myshape shape
	fmt.Println(myshape)

	myshape.breadth = 20
	myshape.length = 10
	myshape.color = "Green"

	myshape.geometry.area = myshape.breadth * myshape.length
	myshape.geometry.perimeter = 2 * (myshape.breadth * myshape.length)

	fmt.Println(myshape)

	//naming the parameter while instantiating
	var breadthLessRect = rectangle{length: 10, color: "color"}
	fmt.Println(breadthLessRect)

	//instatiating a struct using the `new` keyword
	rect1 := new(rectangle)
	fmt.Println(rect1)

	//accessing values in the nested struct
	e := Employee{
		FirstName: "Mark",
		LastName:  "Jones",
		Email:     "mark@gmail.com",
		Age:       25,
		MonthlySalary: []Salary{
			{
				Basic: 15000.00,
				HRA:   5000.00,
				TA:    2000.00,
			},
			{
				Basic: 16000.00,
				HRA:   5000.00,
				TA:    2100.00,
			},
			{
				Basic: 17000.00,
				HRA:   5000.00,
				TA:    2200.00,
			},
		},
	}
	fmt.Println(e.FirstName, e.LastName)
	fmt.Println(e.Age)
	fmt.Println(e.Email)
	fmt.Println(e.MonthlySalary[0])
	fmt.Println(e.MonthlySalary[1])
	fmt.Println(e.MonthlySalary[2])

	fmt.Println(e.EmployeeInfo(""))

	json_string := `
	{
		"firstname": "Rocky",
		"lastname": "Balboa",
		"city": "London"
	}`

	//taking the values in the json and putting them in struct form
	freshEmployee := new(EmployeeJson)
	json.Unmarshal([]byte(json_string), freshEmployee)
	fmt.Println(freshEmployee)

	//taking the values from the struct and putting it back in json form
	fresherEmployee := new(EmployeeJson)
	fresherEmployee.FirstName = "Brah"
	fresherEmployee.LastName = "Bruh"
	fresherEmployee.City = "City"
	jsonString, _ := json.Marshal(fresherEmployee)

	fmt.Println(jsonString)

}
