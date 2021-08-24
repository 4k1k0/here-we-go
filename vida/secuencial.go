package main

import (
	"fmt"
	"log"
	"time"
)

func significadoVidaSecuencial() int {
	log.Println("Calculando significado de la vida")
	total := 0

	for i := 1; i <= 21; i++ {
		time.Sleep(250 * time.Millisecond)
		total = i

	}

	return total
}

func significadoUniversoSecuencial() int {
	log.Println("Calculando significado del universo")
	total := 0

	for i := 1; i <= 21; i++ {
		time.Sleep(250 * time.Millisecond)
		total = i
	}
	return total
}

func main() {
	vida := significadoVidaSecuencial()
	universo := significadoUniversoSecuencial()

	significado := vida + universo

	fmt.Printf("El significado de la vida y del universo es: %d\n", significado)
}
