package main

import (
	"fmt"
	"log"
	"time"
)

func significadoVida(c chan int) {
	log.Println("Calculando significado de la vida")
	total := 0

	for i := 1; i <= 21; i++ {
		time.Sleep(250 * time.Millisecond)
		total = i
	}
	c <- total
	close(c)
}

func significadoUniverso(c chan int) {
	log.Println("Calculando significado del universo")
	total := 0

	for i := 1; i <= 21; i++ {
		time.Sleep(250 * time.Millisecond)
		total = i
	}
	c <- total
	close(c)
}

func main() {
	cVida := make(chan int)
	cUniverso := make(chan int)

	go significadoVida(cVida)
	go significadoUniverso(cUniverso)

	vida := <- cVida
	universo := <- cUniverso

	significado := vida + universo

	fmt.Printf("El significado de la vida y del universo es: %d\n", significado)
}
