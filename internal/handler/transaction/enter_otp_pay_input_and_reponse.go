package transaction

import (
	"errors"
	"net/http"
)

// InputOTP input otp
//type InputOTP struct {
//	OTP string `json:"OTP"`
//}

// PayResponse message of successful transaction
type PayResponse struct {
	Message string `json:"message"`
	Number  string `json:"number"`
	Balance int64  `json:"balance"`
}

func checkInputOTP(r *http.Request) (string, error) {
	value := r.FormValue("otp")
	if value == "" {
		return "", errors.New("invalid OTP")
	}

	return value, nil
}
