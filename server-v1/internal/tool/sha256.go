package tool

import (
	"crypto/sha256"
	"encoding/hex"
)

func GetSha256String(str string) string {
	var p = sha256.Sum256([]byte(str))
	return hex.EncodeToString(p[:])
}
