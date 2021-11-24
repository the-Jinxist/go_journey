package main

import (
	"fmt"
	"reflect"
)

type Employee interface {
	GetTurnt()
	PrintName()
}

type Accountant struct {
	Age  int
	Name string
}

func (a Accountant) GetTurnt() {
	finalAge := a.Age
	fmt.Println(finalAge)
}

func (a Accountant) PrintName() {
	fmt.Println(a.Name)
}

func main() {

	var employee Employee = Accountant{9, "Bob"}
	employee.GetTurnt()

	FunctionTakesArgumentsOfAnyKind(employee, 4.5667, 20, "hey")

}

func FunctionTakesArgumentsOfAnyKind(values ...interface{}) {
	for arg := range values {
		fmt.Println(reflect.TypeOf(arg).Kind())
	}
}
