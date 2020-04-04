package model

// Costumer represents the schema costumers on database.
type Costumer struct {
	ID       int `json:"id"`
	PersonID int `json:"personID"`
}
