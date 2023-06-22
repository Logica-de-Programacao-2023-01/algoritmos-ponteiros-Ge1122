package main

import "fmt"

func numerosprimos(list *[]int, size int) {
	for i := 2; i < size; i++ {
		Numero_primo := true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				Numero_primo = false
				break
			}
		}
		if Numero_primo {
			*list = append(*list, i)
		}
	}
}

func main() {
	var list []int
	numerosprimos(&list, 14)
	fmt.Print(list)
}
