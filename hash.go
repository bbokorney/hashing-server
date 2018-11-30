package main

import (
	"crypto/sha512"
	"encoding/base64"
)

func Hash(password string) string {
	h := sha512.New()
	h.Write([]byte(password))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
