package model

// Person represents the schema of persons in database.
type Person struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Document string `json:"document"`
	UserID   int    `json:"userID"`
}
