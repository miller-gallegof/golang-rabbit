package functions

func Add(a int, b int) int {
	return a + b
}

/*
	para hacer dinamico los parametros de una funcion se utiliza la siguiente sintaxis
*/

func AddDinamycValues(values ...int) {
	/*
		lo que hace esto es almacenar los valores en un array
	*/

	for i := range values {
		println(values[i])
	}

}
