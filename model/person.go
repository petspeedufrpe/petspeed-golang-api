package model

// Person represents the schema of persons in database.
type Person struct {
	PersonID int    `json:"person_id"`
	Name     string `json:"name"`
	Document string `json:"document"`
	UserID   int    `json:"user_id"`
}
