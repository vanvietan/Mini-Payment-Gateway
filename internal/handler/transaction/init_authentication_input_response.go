package transaction

import (
	"encoding/json"
	"errors"
	"math"
	"net/http"
	"pg/internal/model"
	"regexp"
	"time"
)

// InitAuthenticationInput contains card information and amount
type InitAuthenticationInput struct {
	Number      string    `json:"number"`
	ExpiredDate time.Time `json:"expired_date"`
	CVV         string    `json:"CVV"`
	UserID      int64     `json:"userID"`
	Balance     int64     `json:"balance"`
	Amount      int64     `json:"amount"`
}

func (c InitAuthenticationInput) checkValidate() (model.Card, model.Order, error) {
	if c.Amount <= 0 {
		return model.Card{}, model.Order{}, errors.New("invalid amount")
	}
	if c.UserID <= 0 || c.UserID > math.MaxInt64 {
		return model.Card{}, model.Order{}, errors.New("invalid userID")
	}
	match16 := regexp.MustCompile(`^\d{16}$`)
	match3 := regexp.MustCompile(`^\d{3}$`)
	if !match16.MatchString(c.Number) {
		return model.Card{}, model.Order{}, errors.New("invalid number")
	}
	if !match3.MatchString(c.CVV) {
		return model.Card{}, model.Order{}, errors.New("invalid CVV")
	}
	if c.ExpiredDate.Equal(time.Now()) || c.ExpiredDate.Before(time.Now()) {
		return model.Card{}, model.Order{}, errors.New("invalid expired date")
	}
	if c.Balance <= 0 {
		return model.Card{}, model.Order{}, errors.New("invalid balance")
	}
	if c.Balance < c.Amount {
		return model.Card{}, model.Order{}, errors.New("balance is too low")
	}
	return model.Card{
			Number:      c.Number,
			ExpiredDate: c.ExpiredDate,
			CVV:         c.CVV,
			Balance:     c.Balance,
			UserID:      c.UserID,
		}, model.Order{
			Amount: c.Amount,
		}, nil
}

// OTPResponse OTP response
type OTPResponse struct {
	Message string `json:"message"`
	OTP     string `json:"otp"`
}

func checkValidationAndAmount(r *http.Request) (model.Card, model.Order, error) {
	var input InitAuthenticationInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		return model.Card{}, model.Order{}, err
	}
	cardInput, amountInput, err := input.checkValidate()
	if err != nil {
		return model.Card{}, model.Order{}, err
	}

	return cardInput, amountInput, nil
}
