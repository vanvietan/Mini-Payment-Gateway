package transaction

import (
	"errors"
	"net/http"
	"pg/internal/model"
	"strconv"
)

func checkCardIDAndOrderID(r *http.Request) (int64, int64, error) {
	cardID, err := strconv.ParseInt(r.URL.Query().Get("card"), 10, 64)
	if err != nil {
		return 0, 0, errors.New("cardID must be a number")
	}
	orderID, err := strconv.ParseInt(r.URL.Query().Get("order"), 10, 64)
	if err != nil {
		return 0, 0, errors.New("orderID must be a number")
	}
	return cardID, orderID, nil
}

type authenticateResponse struct {
	Message string `json:"message"`
	HTML    string `json:"html"`
}

func toAuthenticatePayerResponse(t model.Transaction) authenticateResponse {
	return authenticateResponse{
		Message: "created a transaction",
		HTML:    "<!DOCTYPE html>\n<html>\n<body>\n<h1>Submit your OTP</h1>\n<form action=\"/authenticateTransaction/{id}\" method=\"post\">\n    <label for=\"otp\">OTP:</label>\n    <input type=\"text\" id=\"otp\" name=\"otp\"><br><br>\n    <input type=\"hidden\" id=\"trans\" name=\"trans\" value={{.trans}}><br><br>\n    <input type=\"submit\" value=\"Submit\">\n</form>\n<p>Click the \"Submit\" button and the form-data will be sent to a page on th server called \"/form\".</p>\n</body>\n</html>",
	}
}
