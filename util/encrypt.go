package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func Encrypt(password string) string {
	h := hmac.New(sha256.New, []byte(password))
	sha := hex.EncodeToString(h.Sum(nil))
	//println(sha)
	return sha
}
