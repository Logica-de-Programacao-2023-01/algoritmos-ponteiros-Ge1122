package main
import "fmt"
func main() {
	variable := 38
	pointer := &variable
	fmt.Println(variable)

	MudarValor(pointer, 21)
	fmt.Print(variable)
}
func MudarValor(pointer *int, variable int) {
	*pointer = variable
}

