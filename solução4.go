package main
import "fmt"
func main() {
	number := 0203
	fmt.Println(num)

	sumLastTwoDigits(&num)
	fmt.Print(num)
}
func sumLastTwoDigits(num *int) {
	lastDigit := *num % 10
	penultimateDigit := (*num / 10) % 10
	*num = lastDigit + penultimateDigit
}

