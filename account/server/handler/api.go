package handler

import (
	"net/http"

	"github.com/MihaiBlebea/coffee-shop/account/user"
)

type Logger interface {
	Info(args ...interface{})
	Trace(args ...interface{})
	Debug(args ...interface{})
	Print(args ...interface{})
	Warn(args ...interface{})
	Warning(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Panic(args ...interface{})
}

type UserService interface {
	NewUser(firstName, lastName string, age uint) (*user.User, error)
}

type Service struct {
	user   UserService
	logger Logger
}

func New(user UserService, logger Logger) *Service {
	return &Service{user, logger}
}

func (s *Service) AuthenticateEndpoint() http.Handler {
	return authenticateEndpoint(s.user, s.logger)
}

func (s *Service) HealthEndpoint() http.Handler {
	return healthEndpoint(s.logger)
}
