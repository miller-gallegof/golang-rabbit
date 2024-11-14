package main

import (
	"fmt"
	_struct "fundamentos2/structure/struct"
)

func main() {
	p1 := _struct.Product{
		ID:       "1",
		Name:     "Laptop",
		Price:    5000.00,
		Quantity: 10,
		Tag: _struct.Type{
			ID:          "1",
			Code:        "001",
			Description: "Electronics",
		},
	}

	fmt.Println(p1)
}
