package main

import (
	"context"
	"fmt"
	"time"
)

func simulateLongOperation(ctx context.Context) {
	for i := 0; i < 5; i++ {
		select {
		case <-time.After(1 * time.Second):
			fmt.Printf("Operation completed %d \n", i)
			/*
				La funcion done del ctx -> basicamente cuando el context se cancela
				la funcion done retorna un valor
				y en este caso se va a ejecutar el case <-ctx.Done()
			*/
		case <-ctx.Done():
			fmt.Println("Operation canceled")
			return
		}
	}
	fmt.Println("Operation completed")
}

/*
	Contextos -> son una forma de comunicación entre gorutinas
	- Se pueden cancelar
	- Se pueden usar para enviar señales
	- Se pueden usar para enviar valores
*/

func main() {

	//Ejemplo de cancelacion de goruntines
	//Creacion del contexto
	/*
		Background: es el contexto padre de todos los contextos
	*/
	ctx := context.Background()

	/*
		Se crea un contexto hijo que depende de ctx
		este nuevo contexto genera un timeout de 3 segundos en este caso
		retorna una copia del contexto padre con el timeout
		cancelFunc -> (forma de cancelar manualmente antes de el timeout) es una funcion que se ejecuta cuando el contexto hijo se cancela
	*/
	ctxHijo, cancelFunc := context.WithTimeout(ctx, 3*time.Second)

	defer cancelFunc()

	go simulateLongOperation(ctxHijo)

	time.Sleep(6 * time.Second)

	/*
			De igual forma el contexto puede almacenar valores que pueden ser utilizados por diferentes goruntine o funciones
		como primer parametro se envia el context al cual se desea agregar el valor
	*/

	ctxHijoV := context.WithValue(ctx, "env", "sandbox")

	fmt.Println(ctxHijoV.Value("env"))

}
