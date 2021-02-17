package model

// User :
type User struct {
	Username  string      `json:"username" bson:"_id"`
	Details   UserDetails `json:"details" bson:"details"`
	Contact   UserContact `json:"contacts" bson:"contacts"`
	CreatedOn int64       `json:"created_on" bson:"created_on"`
	Password  string      `json:"-" bson:"password"`
}

type UserDetails struct {
	Name   string     `json:"name" bson:"name"`
	DOB    int64      `json:"dob" bson:"dob"`
	Gender GenderType `json:"gender" bson:"gender"`
}

type UserContact struct {
	EmailID string `json:"email_id" bson:"email_id"`
}

type GenderType uint8

const (
	GenderTypeMale GenderType = iota + 1
	GenderTypeFemale
)
