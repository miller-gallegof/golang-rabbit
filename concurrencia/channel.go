package main

import (
	"fmt"
	"time"
)

/*
se crea una funcion que acepta diferentes parametros
<-chan es un canal de solo lectura es decir, unicamente va a leer datos de los canales
chan<- es un canal de solo escritura es decir nunca se va a leer datos
*/

func worker(id int, jobs <-chan int, result chan<- int) {
	for j := range jobs {
		fmt.Printf("Trabajador %d empezo trabajo %d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("Trabajador %d termino trabajo %d\n", id, j)
		/*
			se asigna al canal el valor de j*2
		*/
		result <- j * 2
	}
}

func channels() {
	/*
		Es un conducto donde se puede enviar o recibir datos
		se deben tipar los tipos de datos que los canales van a recibir
	*/

	//Creación de canal

	/*
		Channel -> es un conducto donde se puede enviar o recibir datos
		se puede enviar datos entre goruntine
		son seguros
	*/

	//Creación de canal
	//se crea con make y se espeficia el tipo de dato que se va a enviar
	message := make(chan string)
	go func() {
		//envio de datos al canal
		//se asigna con <- el valor que se va a enviar
		message <- "Hello from channel"
	}()
	//asignacion de valor que trae el canal
	msg := <-message
	fmt.Println(msg)

	//la funcion main se ejecuta como una goruntine

	/*canales con bufer
		permite que el canal pueda almacenar un numero determinado de mensajes
	para asignarle un bufer en la funcion make, como segundo parametro se asigna un entero

	*/

	jobs := make(chan int, 100)
	result := make(chan int, 100)
	jobs <- 1
	jobs <- 2

	/*
		Los canales en cierto modo se comporta como publicadores
		dentro de la goruntime se comporta como consumidor
		ejecutando la goruntime cada que el valor se envie por el canal
		siempre y cuando el canal no este en estado close

	*/
	for i := 1; i <= 5; i++ {
		go worker(i, jobs, result)
	}

	//para que la goruntine deje de escuchar los diferentes valores que se envian por el canal
	//se cierra el canal

	/*
		Ya que la goruntine esta escuchando los valores que se envian por el canal
		se pueden asignar los jobs despues de ejecutar la goruntine
	*/
	jobs <- 3
	jobs <- 4
	time.Sleep(time.Second)
	jobs <- 5

	close(jobs)
	fmt.Println(<-result)

}
