package models

type Pessoa struct {
	Nome  string
	Idade int
}

func NovaPessoa(nome string, idade int) *Pessoa {
	return &Pessoa{
		Nome:  nome,
		Idade: idade,
	}
}
