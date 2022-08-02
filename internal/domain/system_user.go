package domain

import (
	"time"

	"gitlab.com/abiewardani/scaffold/internal/system/crypto"
)

var (
	UnverifiedStatus = 0
	VerifiedStatus   = 1
	AdminGenerated   = 2
)

var (
	PasswordStatusUserGenerated  = 0
	PasswordStatusAdminGenerated = 1
)

var (
	StatusActive   = 0
	StatusArchived = 1
)

var IsValidEmailStatus = map[string]int{
	"unverified":      UnverifiedStatus,
	"verified":        VerifiedStatus,
	"admin_generated": AdminGenerated,
}

var GetEmailStatus = map[string]int{
	"unverified":      UnverifiedStatus,
	"verified":        VerifiedStatus,
	"admin_generated": AdminGenerated,
}

var IsValidPasswordStatus = map[int]bool{
	PasswordStatusUserGenerated:  true,
	PasswordStatusAdminGenerated: true,
}

var IsValidStatus = map[int]bool{
	StatusActive:   true,
	StatusArchived: true,
}

type SystemUser struct {
	ID                 string    `gorm:"column:id;primary_key;type:uuid;primary_key;"`
	FirstName          string    `gorm:"column:first_name"`
	LastName           string    `gorm:"column:last_name"`
	CreatedAt          time.Time `gorm:"column:created_at"`
	UpdatedAt          time.Time `gorm:"column:updated_at"`
	Username           string    `gorm:"column:username"`
	EncryptedKmsKey    string    `gorm:"column:encrypted_kms_key"`
	EmailCiphertext    string    `gorm:"column:email_ciphertext"`
	EmailBidx          string    `gorm:"column:email_bidx"`
	MobileNoCiphertext string    `gorm:"mobile_no_ciphertext"`
	MobileNoBidx       string    `gorm:"mobile_no_bidx"`

	RandomKey []byte `gorm:"-"`
}

func (c *SystemUser) TableName() string {
	return `system_users`
}

func (c *SystemUser) MobileNo() string {
	return crypto.HbCrypto.Decrypt(ModelName, c.ID, c.EncryptedKmsKey, c.MobileNoCiphertext)
}

func (c *SystemUser) Email() string {
	return crypto.HbCrypto.Decrypt(ModelName, c.ID, c.EncryptedKmsKey, c.EmailCiphertext)
}

func (c *SystemUser) GetRandomKey() []byte {
	// check random key is already generate
	if c.RandomKey != nil {
		return c.RandomKey
	}

	// check if we have previous EncryptedKmsKey
	if c.EncryptedKmsKey != `` {
		c.RandomKey, _ = crypto.HbCrypto.DecryptKmsKey(ModelName, c.ID, c.EncryptedKmsKey)
		return c.RandomKey
	}

	// Generate Random Key
	c.RandomKey, _ = crypto.HbCrypto.GenerateKey()
	c.EncryptedKmsKey, _ = crypto.HbCrypto.EncryptKmsKey(ModelName, c.ID, c.RandomKey)
	return c.RandomKey
}

func (c *SystemUser) EmailCrypto(email string) error {
	randomKey := c.GetRandomKey()
	chipperText, err := crypto.HbCrypto.EncryptPlainText(ModelName, c.ID, email, randomKey)
	if err != nil {
		return err
	}
	c.EmailCiphertext = chipperText
	c.EmailBidx = crypto.HbCrypto.GenerateBlindIndex(email)
	return nil
}

func (c *SystemUser) MobileNoCrypto(mobileNo string) error {
	chipperText, err := crypto.HbCrypto.EncryptPlainText(ModelName, c.ID, mobileNo, c.GetRandomKey())
	if err != nil {
		return err
	}
	c.MobileNoCiphertext = chipperText
	c.MobileNoBidx = crypto.HbCrypto.GenerateBlindIndex(mobileNo)
	return nil
}

type SystemUserParams struct {
	ID       string
	Username string
	Email    string
}

var ModelName = `SystemUser`
