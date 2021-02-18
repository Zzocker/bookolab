package model

import "time"

// Token :
type Token struct {
	ID        string    `json:"id"`
	ExpireIn  int64     `json:"expire_in"`
	CreatedAt int64     `json:"created_at"`
	Type      tokenType `json:"tokenType"`
	Username  string    `json:"username"`
}

type tokenType uint8

const (
	tokenTypeRefresh tokenType = iota + 1
	tokenTypeAccess
)

const (
	day                int64 = 86400 // seconds
	tokenRefreshExpiry int64 = (31 * day)
	tokenAccessToken   int64 = 7 * day
)

func NewRefreshToken(id string, username string) Token {
	return Token{
		ID:        id,
		ExpireIn:  tokenRefreshExpiry,
		CreatedAt: time.Now().Unix(),
		Type:      tokenTypeRefresh,
		Username:  username,
	}
}

func NewAccessToken(id string, username string) Token {
	return Token{
		ID:        id,
		ExpireIn:  tokenAccessToken,
		CreatedAt: time.Now().Unix(),
		Type:      tokenTypeAccess,
		Username:  username,
	}
}
