package httpd

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// Responds an json formatted error message.
//
// If error is nil then we use the status text as error message.
func Error(w http.ResponseWriter, err error, c int) {
	if err == nil {
		err = errors.New(http.StatusText(c))
	}

	http.Error(w, fmt.Sprintf(`{"error":"%s"}`, err.Error()), c)
}

// Parses a http request body to a value.
//
// If an error occurs it will be handled and returns false.
func Parse(w http.ResponseWriter, r *http.Request, v interface{}) bool {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		Error(w, nil, http.StatusBadRequest)

		return false
	}

	return true
}

// Responds value to a http request.
//
// If value is nil then we respond a json object which contains
// the status text.
// If an error occurs it returns false.
func Respond(w http.ResponseWriter, v interface{}, c int) bool {
	w.WriteHeader(c)

	if v == nil {
		v = map[string]string{
			"status": http.StatusText(c),
		}
	}

	if err := json.NewEncoder(w).Encode(v); err != nil {
		Error(w, err, http.StatusInternalServerError)

		return false
	}

	return true
}
