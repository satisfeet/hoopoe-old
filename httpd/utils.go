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

// Responds value to a http request.
//
// If value is nil then we respond a json object which contains
// the status text.
func Respond(w http.ResponseWriter, v interface{}, c int) {
	w.WriteHeader(c)

	if v == nil {
		v = map[string]string{
			"status": http.StatusText(c),
		}
	}

	if err := json.NewEncoder(w).Encode(v); err != nil {
		Error(w, err, http.StatusInternalServerError)
	}
}
