package main

import "fmt"

/*
 trabajar con la direccion de memoria de las variables
*/

func main() {
	var value int
	value = 100
	fmt.Println("value: ", value)

	/*
		& obtiene la direccion de memoria de la variable
	*/

	fmt.Println("address of value: ", &value)

	/*
		declarar un puntero - se declara con un *
	*/

	var valuePuntero *int
	valuePuntero = &value

	fmt.Println("valuePuntero: ", valuePuntero)
	fmt.Println("valuePuntero: ", *valuePuntero)

	/*
			var p *int -> declara un puntero
			*p -> accede al valor de la direccion de memoria
		un puntero es basicamente una direccion de memoria, donde se almacenan los valores de las variables
	*/

	i := 10

	var p *int

	p = &i

	fmt.Println(p)

	/*
		las variables tienen los siguientes datos:
		name -> nombre de la variable
		Type -> tipo de dato
		Adress -> direccion de la memoria
		Value -> el valor de la variable
	*/

	/*
		valores de myBar
		name -> myBar
		Type -> string
		Adress -> 0x001
		Value -> "hola"
	*/

	var myBar string = "hola"

	/*
			Puntero -> es una variable que almacena la direccion en memoria de un valor
			& -> obtiene la direccion de memoria de una variable
		Nota: al ser un puntero no se puede asignar un dato directamente
		Se debe declarar es un espacio en memoria
	*/

	/*
		valores de ptr
		name -> ptr
		Type -> *string
		Adress -> 0x002
		Value -> 0x001
	*/

	var ptr *string = &myBar

	fmt.Println("ptr: ", ptr)
	fmt.Println("address ptr: ", &ptr)
	fmt.Println("myBar: ", myBar)
	fmt.Println("address myBar: ", &myBar)

	/*
		puedo acceder al valor de la direccion de memoria del puntero
		se referencia al puntero y eso se hace por medio de *
	*/

	fmt.Println("value ptr: ", *ptr)

	/*
		para modificar el valor del espacio en memoria del puntero
		se debe hacer referencia al puntero y asignarle un valor
	*/

	*ptr = "mundo"

	fmt.Println("value ptr: ", *ptr)
	fmt.Println("myBar: ", myBar)

}
