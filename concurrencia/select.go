package main

import "fmt"

func enviar(par chan<- int, inpar chan<- int, salir chan<- bool) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			// Se asignan valores a los canales
			par <- i
		} else {
			// Se asignan valores a los canales
			inpar <- i
		}
	}
	close(salir)
}

func recibir(par <-chan int, inpar <-chan int, salir <-chan bool) {
	for {
		select {
		/*
			Funciona similar al switch
		*/
		//Si hay un valor en el canal par
		// se ejecuta lo que se encuentra en el case
		case v := <-par:
			fmt.Println("Desde el canal par: ", v)
			//Si hay un valor en el canal inpar
			// se ejecuta lo que se encuentra en el case
		case v := <-inpar:
			fmt.Println("Desde el canal inpar: ", v)
			//Si hay un valor en el canal salir
		case v := <-salir:
			fmt.Println("Desde el canal salir: ", v)
			return
		}
	}
}

func selectChannel() {

	/*
		Select es una estructura de control que permite ejecutar multiples canales
		cuando se tienen varios canales listos para ser extraidos
		se utiliza select para extraer esos valores
	*/

	//Creacion de canales

	par := make(chan int)
	inpar := make(chan int)
	salir := make(chan bool)

	/*
		Enviar de valores a canales en goruntines
	*/
	go enviar(par, inpar, salir)

	/*
		Recepcion de valores
	*/

	recibir(par, inpar, salir)

	fmt.Println("finalizando")
}
