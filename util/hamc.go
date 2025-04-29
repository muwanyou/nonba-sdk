package util

import (
	"crypto/hmac"
	"hash"
)

func GenerateHMAC(hashFunc func() hash.Hash, secretKey []byte, message []byte) []byte {
	hm := hmac.New(hashFunc, secretKey)
	hm.Write(message)
	return hm.Sum(nil)
}
