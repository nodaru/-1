package util

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

//Hash return a hashed string
func Hash(str string) string {
	h := sha256.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

//HashPass return a hashed pass
func HashPass(password string, salt string) string {
	return Hash(salt + Hash(strings.ToLower(password)))
}