package crypto4go

import (
	"bytes"
)

func ZeroPadding(text []byte, blockSize int) []byte {
	var diff = blockSize - len(text)%blockSize
	var paddingText = bytes.Repeat([]byte{0}, diff)
	return append(text, paddingText...)
}

func PKCS7Padding(text []byte, blockSize int) []byte {
	var diff = blockSize - len(text)%blockSize
	var paddingText = bytes.Repeat([]byte{byte(diff)}, diff)
	return append(text, paddingText...)
}

func PKCS7UnPadding(text []byte) []byte {
	var length = len(text)
	var unpadding = int(text[length-1])
	return text[:(length - unpadding)]
}
