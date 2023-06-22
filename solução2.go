package main

import "fmt"

func main() {
	verificar := 5
	atualizarparimpar(&verificar)
	fmt.Print("O número esta dentro dos parametros: ", verificar)
}

func atualizarparimpar(verificar *int) {
	if *verificar%2 == 0 {
		*verificar = 0
	} else {
		*verificar = 1
	}
}

