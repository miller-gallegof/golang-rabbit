package main

import "fmt"

func main() {
	/*
		para que sea considerado un vector se debe tener un tamaño determinado al array
	*/

	var array [5]int
	fmt.Println(array)

	/*
		los slice se comportan como arreglos dinamicos
		para crear un slice unicamente no se asigna tamaño al array
	*/
	var slice []int
	fmt.Println(slice)
}
