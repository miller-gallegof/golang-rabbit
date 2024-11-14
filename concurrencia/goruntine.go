package main

import (
	"fmt"
	"sync"
)

func printMessage(message string, wg *sync.WaitGroup) {
	fmt.Println(message)
	defer wg.Done()
}

func goruntine() {
	/*
			Goruntine es una funcion que se ejecuta de manera concurrente
			son ligeras y se pueden ejecutar en paralelo
			es un hilo de ejecucion que se ejecuta simultaneamente
		    se pueden crear miles de goruntine y no se ve afectado el rendimiento
			a diferencia de los threads en otros lenguajes
	*/

	//Creación de goruntine
	/*
		al ejecutar una goruntine se crea un hilo de ejecucion
		que se ejecuta de manera independiente
		por lo que si no esperamos que la goruntine termine, antes de que el programe finalice no se podra
		observar el resultado de la goruntine
		en este caso nunca se ejecutara la goruntine
	*/
	//go printMessage("Hello from goruntine")
	fmt.Println("Hello from main")

	/*
		Paquete sync.WaitGroup
		Es un paquete que nos permite esperar a que todas las goruntine terminen
		para que el programa principal pueda finalizar
	*/

	var wg sync.WaitGroup
	// wg.Add -> agrega al grupo el numero de goruntine que deberian terminar para que el flujo continue
	// wg.Done -> indica que la goruntine termino
	// wg.Wait -> espera a que todas las goruntine terminen para que el flujo pueda continuar
	wg.Add(1)
	go printMessage("Hello from goruntine with wait group", &wg)
	fmt.Println("Hello from main with wait group")
	wg.Wait()

}
