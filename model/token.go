package model

// Token :
type Token struct {
	ID       string `json:"id"`
	ExpireIn int64  `json:"-"`
}
