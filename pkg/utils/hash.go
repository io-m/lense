package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"hash"
)

const keyString = "lenses-key-string"

// HMAC is struct that holds hash field
type HMAC struct {
	hmac hash.Hash
}

// NewHMAC is constructor function for making new instances of HMAC structure
func NewHMAC() *HMAC {
	h := hmac.New(sha256.New, []byte(keyString))
	return &HMAC{
		hmac: h,
	}
}

// HashFunc is function for generating hashed strings
// Used for rembember tokens
// It takes input string from users
// Restes it first so it can be hashed from clean start
// it write hmac encoded into provided input
// and returns summed value from the whole process
func (h *HMAC) HashFunc(inputString string) string {
	h.hmac.Reset()
	h.hmac.Write([]byte(inputString))
	bytes := h.hmac.Sum(nil)
	return base64.URLEncoding.EncodeToString(bytes)
}
