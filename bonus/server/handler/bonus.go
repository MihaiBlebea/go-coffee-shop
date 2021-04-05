package handler

import (
	"errors"
	"net/http"
)

type BonusRequest struct {
	UserID string `json:"user_id"`
}

type BonusResponse struct {
	StampCount uint   `json:"stamp_count"`
	BonusCount uint   `json:"bonus_count"`
	Success    bool   `json:"success"`
	Message    string `json:"message,omitempty"`
}

func bonusEndpoint(stamp StampService, logger Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := BonusResponse{}

		request, err := validateBonusRequest(r)
		if err != nil {
			response.Message = err.Error()
			sendResponse(w, response, http.StatusBadRequest)
			return
		}

		stamps, err := stamp.GetStampsForUser(request.UserID)
		if err != nil {
			response.Message = err.Error()
			sendResponse(w, response, http.StatusBadRequest)
			return
		}

		response.StampCount = uint(len(stamps))
		response.BonusCount = uint(len(stamps))
		response.Success = true

		sendResponse(w, response, 200)
	})
}

func validateBonusRequest(r *http.Request) (*BonusRequest, error) {
	request := BonusRequest{}

	userID := r.URL.Query().Get("user_id")

	if userID == "" || len(userID) < 3 {
		return &request, errors.New("Invalid request param user_id")
	}

	request.UserID = userID

	return &request, nil
}
