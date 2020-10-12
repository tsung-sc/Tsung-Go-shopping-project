package crypto4go

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"golang.org/x/crypto/pbkdf2"
	"hash"
)

const (
	K_PKCS5_SALT_LEN      = 8
	K_PKCS5_DEFAULT_ITER  = 2048
	K_PKCS5_DEFAULT_MAGIC = "Salted__"
	K_EVP_MAX_IV_LENGTH   = 16
)

func randBytes(length int) (data []byte, err error) {
	data = make([]byte, length)
	_, err = rand.Read(data)
	return data, err
}

func AESCBCEncryptWithSalt(plaintext, key []byte, iterCount int, magic string, h func() hash.Hash) ([]byte, error) {
	return AESEncryptWithSalt(plaintext, key, iterCount, magic, h, AESCBCEncrypt)
}

func AESCBCDecryptWithSalt(ciphertext, key []byte, iterCount int, magic string, h func() hash.Hash) ([]byte, error) {
	return AESDecryptWithSalt(ciphertext, key, iterCount, magic, h, AESCBCDecrypt)
}

func AESEncryptWithSalt(plaintext, key []byte, iterCount int, magic string, h func() hash.Hash, f func(plaintext, key, iv []byte) (dst []byte, err error)) (dst []byte, err error) {
	if iterCount <= 0 {
		iterCount = K_PKCS5_DEFAULT_ITER
	}

	if h == nil {
		h = md5.New
	}

	var salt, _ = randBytes(K_PKCS5_SALT_LEN)
	var sKey = pbkdf2.Key(key, salt, iterCount, len(key), h)
	var sIV = pbkdf2.Key(sKey, salt, iterCount, K_EVP_MAX_IV_LENGTH, h)

	dst, err = f(plaintext, sKey, sIV)

	dst = append(salt, dst...)
	dst = append([]byte(magic), dst...)

	return dst, err
}

func AESDecryptWithSalt(ciphertext, key []byte, iterCount int, magic string, h func() hash.Hash, f func(ciphertext, key, iv []byte) ([]byte, error)) (dst []byte, err error) {
	if iterCount <= 0 {
		iterCount = K_PKCS5_DEFAULT_ITER
	}

	if h == nil {
		h = md5.New
	}

	//if len(ciphertext) <= len(magic) + K_PKCS5_SALT_LEN {
	//	return nil, errors.New("Error")
	//}

	var salt = ciphertext[len(magic) : len(magic)+K_PKCS5_SALT_LEN]
	var sKey = pbkdf2.Key(key, salt, iterCount, len(key), h)
	var sIV = pbkdf2.Key(sKey, salt, iterCount, K_EVP_MAX_IV_LENGTH, h)

	dst, err = f(ciphertext[len(magic)+K_PKCS5_SALT_LEN:], sKey, sIV)

	return dst, err
}

// AESCBCEncrypt 由key的长度决定是128, 192 还是 256
func AESCBCEncrypt(plaintext, key, iv []byte) ([]byte, error) {
	var block, err = aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	var blockSize = block.BlockSize()
	iv = iv[:blockSize]

	var src = PKCS7Padding(plaintext, blockSize)
	var dst = make([]byte, len(src))

	var mode = cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(dst, src)
	return dst, nil
}

func AESCBCDecrypt(ciphertext, key, iv []byte) ([]byte, error) {
	var block, err = aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	var blockSize = block.BlockSize()
	iv = iv[:blockSize]

	var dst = make([]byte, len(ciphertext))

	var mode = cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(dst, ciphertext)
	dst = PKCS7UnPadding(dst)
	return dst, nil
}

func AESCFBEncrypt(plaintext, key, iv []byte) ([]byte, error) {
	var block, err = aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	var blockSize = block.BlockSize()
	iv = iv[:blockSize]

	var dst = make([]byte, len(plaintext))

	var mode = cipher.NewCFBEncrypter(block, iv)
	mode.XORKeyStream(dst, plaintext)
	return dst, nil
}

func AESCFBDecrypt(ciphertext, key, iv []byte) ([]byte, error) {
	var block, err = aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	var blockSize = block.BlockSize()
	iv = iv[:blockSize]

	var dst = make([]byte, len(ciphertext))

	var mode = cipher.NewCFBDecrypter(block, iv)
	mode.XORKeyStream(dst, ciphertext)
	return dst, nil
}
