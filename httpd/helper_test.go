package httpd

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestError(t *testing.T) {
	Convey("Given an response and error nil", t, func() {
		res := httptest.NewRecorder()

		Convey("Error()", func() {
			Error(res, nil, http.StatusInternalServerError)

			Convey("Should set response status", func() {
				So(res.Code, ShouldEqual, http.StatusInternalServerError)
			})
			Convey("Should set response body", func() {
				So(res.Body.String(), ShouldContainSubstring, `{"error":"Internal Server Error"}`)
			})
		})
	})
	Convey("Given an response and an error", t, func() {
		res := httptest.NewRecorder()
		err := errors.New("Some error")

		Convey("Error()", func() {
			Error(res, err, http.StatusInternalServerError)

			Convey("Should set response status", func() {
				So(res.Code, ShouldEqual, http.StatusInternalServerError)
			})
			Convey("Should set response body", func() {
				So(res.Body.String(), ShouldContainSubstring, `{"error":"Some error"}`)
			})
		})
	})
}
