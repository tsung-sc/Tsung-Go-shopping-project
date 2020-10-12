package crypto4go

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

func HmacMD5(plaintext, key []byte) []byte {
	var h = hmac.New(md5.New, key)
	h.Write(plaintext)
	return h.Sum(nil)
}

func HmacMD5String(plaintext, key string) string {
	return hex.EncodeToString(HmacMD5([]byte(plaintext), []byte(key)))
}

func HmacSHA1(plaintext, key []byte) []byte {
	var h = hmac.New(sha1.New, key)
	h.Write(plaintext)
	return h.Sum(nil)
}

func HmacSHA1String(plaintext, key string) string {
	return hex.EncodeToString(HmacSHA1([]byte(plaintext), []byte(key)))
}

func HmacSHA256(plaintext, key []byte) []byte {
	var h = hmac.New(sha256.New, key)
	h.Write(plaintext)
	return h.Sum(nil)
}

func HmacSHA256String(plaintext, key string) string {
	return hex.EncodeToString(HmacSHA256([]byte(plaintext), []byte(key)))
}

func HmacSHA512(plaintext, key []byte) []byte {
	var h = hmac.New(sha512.New, key)
	h.Write(plaintext)
	return h.Sum(nil)
}

func HmacSHA512String(plaintext, key string) string {
	return hex.EncodeToString(HmacSHA512([]byte(plaintext), []byte(key)))
}
