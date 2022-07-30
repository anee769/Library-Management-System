package common

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func generateId(title, bookMedium string, bookType BookType) string {
	h := sha256.New()
	h.Write([]byte(title + fmt.Sprintf("%v%v", bookType, bookMedium)))
	sha256_hash := hex.EncodeToString(h.Sum(nil))
	return sha256_hash
}
