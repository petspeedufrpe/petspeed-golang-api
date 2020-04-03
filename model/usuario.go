package model

// Usuario representa a tabela de usuário no banco de dados.
type Usuario struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Senha string `json:"senha"`
}
