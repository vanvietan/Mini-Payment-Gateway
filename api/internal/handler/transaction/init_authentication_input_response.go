package transaction

import (
	"encoding/json"
	"errors"
	"net/http"
	model2 "pg/api/internal/model"
	"regexp"
	"time"
)

const maxNumber = 9999999999

// InitAuthenticationInput contains card information and amount
type InitAuthenticationInput struct {
	Number      string    `json:"number"`
	ExpiredDate time.Time `json:"expired_date"`
	CVV         string    `json:"CVV"`
	Amount      int64     `json:"amount"`
}

func (i InitAuthenticationInput) checkValidateAndMap() (model2.Card, model2.Order, error) {
	if i.Amount < 0 || i.Amount > maxNumber {
		return model2.Card{}, model2.Order{}, errors.New("invalid amount")
	}
	match16 := regexp.MustCompile(`^\d{16}$`)
	match3 := regexp.MustCompile(`^\d{3}$`)
	if !match16.MatchString(i.Number) {
		return model2.Card{}, model2.Order{}, errors.New("invalid number")
	}
	if !match3.MatchString(i.CVV) {
		return model2.Card{}, model2.Order{}, errors.New("invalid CVV")
	}
	if i.ExpiredDate.Before(time.Now()) {
		return model2.Card{}, model2.Order{}, errors.New("invalid expired date")
	}
	return model2.Card{
			Number:      i.Number,
			ExpiredDate: i.ExpiredDate,
			CVV:         i.CVV,
		}, model2.Order{
			Amount: i.Amount,
		}, nil
}

// InitAuthenticateResponse init authenticate response
type InitAuthenticateResponse struct {
	Message string `json:"message"`
	CardID  int64  `json:"cardID"`
	OrderID int64  `json:"orderID"`
}

func checkValidationAndAmount(r *http.Request) (model2.Card, model2.Order, error) {
	var input InitAuthenticationInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		return model2.Card{}, model2.Order{}, err
	}
	cardInput, amountInput, err := input.checkValidateAndMap()
	if err != nil {
		return model2.Card{}, model2.Order{}, err
	}

	return cardInput, amountInput, nil
}

//	func toGetAInitAuthenticateResponse() InitAuthenticateResponse {
//		return InitAuthenticateResponse{
//			Message: "created a transaction",
//			HTML:    "<!DOCTYPE html>\n<html>\n<body>\n<h1>Submit your OTP</h1>\n<form action=\"/transactions\" method=\"post\">\n    <label for=\"otp\">OTP:</label>\n    <input type=\"text\" id=\"otp\" name=\"otp\"><br><br>\n    <input type=\"hidden\" id=\"trans\" name=\"trans\" value={{.trans}}><br><br>\n    <input type=\"submit\" value=\"Submit\">\n</form>\n<p>Click the \"Submit\" button and the form-data will be sent to a page on th server called \"/form\".</p>\n</body>\n</html>",
//		}
//	}
func toGetAInitAuthenticateResponse(c model2.Card, o model2.Order) InitAuthenticateResponse {
	return InitAuthenticateResponse{
		Message: "card " + c.Number + " is valid",
		CardID:  c.ID,
		OrderID: o.ID,
	}
}
