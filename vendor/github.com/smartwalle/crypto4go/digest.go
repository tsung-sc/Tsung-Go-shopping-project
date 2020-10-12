package crypto4go

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

func MD5(value []byte) []byte {
	var m = md5.New()
	m.Write(value)
	return m.Sum(nil)
}

func MD5String(value string) string {
	return hex.EncodeToString(MD5([]byte(value)))
}

func SHA1(value []byte) []byte {
	var s = sha1.New()
	s.Write(value)
	return s.Sum(nil)
}

func SHA1String(value string) string {
	return hex.EncodeToString(SHA1([]byte(value)))
}

func SHA256(value []byte) []byte {
	var s = sha256.New()
	s.Write(value)
	return s.Sum(nil)
}

func SHA256String(value string) string {
	return hex.EncodeToString(SHA256([]byte(value)))
}

func SHA512(value []byte) []byte {
	var s = sha512.New()
	s.Write(value)
	return s.Sum(nil)
}

func SHA512String(value string) string {
	return hex.EncodeToString(SHA512([]byte(value)))
}
