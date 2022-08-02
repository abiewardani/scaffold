package secureRandom

import (
	"crypto/rand"
	"encoding/base64"
	"math/big"

	"github.com/btcsuite/btcutil/base58"
)

var NonceBytes = 12

func RandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func RandomString(n int) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		ret[i] = letters[num.Int64()]
	}

	return string(ret), nil
}

func RandomStringURLSafe(n int) (string, error) {
	b, err := RandomBytes(n)
	return base64.URLEncoding.EncodeToString(b), err
}

func GenerateUniqueSecureToken(n int) (string, error) {
	keyBytes, err := RandomBytes(n)
	if err != nil {
		return ``, err
	}

	encoded := base58.Encode(keyBytes)
	return encoded, nil
}
