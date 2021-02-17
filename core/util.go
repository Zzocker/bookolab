package core

import (
	"crypto/md5"
	"encoding/hex"
)

func hash(s string) string {
	return hex.EncodeToString(md5.New().Sum([]byte(s)))
}
