package transaction

import (
	"net/http"
	"text/template"
)

// OTPForm form for clients
func (h Handler) OTPForm(w http.ResponseWriter, r *http.Request) {

	tmp, _ := template.ParseFiles("internal/views/otp.html")

	tmp.Execute(w, nil)
}
