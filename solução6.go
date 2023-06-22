package main
import "fmt"
type Livro struct {
	Titulo string
	Autor  string
}
func main() {
	livro := Livro{
		Titulo: "A escalada do princípe",
		Autor:  "Anônimo",
	}
	fmt.Println(livro)
	setUnknown(&livro)
	fmt.Print(livro)
}
func setUnknown(livro *Livro) {
	if livro.Autor == "Anônimo" {
		livro.Titulo = "Desconhecido"
	}
}
