package main

import (
	"fmt"
	"sync"
)

// Estructura que representa la sala de conferencias
type ConferenceRoom struct {
	capacity int
	count    int
	mutex    sync.Mutex
	enterCh  chan struct{} // Canal para permitir la entrada
	exitCh   chan struct{} // Canal para permitir la salida
}

// Función para que una persona intente entrar a la sala de conferencias
func (cr *ConferenceRoom) enterConferenceRoom() {
	cr.enterCh <- struct{}{} // Intentar enviar un valor al canal de entrada
}

// Función para que una persona salga de la sala de conferencias
func (cr *ConferenceRoom) exitConferenceRoom() {
	cr.exitCh <- struct{}{} // Intentar enviar un valor al canal de salida
}

// Función para controlar el acceso a la sala de conferencias
func manageAccess(cr *ConferenceRoom) {
	for {
		select {
		case <-cr.enterCh:
			// Intento de entrada a la sala
			cr.mutex.Lock() // Adquirir el bloqueo del mutex
			if cr.count < cr.capacity {
				cr.count++ // Incrementar el conteo de personas en la sala
				fmt.Println("Persona ingresó a la sala")
			} else {
				fmt.Println("La sala está llena. La persona debe esperar.")
			}
			cr.mutex.Unlock() // Liberar el bloqueo del mutex

		case <-cr.exitCh:
			// Intento de salida de la sala
			cr.mutex.Lock() // Adquirir el bloqueo del mutex
			if cr.count > 0 {
				cr.count-- // Decrementar el conteo de personas en la sala
				fmt.Println("Persona salió de la sala")
			} else {
				fmt.Println("La sala está vacía. No hay personas para salir.")
			}
			cr.mutex.Unlock() // Liberar el bloqueo del mutex
		}
	}
}

func main() {
	cr := &ConferenceRoom{
		capacity: 10,
		count:    0,
		enterCh:  make(chan struct{}),
		exitCh:   make(chan struct{}),
	}

	// Iniciar goroutine para controlar el acceso a la sala de conferencias
	go manageAccess(cr)

	// Simulación de personas intentando ingresar y salir de la sala
	for i := 1; i <= 15; i++ {
		if i%2 == 0 {
			cr.enterConferenceRoom() // Intento de entrada a la sala
		} else {
			cr.exitConferenceRoom() // Intento de salida de la sala
		}
	}

	// Esperar a que todas las personas hayan terminado de ingresar y salir
	// Esto se hace para evitar que el programa finalice antes de completar todas las operaciones
	wg := sync.WaitGroup{}
	wg.Add(15)
	for i := 0; i < 15; i++ {
		go func(i int) {
			defer wg.Done()
			if i%2 == 0 {
				cr.enterConferenceRoom() // Intento de entrada a la sala
			} else {
				cr.exitConferenceRoom() // Intento de salida de la sala
			}
		}(i)
	}
	wg.Wait()

	fmt.Println("Programa finalizado")
}
