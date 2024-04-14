package tools

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"

	"io"
)

func passwordToKey(password []byte) []byte {
	hash := sha256.New()
	hash.Write([]byte(password))
	byteString := hash.Sum(nil)
	return byteString
}

func aesEncrypt(text []byte, key []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}
	encryptedText := gcm.Seal(nonce, nonce, text, nil)
	return encryptedText, nil
}

func Enrypt(plaintext []byte, password []byte) ([]byte, error) {
	encrypted, err := aesEncrypt(plaintext, passwordToKey(password))
	if err != nil {
		return nil, err
	}
	return encrypted, nil
}
