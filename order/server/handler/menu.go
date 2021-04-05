package handler

import (
	"net/http"

	"github.com/MihaiBlebea/coffee-shop/order/coffee"
)

type MenuResponse struct {
	Coffees []coffee.Coffee `json:"coffees,omitempty"`
	Success bool            `json:"success"`
	Message string          `json:"message,omitempty"`
}

func menuEndpoint(coffeeService CoffeeService, logger Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := MenuResponse{}

		response.Coffees = coffeeService.GetAll()
		response.Success = true

		sendResponse(w, response, 200)
	})
}
