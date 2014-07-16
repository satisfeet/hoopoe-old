package httpd

import (
	"errors"
	"fmt"
	"net/http"
)

// Responds an json formatted error message.
// If error is nil then we use the status text as error message.
func Error(w http.ResponseWriter, err error, code int) {
	if err == nil {
		err = errors.New(http.StatusText(code))
	}

	http.Error(w, fmt.Sprintf(`{"error":"%s"}`, err.Error()), code)
}
