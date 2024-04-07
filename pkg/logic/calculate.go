package logic

import (
	"crypto/sha256"
	"encoding/base64"
)

func GenerateBase64Hash(data string) string {
	hasher := sha256.New()
	hasher.Write([]byte(data))

	result := hasher.Sum(nil)

	resultBase64 := base64.StdEncoding.EncodeToString(result)
	return resultBase64
}
