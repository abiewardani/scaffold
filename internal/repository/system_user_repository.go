package repository

import (
	"context"
	"errors"

	"gitlab.com/abiewardani/scaffold/internal/domain"
	"gitlab.com/abiewardani/scaffold/internal/system/connection"
	"gitlab.com/abiewardani/scaffold/internal/system/connection/database"
	"gorm.io/gorm"
)

// SystemUserRepository ...
type SystemUserRepository interface {
	FindOne(ctx context.Context, params *domain.SystemUserParams) (sysUser *domain.SystemUser, err error)
	Update(ctx context.Context, systemUser *domain.SystemUser) (sysUser *domain.SystemUser, err error)
	Create(ctx context.Context, systemUser *domain.SystemUser) (sysUser *domain.SystemUser, err error)
}

// sytemUserRepositoryCtx ...
type sytemUserRepositoryCtx struct {
	DB database.GormDatabase
}

func NewSystemUserRepository(conn connection.Connection) SystemUserRepository {
	return &sytemUserRepositoryCtx{
		DB: conn.DB(),
	}
}

func (c *sytemUserRepositoryCtx) FindOne(ctx context.Context, params *domain.SystemUserParams) (sysUser *domain.SystemUser, err error) {
	res := domain.SystemUser{}
	db := c.DB.Slave()

	if params.ID != `` {
		db = db.Where(`system_users.id = ?`, params.ID)
	}

	err = db.First(&res).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &res, nil
}

func (c *sytemUserRepositoryCtx) Update(ctx context.Context, systemUser *domain.SystemUser) (sysUser *domain.SystemUser, err error) {
	db := c.DB.Master()

	err = db.Model(&systemUser).Updates(systemUser).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return systemUser, nil
}

func (c *sytemUserRepositoryCtx) Create(ctx context.Context, systemUser *domain.SystemUser) (sysUser *domain.SystemUser, err error) {
	if err := c.DB.Master().Create(systemUser).Error; err != nil {
		return nil, err
	}

	return systemUser, nil
}
