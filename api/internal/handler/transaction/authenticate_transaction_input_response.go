package transaction

import (
	"errors"
	"net/http"
	"regexp"
)

type AuthenticateTransaction struct {
	Message string `json:"message"`
}

func toAuthenticateTransactionResponse() AuthenticateTransaction {
	return AuthenticateTransaction{Message: "Successful Authenticated"}
}

func checkOTP(r *http.Request) (string, error) {
	err := r.ParseForm()
	if err != nil {
		return "", err
	}
	otp := r.Form.Get("otp")
	if otp == "" {
		return "", errors.New("invalid OTP")
	}
	match6 := regexp.MustCompile(`^\d{6}$`)
	if !match6.MatchString(otp) {
		return "", errors.New("invalid OTP")
	}
	return otp, nil
}
