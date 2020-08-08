package main

import "fmt"

// structs permitem a criação de tipos customizados
type pessoa struct {
	nome  string
	idade int
}

// declaração  global com zero value
var y int // zerio value de int é 0 é de um bool é um false e de um string é uma string vazia

func main() {
	// declaração do tipo de dado pessoa com os dados agrupados
	p1 := pessoa{
		"wagner",
		29,
	}

	fmt.Println(p1)
}

func variables() {
	// declaração curta de variável
	// escopo de x é a variables func
	x := 7
	fmt.Printf("%T", x)
}

func slices() {
	// sllices é uma estrutura de dados para agrupar todo tipo de coisa, inteiros, strings, strucs...
	xi := []int{2, 3, 4, 5, 65}
	fmt.Println(xi)
}

func maps() {
	// maps são outro tipo de estrutura da dados baseado em chave e valor, como hashes de ruby
	m := map[string]int{"wagner": 29}
	fmt.Println(m)
}
