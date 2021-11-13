package main

import (
	"fmt"
	"strconv"
)

func main() {

	//Converting data types seems to be done mainly throught the "str conv library"

	//Converting string to integer: the result is two values, the converted value or the error occurred
	//while the conversion happens
	var valueInString string = "45"
	var sameValueInInt, err = strconv.ParseInt(valueInString, 0, 16)

	fmt.Println("Value in integer: ", sameValueInInt)
	fmt.Printf("err: %v\n", err)

	//Converting string to float: same with earlier: the error and the result
	var floatValueInString string = "4.3456"
	var sameValueInFloat, floatErr = strconv.ParseFloat(floatValueInString, 0)

	fmt.Println("Value in float: ", sameValueInFloat)
	fmt.Printf("floatErr: %v\n", floatErr)

	//Okay, seems that the Parse<Datatype>(data) functions on `strconv` converts string to the data type
	//and it seems that Format<Datatype>(data) functions on `strconv` converts the data type to the string

	var intToString string = strconv.FormatInt(34, 16)
	fmt.Println("string from int: ", intToString)

}
