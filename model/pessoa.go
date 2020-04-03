package model

// Pessoa representa a tabela pessoa do banco de dados.
type Pessoa struct {
	ID        int    `json:"id"`
	Nome      string `json:"nome"`
	Documento string `json:"documento"`
	IDUsuario int    `json:"idUsuario"`
}
