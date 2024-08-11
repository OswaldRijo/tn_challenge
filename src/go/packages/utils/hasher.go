package utils

import (
	"encoding/base64"

	"golang.org/x/crypto/scrypt"
)

func HashString(str string, saltStr string) (string, error) {
	salt := []byte(saltStr)
	dk, err := scrypt.Key([]byte(str), salt, 16384, 8, 1, 32)

	return base64.StdEncoding.EncodeToString(dk), err
}
