package system

import (
	"gitlab.com/abiewardani/scaffold/internal/system/config"
	"gitlab.com/abiewardani/scaffold/internal/system/connection"
	"gitlab.com/abiewardani/scaffold/internal/system/crypto"
	"gitlab.com/abiewardani/scaffold/pkg/storage"
)

type System struct {
	Config config.Config
	Conn   connection.Connection
	Cryp   crypto.Crypto
}

func New() *System {
	cfg := config.LoadConfig()
	conn := connection.LoadConnection(&cfg)
	crp, err := crypto.LoadCrypto(&cfg)
	if err != nil {
		return nil
	}

	// init HbStorageAWS
	_, err = storage.NewStorageAWS(&storage.StorageAwsConfig{
		AwsAccessKeyID:     cfg.AwsAccessKeyID,
		AwsSecretAccessKey: cfg.AwsSecretAccessKey,
		Region:             cfg.Region,
		Bucket:             cfg.S3Bucket,
		ExpiredTime:        cfg.S3ExpiredTime,
	})
	if err != nil {
		return nil
	}

	return &System{
		Config: cfg,
		Conn:   conn,
		Cryp:   crp,
	}
}

func (c *System) GetConn() connection.Connection {
	return c.Conn
}
