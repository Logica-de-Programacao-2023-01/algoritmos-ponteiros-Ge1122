package main

import 	"fmt"

func main() {
	number := 2.74
	fmt.Println(number)
	media(&number)
	fmt.Print(number)
}
func media(number *float64) {
	*number = (*number + 3.14) / 2
}
