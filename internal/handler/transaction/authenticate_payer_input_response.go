package transaction

import (
	"errors"
	"net/http"
	"pg/internal/handler/common"
	"regexp"
)

type AuthenticatePayerResponse struct {
	Message string `json:"message"`
}

func toAuthenticatePayerResponse() AuthenticatePayerResponse {
	return AuthenticatePayerResponse{Message: "Successful Authenticated"}
}

func toWrongOTPResponse() common.CommonErrorResponse {
	return common.CommonErrorResponse{
		Code:        "internal_server_error",
		Description: "Incorrect OTP authenticate",
	}
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
