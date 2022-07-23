package signup

import (
	"net/http"

	"github.com/kchang4/auth/internal/data"
)

func SignupHandler(rw http.ResponseWriter, r *http.Request) {
	if _, ok := r.Header["Email"]; !ok {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Email Missing"))
		return
	}

	if _, ok := r.Header["Username"]; !ok {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Username Missing"))
		return
	}

	if _, ok := r.Header["Passwordhash"]; !ok {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Passwordhas Missing"))
		return
	}

	if _, ok := r.Header["Fullname"]; !ok {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Fullname Missing"))
	}

	// Validate and then add user
	check := data.AddUser(r.Header["Email"][0], r.Header[""])
}
