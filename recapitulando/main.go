package main

import "fmt"

// structs permitem a criação de tipos customizados
type pessoa struct {
	nome  string
	idade int
}

type agenteScreto struct {
	pessoa
	licençaParaMatar bool
}

// declaração  global com zero value
var y int // zerio value de int é 0 é de um bool é um false e de um string é uma string vazia

func main() {
	// declaração do tipo de dado pessoa com os dados agrupados
	p1 := pessoa{
		"wagner",
		29,
	}

	// composição atraves de tipos de dados customizados
	agente := agenteScreto{
		pessoa{
			"James Bond",
			49,
		},
		true,
	}

	fmt.Println(agente)

	// então falar é uma função e também um metodo, e metodos podem ser encadeados tbm com uma chain
	p1.falar()
	// chamada do novo metodo
	agente.horademata()
}

// funcções além de argumentos e retorno possuem receivers que podem ser dados customizados
// todos os dados de pessoa ficam disponiveis dentro da função que tem um receiver desse tipo
func (p pessoa) falar() {
	fmt.Println(p.nome)
}

func (as agenteScreto) horademata() {
	fmt.Println(as.nome, `diz, "minha licença é true, ovo mata oce" `)
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
