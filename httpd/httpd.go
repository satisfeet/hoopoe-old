package httpd

import (
	"net/http"

	"gopkg.in/mgo.v2"

	"github.com/satisfeet/go-validation"
	"github.com/satisfeet/hoopoe/store/mongo"
)

func ErrorCode(err error) int {
	switch err {
	case mgo.ErrNotFound:
		return http.StatusNotFound
	case mongo.ErrBadId:
		return http.StatusBadRequest
	}

	switch err.(type) {
	case validation.Error:
		return http.StatusBadRequest
	}

	return http.StatusInternalServerError
}
