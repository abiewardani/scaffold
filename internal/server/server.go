package server

import (
	"fmt"
	"log"
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	v1 "gitlab.com/abiewardani/scaffold/internal/delivery/v1"
	"gitlab.com/abiewardani/scaffold/internal/system"
	"gitlab.com/abiewardani/scaffold/internal/system/manager"
)

type Server struct {
	httpServer *echo.Echo
	uc         manager.UsecaseManager
	sys        *system.System
}

func InitServer() *Server {
	echoServer := echo.New()
	sys := system.New()
	uc := manager.NewUsecaseManager(sys)

	echoServer.Validator = &CustomValidator{validator: validator.New()}
	return &Server{httpServer: echoServer, uc: uc, sys: sys}
}

// Run ...
func (s *Server) Run() {
	systemUserHandler := v1.NewSystemUserHandler(s.uc.SystemUser())
	systemUserGroup := s.httpServer.Group(`/system_user`)
	systemUserHandler.Mount(systemUserGroup, s.sys.Config.SecretKeyBase)

	if err := s.httpServer.Start(fmt.Sprintf(":%d", s.sys.Config.Port)); err != nil {
		log.Panic(err)
	}
}

func Start() {
	server := InitServer()
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		server.Run()
	}()

	wg.Wait()
}

// CustomValidator ...
type CustomValidator struct {
	validator *validator.Validate
}

// Validate ...
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
