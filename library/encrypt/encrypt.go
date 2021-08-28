package lib

import (
	"crypto/sha512"
	"encoding/hex"
)

// encryptSHA512 lib for encrypt data
func EncryptSHA512(data string) string {
	encData := sha512.New()
	encData.Write([]byte(data))
	strEnc := encData.Sum(nil)
	return hex.EncodeToString(strEnc)
}
