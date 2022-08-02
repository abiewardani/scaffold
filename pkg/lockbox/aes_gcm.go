package lockbox

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"log"

	secureRandom "gitlab.com/abiewardani/scaffold/pkg/secure_random"
)

type AesGcm interface {
	Decrypt(cipherText string, key []byte) (plainText string)
	Encrypt(plainText []byte, key []byte) (cipherText string, err error)
}

type aesGcmCtx struct {
	Key string
}

func NewAesGcm() AesGcm {
	return &aesGcmCtx{}
}

func (c *aesGcmCtx) Encrypt(message []byte, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err.Error())
		return ``, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return ``, err
	}

	nonce := make([]byte, aesgcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return ``, err
	}

	ciphertext := aesgcm.Seal(nonce, nonce, message, nil)
	dst := make([]byte, base64.StdEncoding.EncodedLen(len(ciphertext)))
	base64.StdEncoding.Encode(dst, ciphertext)

	return string(dst), nil
}

func (c *aesGcmCtx) Decrypt(encryptedString string, key []byte) (decryptedString string) {
	if encryptedString == `` {
		return ``
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Fatal(err)
	}

	sz := aesgcm.NonceSize()
	enc, err := base64.StdEncoding.DecodeString(encryptedString)
	if err != nil {
		fmt.Println(err.Error())
	}

	nonce, ciphertext := enc[:sz], enc[sz:]
	plainText, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		log.Fatal(err)
	}

	return string(plainText)
}

func GenerateNonce() ([]byte, error) {
	return secureRandom.RandomBytes(secureRandom.NonceBytes)
}
