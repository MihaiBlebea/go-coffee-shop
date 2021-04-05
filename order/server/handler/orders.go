package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/MihaiBlebea/coffee-shop/order/payment"
)

type OrdersRequest struct {
	UserID string
}

type OrdersResponse struct {
	Orders  []payment.Payment `json:"orders,omitempty"`
	Success bool              `json:"success"`
	Message string            `json:"message,omitempty"`
}

func ordersEndpoint(paymentService PaymentService, logger Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := OrdersResponse{}

		request, err := validateOrdersRequest(r)
		if err != nil {
			response.Message = err.Error()
			sendResponse(w, response, http.StatusBadRequest)
			return
		}

		fmt.Println(request.UserID)
		payments, err := paymentService.AllByUserID(request.UserID)
		if err != nil {
			response.Message = err.Error()
			sendResponse(w, response, http.StatusBadRequest)
			return
		}

		response.Orders = payments
		response.Success = true

		sendResponse(w, response, 200)
	})
}

func validateOrdersRequest(r *http.Request) (*OrderRequest, error) {
	request := OrderRequest{}

	userID := r.URL.Query().Get("user_id")

	if userID == "" || len(userID) < 3 {
		return &request, errors.New("Invalid request param user_id")
	}

	request.UserID = userID

	return &request, nil
}
