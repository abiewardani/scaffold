package usecase

import (
	"context"

	"gitlab.com/abiewardani/scaffold/internal/domain"
	"gitlab.com/abiewardani/scaffold/internal/repository"
	"gitlab.com/abiewardani/scaffold/internal/system"
	serializer "gitlab.com/abiewardani/scaffold/internal/usecase/serializers"
	"gitlab.com/abiewardani/scaffold/pkg/shared"
)

type SystemUser interface {
	Show(ctx context.Context, id string) (sysUser *serializer.SystemUserShortAttributesSerializer, err error)
}

// systemUserCtx ..
type systemUserCtx struct {
	sys                  system.System
	systemUserRepository repository.SystemUserRepository
	activeStorageRepo    repository.ActiveStorageRepository
}

func NewSystemUserUc(sys system.System, systemUserRepo repository.SystemUserRepository, activeStorageRepository repository.ActiveStorageRepository) SystemUser {
	return &systemUserCtx{
		sys:                  sys,
		systemUserRepository: systemUserRepo,
		activeStorageRepo:    activeStorageRepository,
	}
}

func (c *systemUserCtx) Show(ctx context.Context, id string) (sysUser *serializer.SystemUserShortAttributesSerializer, err error) {
	systemUser, err := c.systemUserRepository.FindOne(ctx, &domain.SystemUserParams{
		ID: id,
	})

	if err != nil {
		return nil, shared.NewMultiStringBadRequestError(shared.HTTPErrorBadRequest, "An error occurred while getting system user")
	}

	if systemUser == nil {
		return nil, shared.NewMultiStringBadRequestError(shared.HTTPErrorDataNotFound, "Data not found")
	}

	return serializer.NewSystemUserShortAttributesSerializer(systemUser), nil
}
