package main

import "fmt"

type Conta struct {
	Saldo   float64
	Titular string
}
func main() {
	conta := Conta{
		Saldo:   4026,
		Titular: "Pericles",
	}
	fmt.Println(conta)

	MaisDinheiro(&conta)
	fmt.Print(conta)
}
func MaisDinheiro(conta *Conta) {
	conta.Saldo = conta.Saldo + 1212
}

