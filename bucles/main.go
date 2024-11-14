package main

import "fmt"

func main() {

	array := []int{1, 2, 3, 45, 6}

	for i := range array {
		fmt.Println(array[i])
	}

	// iterando un map
	mapa := map[int]string{1: "a", 2: "b", 3: "c"}

	for key, value := range mapa {
		fmt.Println(key, value)
	}

	arrayEjercicio := [5]int{4, 2, 5, 6, 7}

	for item := range arrayEjercicio {
		arrayEjercicio[item] = arrayEjercicio[item] + 20
	}

	fmt.Println(arrayEjercicio)

	var value int
	newArray := []int{}
	for {
		fmt.Println("Ingrese un número")
		fmt.Scan(&value)
		newArray = append(newArray, value)
		if value == 0 {
			break
		}
	}

	fmt.Println(newArray)

}
