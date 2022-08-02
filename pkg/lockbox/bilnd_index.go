package lockbox

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

type BlindIndex interface {
	Generate(value string) string
}

type blindIndexCtx struct {
	MasterKey string
}

func NewBlindIndex(masterKey string) BlindIndex {
	return &blindIndexCtx{
		MasterKey: masterKey,
	}
}

func (c *blindIndexCtx) Generate(value string) string {
	if value == `` {
		return ``
	}
	decodeKey, err := hex.DecodeString(c.MasterKey)
	if err != nil {
		fmt.Println(err)
		return ``
	}

	key := value + string(decodeKey)
	h := sha1.New()
	h.Write([]byte(value + key))
	blindIndex := hex.EncodeToString(h.Sum(nil))

	return blindIndex
}
