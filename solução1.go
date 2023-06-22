package main
import "fmt"

func refreshValue(n int, sum *int) {
	*sum = 0
	for i := 1; i <= n; i++ {
		*sum += i
	}
}

func main() {
	var sum int
	refreshValue(15, &sum)
	fmt.Print(sum)
}
