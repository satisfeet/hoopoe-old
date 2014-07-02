package httpd

import (
	"encoding/json"
	"errors"
	"net/http"
)

type Handler struct {
	response http.ResponseWriter
	request  *http.Request
}

func NewHandler(w http.ResponseWriter, r *http.Request) *Handler {
	return &Handler{w, r}
}

func (h *Handler) Parse(v interface{}) error {
	err := json.NewDecoder(h.request.Body).Decode(v)

	return err
}

func (h *Handler) Respond(v interface{}, s int) {
	if v == nil {
		h.response.WriteHeader(s)

		return
	}

	j, err := json.Marshal(&v)

	if err != nil {
		h.Error(err, 500)

		return
	}

	h.response.WriteHeader(s)
	h.response.Write(j)
}

func (h *Handler) Error(e error, s int) {
	if e == nil {
		e = errors.New(http.StatusText(s))
	}

	h.response.WriteHeader(s)
	h.response.Write([]byte(e.Error()))
}
