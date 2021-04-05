package handler

import (
	"net/http"

	"github.com/MihaiBlebea/coffee-shop/order/coffee"
	"github.com/MihaiBlebea/coffee-shop/order/payment"
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

type PaymentService interface {
	NewPayment(userID, coffeeID string, amount uint) (*payment.Payment, error)
	AllByUserID(userID string) ([]payment.Payment, error)
}

type CoffeeService interface {
	GetAll() []coffee.Coffee
}

type Service struct {
	paymentService PaymentService
	coffeeService  CoffeeService
	logger         Logger
}

func New(paymentService PaymentService, coffeeService CoffeeService, logger Logger) *Service {
	return &Service{paymentService, coffeeService, logger}
}

func (s *Service) OrderEndpoint() http.Handler {
	return orderEndpoint(s.paymentService, s.logger)
}

func (s *Service) OrdersEndpoint() http.Handler {
	return ordersEndpoint(s.paymentService, s.logger)
}

func (s *Service) HealthEndpoint() http.Handler {
	return healthEndpoint(s.logger)
}

func (s *Service) MenuEndpoint() http.Handler {
	return menuEndpoint(s.coffeeService, s.logger)
}
