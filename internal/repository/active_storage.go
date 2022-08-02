package repository

import (
	"context"
	"errors"

	"gitlab.com/abiewardani/scaffold/internal/domain"
	"gitlab.com/abiewardani/scaffold/internal/system/connection"
	"gitlab.com/abiewardani/scaffold/internal/system/connection/database"
	"gorm.io/gorm"
)

type ActiveStorageRepository interface {
	FindOne(ctx context.Context, params domain.ActiveStorageParamsView) (*domain.ActiveStorageAttachment, error)
	GetStorageUrl(ctx context.Context, recordID, recordType string) string
}

type activeStorageCtx struct {
	DB database.GormDatabase
}

func NewActiveStorageRepository(conn connection.Connection) ActiveStorageRepository {
	return &activeStorageCtx{
		DB: conn.DB(),
	}
}

func (c *activeStorageCtx) FindOne(ctx context.Context, params domain.ActiveStorageParamsView) (*domain.ActiveStorageAttachment, error) {
	res := domain.ActiveStorageAttachment{}
	db := c.DB.Slave()
	var err error

	if params.RecordID != `` {
		db = db.Where(`record_id = ?`, params.RecordID)
	}

	if params.RecordType != `` {
		db = db.Where(`record_type = ?`, params.RecordType)
	}

	err = db.Preload("ActiveStorageBlob").First(&res).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &res, nil
}

func (c *activeStorageCtx) GetStorageUrl(ctx context.Context, recordID, recordType string) string {
	res := domain.ActiveStorageAttachment{}
	db := c.DB.Slave().Where(`record_id = ?`, recordID).Where(`record_type = ?`, recordType)

	err := db.Preload("ActiveStorageBlob").First(&res).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ``
		}
		return ``
	}

	return res.GetStorageUrl()
}
