package transaction

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"html/template"
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

type data struct {
	ID string
}

//func toAuthenticatePayerResponse(w http.ResponseWriter, t model.Transaction) authenticateResponse {
//	id := data{ID: strconv.FormatInt(t.ID, 10)}
//	tmp, err := template.ParseFiles("internal/views/otp.html")
//	if err != nil {
//		log.Print(err)
//	}
//
//	err2 := tmp.Execute(w, id)
//	if err2 != nil {
//		log.Print(err2)
//	}
//
//	return authenticateResponse{
//		Message: "created a transaction",
//		HTML:    "<!DOCTYPE html>\n<html>\n<body>\n<h1>Submit your OTP</h1>\n<form action=\"/authenticateTransaction/{{.ID}}\" method=\"post\">\n    <label for=\"otp\">OTP:</label>\n    <input type=\"text\" id=\"otp\" name=\"otp\"><br><br>\n    <input type=\"hidden\" id=\"trans\" name=\"trans\" value={{.trans}}><br><br>\n    <input type=\"submit\" value=\"Submit\">\n</form>\n<p>Click the \"Submit\" button and the form-data will be sent to a page on th server called \"/form\".</p>\n</body>\n</html>",
//	}
//}

func toOTPResponse(w http.ResponseWriter, t model.Transaction) {
	id := data{ID: strconv.FormatInt(t.ID, 10)}
	tmp, err := template.ParseFiles("internal/views/otp.html")
	if err != nil {
		log.Print(err)
	}

	err2 := tmp.Execute(w, id)
	if err2 != nil {
		log.Print(err2)
	}

}
