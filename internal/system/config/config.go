package config

import (
	"os"
	"strconv"
	"time"

	"gitlab.com/abiewardani/scaffold/internal/system/config/environment"
)

type Config struct {
	Environment             environment.Environment
	Port                    int
	AppName                 string
	GracefulShutdownSeconds int

	DbMasterHost     string
	DbMasterPort     int
	DbMasterUser     string
	DbMasterPassword string
	DbMasterName     string

	DbSlaveHost     string
	DbSlavePort     int
	DbSlaveUser     string
	DbSlavePassword string
	DbSlaveName     string

	SecretKeyBase          string
	AwsAccessKeyID         string
	AwsSecretAccessKey     string
	KmsKeyID               string
	Region                 string
	PrivateGateWayEndpoint string

	S3Bucket      string
	S3ExpiredTime int

	BlindIndexKey string

	TokenName              string
	TokenType              string
	TokenTemporaryLifetime time.Duration
	AccessTokenLifeTime    time.Duration
}

// LoadConfig ...
func LoadConfig() Config {
	var err error
	conf := Config{}
	if port, err := strconv.Atoi(os.Getenv("PORT")); err == nil {
		conf.Port = port
	}

	conf.DbMasterHost = os.Getenv("DB_MASTER_HOST")
	if port, err := strconv.Atoi(os.Getenv("DB_MASTER_PORT")); err == nil {
		conf.DbMasterPort = port
	}

	conf.DbMasterName = os.Getenv("DB_MASTER_NAME")
	conf.DbMasterUser = os.Getenv("DB_MASTER_USER")
	conf.DbMasterPassword = os.Getenv("DB_MASTER_PASSWORD")

	conf.DbSlaveHost = os.Getenv("DB_SLAVE_HOST")
	if port, err := strconv.Atoi(os.Getenv("DB_SLAVE_PORT")); err == nil {
		conf.DbSlavePort = port
	}
	conf.DbSlaveName = os.Getenv("DB_SLAVE_NAME")
	conf.DbSlaveUser = os.Getenv("DB_SLAVE_USER")
	conf.DbSlavePassword = os.Getenv("DB_SLAVE_PASSWORD")

	conf.SecretKeyBase = os.Getenv("SECRET_KEY_BASE")
	conf.AwsAccessKeyID = os.Getenv("AWS_ACCESS_KEY_ID")
	conf.AwsSecretAccessKey = os.Getenv("AWS_SECRET_ACCESS_KEY")
	conf.Region = os.Getenv("REGION")
	conf.KmsKeyID = os.Getenv("KMS_KEY_ID")

	conf.BlindIndexKey = os.Getenv("BLIND_INDEX_KEY")
	conf.S3Bucket = os.Getenv("S3_BUCKET")

	if s3ExpiredTime, err := strconv.Atoi(os.Getenv("S3_EXPIRED_TIME")); err != nil {
		conf.S3ExpiredTime = s3ExpiredTime
	}

	conf.TokenTemporaryLifetime = 72 * time.Second
	temporaryLifetimeStr := os.Getenv("TOKEN_TEMPORARY_LIFETIME")
	temporaryLifetime, err := strconv.Atoi(temporaryLifetimeStr)
	if err == nil {
		conf.TokenTemporaryLifetime = time.Duration(temporaryLifetime) * time.Second
	}

	conf.AccessTokenLifeTime = 300 * time.Second
	accessTokenLifetimeStr := os.Getenv("ACCESS_TOKEN_LIFETIME")
	accessTokenLifetime, err := strconv.Atoi(accessTokenLifetimeStr)
	if err != nil {
		conf.AccessTokenLifeTime = time.Duration(accessTokenLifetime) * time.Hour
	}

	return conf
}
