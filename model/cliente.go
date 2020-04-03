package model

// Cliente representa a tabela cliente no banco de dados.
type Cliente struct {
	ID       int `json:"id"`
	IDPessoa int `json:"IDPessoa"`
}
