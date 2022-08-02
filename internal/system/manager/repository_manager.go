package manager

import (
	"sync"

	"gitlab.com/abiewardani/scaffold/internal/repository"
	"gitlab.com/abiewardani/scaffold/internal/system"
	"gitlab.com/abiewardani/scaffold/internal/system/connection"
)

type RepositoryManager interface {
	SystemUser() repository.SystemUserRepository
	ActiveStorage() repository.ActiveStorageRepository
}

// repositoryManager ..
type repositoryManager struct {
	sys  system.System
	conn connection.Connection
}

func NewRepositoryManager(sys *system.System) RepositoryManager {
	var r repositoryManager
	var once sync.Once

	once.Do(func() {
		r = repositoryManager{conn: sys.GetConn(), sys: *sys}
	})

	return r
}

func (r repositoryManager) SystemUser() repository.SystemUserRepository {
	var once sync.Once
	var systemUserRepo repository.SystemUserRepository

	once.Do(func() {
		systemUserRepo = repository.NewSystemUserRepository(r.conn)
	})

	return systemUserRepo
}

func (r repositoryManager) ActiveStorage() repository.ActiveStorageRepository {
	var activeStorageRepository repository.ActiveStorageRepository
	var once sync.Once

	once.Do(func() {
		activeStorageRepository = repository.NewActiveStorageRepository(r.conn)
	})

	return activeStorageRepository
}
