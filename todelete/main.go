package main

import (
	"fmt"
	// "reflect"
)

func main() {

	fmt.Println([]byte("a sentence"))
	fmt.Println([]byte("another etxt"))

	data1 := []byte{97, 32, 115, 101, 110, 116, 101, 110, 99, 101}

	fmt.Println(data1)
	fmt.Println(string(data1))
}
