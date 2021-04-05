package handler

import (
	"net/http"

	"github.com/MihaiBlebea/coffee-shop/bonus/stamp"
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

type StampService interface {
	GetStampsForUser(userID string) ([]stamp.Stamp, error)
}

type Service struct {
	stamp  StampService
	logger Logger
}

func New(stamp StampService, logger Logger) *Service {
	return &Service{stamp, logger}
}

func (s *Service) BonusEndpoint() http.Handler {
	return bonusEndpoint(s.stamp, s.logger)
}

func (s *Service) HealthEndpoint() http.Handler {
	return healthEndpoint(s.logger)
}
