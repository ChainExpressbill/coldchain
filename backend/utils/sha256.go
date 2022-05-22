package utils

import (
	"crypto/sha256"
	"fmt"
)

func EncryptSHA256(text string) string {
	encryptText := sha256.Sum256([]byte(text))
	return fmt.Sprintf("%x", encryptText)
}
