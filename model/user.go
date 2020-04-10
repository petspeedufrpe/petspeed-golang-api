package model

// User represents the schema users on database.
type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UserPerson is the aggregate of user and person schemas.
type UserPerson struct {
	User   `json:"user"`
	Person `json:"person"`
}

// UserCostumer is the aggregate of user, person and costumer schemas.
type UserCostumer struct {
	User     `json:"user"`
	Person   `json:"person"`
	Costumer `json:"costumer"`
}
