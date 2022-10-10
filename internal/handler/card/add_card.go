package card

import (
	"encoding/json"
	"net/http"
	"pg/internal/handler/common"
	"pg/internal/model"
)

// AddCard add a card
func (h Handler) AddCard(w http.ResponseWriter, r *http.Request) {
	reqBody, err := checkValidation(r)
	if err != nil {
		common.ResponseJSON(w, http.StatusBadRequest, common.CommonErrorResponse{
			Code:        "invalid_request",
			Description: err.Error(),
		})
		return
	}
	cardS, errS := h.CardSvc.AddCard(r.Context(), reqBody)
	if errS != nil {
		common.ResponseJSON(w, http.StatusInternalServerError, common.InternalCommonErrorResponse)
		return
	}
	common.ResponseJSON(w, http.StatusOK, toGetACardResponse(cardS))
}

func checkValidation(r *http.Request) (model.Card, error) {
	var input ACardInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		return model.Card{}, err
	}
	cardInput, err := input.validateAndMap()
	if err != nil {
		return model.Card{}, err
	}
	return cardInput, nil
}
