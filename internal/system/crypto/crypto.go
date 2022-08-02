package crypto

import (
	"fmt"

	"gitlab.com/abiewardani/scaffold/internal/system/config"
	"gitlab.com/abiewardani/scaffold/pkg/lockbox"
	secureRandom "gitlab.com/abiewardani/scaffold/pkg/secure_random"
)

var HbCrypto Crypto

// Connection ...
type Crypto struct {
	KmsAws     lockbox.KmsAws
	AesGcm     lockbox.AesGcm
	BlindIndex lockbox.BlindIndex
}

// LoadConnection ...
func LoadCrypto(cfg *config.Config) (Crypto, error) {
	cryptoObj := Crypto{}
	kmsAws, err := lockbox.NewKmsAws(&lockbox.KmsAwsConfig{
		AwsAccessKeyID:     cfg.AwsAccessKeyID,
		AwsSecretAccessKey: cfg.AwsSecretAccessKey,
		KmsKeyID:           cfg.KmsKeyID,
		Region:             cfg.Region,
	})

	if err != nil {
		return cryptoObj, err
	}

	cryptoObj.KmsAws = kmsAws
	cryptoObj.AesGcm = lockbox.NewAesGcm()

	//BlindIndex
	cryptoObj.BlindIndex = lockbox.NewBlindIndex(cfg.BlindIndexKey)
	HbCrypto = cryptoObj
	return cryptoObj, nil
}

func (c *Crypto) Decrypt(modelName, modelID, kmsKey, ciphertext string) (plainText string) {
	if ciphertext == `` || kmsKey == `` {
		return ``
	}

	encryptionContext := map[string]*string{
		"model_name": &modelName,
		"model_id":   &modelID,
	}

	key, err := c.KmsAws.Decrypt(kmsKey, encryptionContext)
	if err != nil {
		return ``
	}

	plainTextStr := c.AesGcm.Decrypt(ciphertext, key)
	return plainTextStr
}

func (c *Crypto) EncryptKmsKey(modelName, modelID string, keyRandom []byte) (encryptedKmsKey string, err error) {
	encryptionContext := map[string]*string{
		"model_name": &modelName,
		"model_id":   &modelID,
	}

	encryptedKmsKey, err = c.KmsAws.Encrypt(keyRandom, encryptionContext)
	if err != nil {
		return ``, err
	}

	return encryptedKmsKey, nil
}

func (c *Crypto) DecryptKmsKey(modelName, modelID, chipperText string) (keyRandom []byte, err error) {
	encryptionContext := map[string]*string{
		"model_name": &modelName,
		"model_id":   &modelID,
	}

	keyRandom, err = c.KmsAws.Decrypt(chipperText, encryptionContext)
	if err != nil {
		return nil, err
	}

	return keyRandom, nil
}

func (c *Crypto) EncryptPlainText(modelName, modelID, plaintext string, keyRandom []byte) (cipherText string, err error) {
	plainTextStr, err := c.AesGcm.Encrypt([]byte(plaintext), keyRandom)
	if err != nil {
		fmt.Println(err)
		return ``, err
	}

	return plainTextStr, nil
}

func (c *Crypto) GenerateKey() (randomKey []byte, err error) {
	keyRandomParams, err := secureRandom.RandomBytes(32)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return keyRandomParams, nil
}

func (c *Crypto) GenerateBlindIndex(params string) string {
	return c.BlindIndex.Generate(params)
}
