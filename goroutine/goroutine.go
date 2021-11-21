package main

import "fmt"

func main() {
	go func() {
		defer fmt.Println("closing")
		fmt.Println("doing somethig")
	}()
}
