package card

import (
	"encoding/json"
	"net/http"
	common2 "pg/api/internal/handler/common"
	"pg/api/internal/model"
)

// AddCard add a card
func (h Handler) AddCard(w http.ResponseWriter, r *http.Request) {
	reqBody, err := checkValidation(r)
	if err != nil {
		common2.ResponseJSON(w, http.StatusBadRequest, common2.CommonErrorResponse{
			Code:        "invalid_request",
			Description: err.Error(),
		})
		return
	}
	cardS, errS := h.CardSvc.AddCard(r.Context(), reqBody)
	if errS != nil {
		common2.ResponseJSON(w, http.StatusInternalServerError, common2.InternalCommonErrorResponse)
		return
	}
	common2.ResponseJSON(w, http.StatusOK, toGetACardResponse(cardS))
}

func checkValidation(r *http.Request) (model.Card, error) {
	var input ACardInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		return model.Card{}, err
	}
	cardInput, err := input.ValidateAndMap()
	if err != nil {
		return model.Card{}, err
	}
	return cardInput, nil
}
