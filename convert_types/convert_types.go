package main

import (
	"fmt"
	"reflect"
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

	//Converting from int to another base int, float to another base float, int to float(and vice versa)
	//int to int64
	var normalInt int16 = 10
	var newInt64 int64 = int64(normalInt)

	fmt.Println(reflect.TypeOf(newInt64))

	var normalFloat float32 = 4.56
	var newFloat64 = float64(normalFloat)

	fmt.Println(reflect.TypeOf(newFloat64))

	//int to float, also works vice versa too bruv
	var inter = 7
	var convertedFloat = float64(inter)
	fmt.Println(reflect.TypeOf(convertedFloat))

}
