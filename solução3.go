package main

import "fmt"

func inerter(x *string) {
	runes := []rune(*x)
	for i, j := 0, len(runes)-1; i < len(runes)/2; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	*x = string(runes)
}

func main() {
	str := "Cachorro caramelo"
	fmt.Println("Antes da inversão:", str)

	inerter(&str)
	fmt.Print("Após a inversão: ", str)
}
