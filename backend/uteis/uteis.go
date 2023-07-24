package uteis

import (
	"crypto/sha256"
	"encoding/hex"
)

func CreatedHash256(words string) string {
	hash := sha256.New()
	hash.Write([]byte(words))
	return hex.EncodeToString(hash.Sum(nil))
}

func CompareHash(input string, hashString string) bool {
	hash := sha256.New()
	hash.Write([]byte(input))
	wordsHash := hex.EncodeToString(hash.Sum(nil))

	return wordsHash == hashString
}
