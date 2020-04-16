package model

// Costumer represents the schema costumers on database.
type Costumer struct {
	CostumerID       int `json:"costumer_id"`
	CostumerPersonID int `json:"person_id"`
}
