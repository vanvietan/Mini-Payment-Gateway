package transaction

import (
	"encoding/json"
	"errors"
	"net/http"
	"pg/internal/model"
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

func (i InitAuthenticationInput) checkValidateAndMap() (model.Card, model.Order, error) {
	if i.Amount < 0 || i.Amount > maxNumber {
		return model.Card{}, model.Order{}, errors.New("invalid amount")
	}
	match16 := regexp.MustCompile(`^\d{16}$`)
	match3 := regexp.MustCompile(`^\d{3}$`)
	if !match16.MatchString(i.Number) {
		return model.Card{}, model.Order{}, errors.New("invalid number")
	}
	if !match3.MatchString(i.CVV) {
		return model.Card{}, model.Order{}, errors.New("invalid CVV")
	}
	return model.Card{
			Number:      i.Number,
			ExpiredDate: i.ExpiredDate,
			CVV:         i.CVV,
		}, model.Order{
			Amount: i.Amount,
		}, nil
}

// InitAuthenticateResponse init authenticate response
type InitAuthenticateResponse struct {
	Message string `json:"message"`
	CardID  int64  `json:"cardID"`
	OrderID int64  `json:"orderID"`
}

func checkValidationAndAmount(r *http.Request) (model.Card, model.Order, error) {
	var input InitAuthenticationInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		return model.Card{}, model.Order{}, err
	}
	cardInput, amountInput, err := input.checkValidateAndMap()
	if err != nil {
		return model.Card{}, model.Order{}, err
	}

	return cardInput, amountInput, nil
}

//	func toGetAInitAuthenticateResponse() InitAuthenticateResponse {
//		return InitAuthenticateResponse{
//			Message: "created a transaction",
//			HTML:    "<!DOCTYPE html>\n<html>\n<body>\n<h1>Submit your OTP</h1>\n<form action=\"/transactions\" method=\"post\">\n    <label for=\"otp\">OTP:</label>\n    <input type=\"text\" id=\"otp\" name=\"otp\"><br><br>\n    <input type=\"hidden\" id=\"trans\" name=\"trans\" value={{.trans}}><br><br>\n    <input type=\"submit\" value=\"Submit\">\n</form>\n<p>Click the \"Submit\" button and the form-data will be sent to a page on th server called \"/form\".</p>\n</body>\n</html>",
//		}
//	}
func toGetAInitAuthenticateResponse(c model.Card, o model.Order) InitAuthenticateResponse {
	return InitAuthenticateResponse{
		Message: "card " + c.Number + " is valid",
		CardID:  c.ID,
		OrderID: o.ID,
	}
}
