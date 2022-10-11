package transaction

import (
	"encoding/json"
	"errors"
	"net/http"
	"pg/internal/handler/card"
	"pg/internal/model"
)

// AmountInput input from order amount
type AmountInput struct {
	Amount int64 `json:"amount"`
}

// OTPResponse OTP response
type OTPResponse struct {
	Message string `json:"message"`
	OTP     string `json:"otp"`
}

func (a AmountInput) validateAmount() (model.Order, error) {
	if a.Amount <= 0 {
		return model.Order{}, errors.New("invalid amount")
	}
	return model.Order{
		Amount: a.Amount,
	}, nil
}

func checkValidationAndAmount(r *http.Request) (model.Card, model.Order, error) {
	var input card.ACardInput
	var amount AmountInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		return model.Card{}, model.Order{}, err
	}
	cardInput, err := input.ValidateAndMap()
	if err != nil {
		return model.Card{}, model.Order{}, err
	}
	amountInput, errA := amount.validateAmount()
	if errA != nil {
		return model.Card{}, model.Order{}, err
	}

	return cardInput, amountInput, nil
}
