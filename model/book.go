package model

// Book :
type Book struct {
	ID        string    `json:"id" bson:"_id"`
	Ownership Ownership `json:"ownership" bson:"ownership"`
}

type Ownership struct {
	Owner   string `json:"owner" bson:"owner"`
	Current string `json:"current" bson:"current"`
}
