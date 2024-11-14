package main

import (
	"fmt"
	"strconv"
)

func main() {
	const myConst = "hello world"
	/*
		para imprimir una constante se usa el formato fmt.Printf
		funciona muy parecido a lo que es los template string en js
		sin embargo se utiliza el %s para indicar la variable que se va a imprimir
		como segundo parametro se envian las variables que se van a imprimir
		%d -> son datos tipo entero
		%f -> son datos tipo flotante
		%s -> son datos tipo string
		%t -> son datos tipo boolean
		%v -> son datos de cualquier tipo
	*/
	fmt.Printf("%s\n", myConst)

	/*
		para concatenar y retornar valores se utiliza la funcion fmt.Sprintf
	*/

	newString := fmt.Sprintf("my nuevo %s", myConst)
	fmt.Println(newString)

	/*
		de igual forma con este metodo se puede cambiar el formato de los valores
	*/
	varFloat := 3.1416
	newString = fmt.Sprintf("my nuevo %f", varFloat)

	// de string a entero
	inVal, _ := strconv.ParseInt("1333", 0, 64)
	fmt.Printf("%d\n", inVal)

}
