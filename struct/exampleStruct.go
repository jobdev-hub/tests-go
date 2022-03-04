package main

import "fmt"

type Cliente struct {
	Id       int
	Name     string
	Age      int
	Endereco []Endereco
}

func (c Cliente) toString() string {

	cidades := ""
	qtd := len(c.Endereco)

	for _, cidade := range c.Endereco {
		if qtd > 1 {
			cidades += cidade.Cidade + ", "
		} else {
			cidades += cidade.Cidade
		}
		qtd--
	}
	return fmt.Sprintf("Id: %d; Name: %s; Age: %d\nCidades: %s", c.Id, c.Name, c.Age, cidades)
}

type Endereco struct {
	Id          string
	Rua         string
	Numero      int
	Complemento string
	Cidade      string
	Estado      string
}

func (e Endereco) toString() string {
	return fmt.Sprintf("Id: %s, Rua: %s, Numero: %d, Complemento: %s, Cidade: %s, Estado: %s", e.Id, e.Rua, e.Numero, e.Complemento, e.Cidade, e.Estado)
}

func main() {

	endereco1 := Endereco{
		"1",
		"Rua Teste",
		1,
		"Complemento Teste",
		"Cidade Teste",
		"UF Teste",
	}

	endereco2 := Endereco{}
	endereco2.Id = "2"
	endereco2.Rua = "Rua Teste 2"
	endereco2.Numero = 2
	endereco2.Complemento = "Complemento Teste 2"
	endereco2.Cidade = "Cidade Teste 2"
	endereco2.Estado = "UF Teste 2"

	enderecos := []Endereco{endereco1, endereco2}
	cliente := Cliente{Id: 1, Name: "Cliente Teste", Age: 20, Endereco: enderecos}
	fmt.Println("\n" + cliente.Name)
	fmt.Println(cliente.toString())
}
