package main

import (
	"fmt"
	"strconv"
)

type Book struct {
	ID     int64  `json:"id"`
	Author string `json:"author"`
	Title  string `json:"title"`
	Year   string `json:"year"`
}

func main() {
	fmt.Print("omo!")
}

func (b Book) printBookValues() string {
	return strconv.FormatInt(b.ID, 64)
}
