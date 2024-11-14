package main

import (
	"fmt"
	"fundamentos2/function/functions"
)

func main() {
	result := functions.Add(1, 2)
	fmt.Println(result)

	functions.AddDinamycValues(1, 2, 3, 4, 5)
}
