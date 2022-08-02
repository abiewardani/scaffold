package lockbox

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
)

type KmsAws interface {
	Decrypt(chipperText string, encryptionContext map[string]*string) (plainText []byte, err error)
	Encrypt(message []byte, encryptionContext map[string]*string) (cipherText string, err error)
}

type kmsAwsCtx struct {
	Session  *session.Session
	SvcKms   *kms.KMS
	KmsKeyID string
}

type KmsAwsConfig struct {
	AwsAccessKeyID     string
	AwsSecretAccessKey string
	KmsKeyID           string
	Region             string
}

func NewKmsAws(config *KmsAwsConfig) (KmsAws, error) {
	kmsAws := kmsAwsCtx{}
	var err error

	kmsAws.Session, err = session.NewSession(&aws.Config{
		Region:      aws.String(config.Region),
		Credentials: credentials.NewStaticCredentials(config.AwsAccessKeyID, config.AwsSecretAccessKey, ""),
	})
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	kmsAws.SvcKms = kms.New(kmsAws.Session)
	kmsAws.KmsKeyID = config.KmsKeyID
	return &kmsAws, nil
}

func (c *kmsAwsCtx) Decrypt(chipperText string, encryptionContext map[string]*string) (plainText []byte, err error) {
	chipperTextArr := strings.Split(chipperText, `:`)
	if len(chipperTextArr) < 2 {
		return nil, errors.New(`Invalid Chippertext`)
	}

	chipperText = chipperTextArr[1]
	ciphertextBlob, err := base64.StdEncoding.DecodeString(chipperText)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	inputDecrypt := &kms.DecryptInput{
		CiphertextBlob:    ciphertextBlob,
		EncryptionContext: encryptionContext,
	}

	respDecrypt, err := c.SvcKms.Decrypt(inputDecrypt)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return respDecrypt.Plaintext, nil
}

func (c *kmsAwsCtx) Encrypt(message []byte, encryptionContext map[string]*string) (cipherText string, err error) {
	inputEncrypt := &kms.EncryptInput{
		KeyId:             aws.String(c.KmsKeyID),
		Plaintext:         []byte(message),
		EncryptionContext: encryptionContext,
	}

	respEncrypt, err := c.SvcKms.Encrypt(inputEncrypt)
	if err != nil {
		fmt.Println(err.Error())
		return ``, err
	}

	dst := make([]byte, base64.StdEncoding.EncodedLen(len(respEncrypt.CiphertextBlob)))
	base64.StdEncoding.Encode(dst, respEncrypt.CiphertextBlob)

	cipherText = `v1:` + string(dst)
	return cipherText, nil
}
