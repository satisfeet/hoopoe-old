package context

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	. "github.com/smartystreets/goconvey/convey"
)

func TestParam(t *testing.T) {
	Convey("Given a Request with parameter", t, func() {
		_, _, ctx := request("GET", "/bar", "")

		Convey("Param()", func() {
			value := ctx.Param("foo")

			Convey("Should return parameter", func() {
				So(value, ShouldEqual, "bar")
			})
		})
	})
}

func TestQuery(t *testing.T) {
	Convey("Given a request with query string", t, func() {
		_, _, ctx := request("POST", "/?foo=bar", "")

		Convey("Query()", func() {
			value := ctx.Query("foo")

			Convey("Should return query value", func() {
				So(value, ShouldEqual, "bar")
			})
		})
	})
}

func TestParse(t *testing.T) {
	Convey("Given a request with json body", t, func() {
		req, _, ctx := request("POST", "/", "{\"foo\":\"bar\"}")
		req.Header.Add("Content-Type", "application/json")

		Convey("Parse()", func() {
			value := map[string]string{}

			err := ctx.Parse(&value)

			Convey("Should return no error", func() {
				So(err, ShouldBeNil)
			})
			Convey("Should decoded body to value", func() {
				So(value["foo"], ShouldEqual, "bar")
			})
		})
	})
	Convey("Given a request with invalid json body", t, func() {
		req, _, ctx := request("POST", "/", "{foo:bar}")
		req.Header.Add("Content-Type", "application/json")

		Convey("Parse()", func() {
			err := ctx.Parse(map[string]string{})

			Convey("Should return error", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
	Convey("Given a request with invalid content type", t, func() {
		_, _, ctx := request("POST", "/", "")

		Convey("Parse()", func() {
			err := ctx.Parse(map[string]string{})

			Convey("Should return error", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestError(t *testing.T) {
	Convey("Given nil and code", t, func() {
		_, res, ctx := request("GET", "/", "")

		Convey("Error()", func() {
			ctx.Error(nil, 404)

			Convey("Should set response status", func() {
				So(res.Code, ShouldEqual, 404)
			})
			Convey("Should set response body", func() {
				So(res.Body.String(), ShouldEqual, "{\"error\":\"Not Found\"}\n")
			})
		})
	})
	Convey("Given error and code", t, func() {
		_, res, ctx := request("GET", "/", "")

		Convey("Error()", func() {
			ctx.Error(errors.New("Invalid"), 400)

			Convey("Should set response status", func() {
				So(res.Code, ShouldEqual, 400)
			})
			Convey("Should set response body", func() {
				So(res.Body.String(), ShouldEqual, "{\"error\":\"Invalid\"}\n")
			})
		})
	})
}

func TestRespond(t *testing.T) {
	Convey("Given nil and code", t, func() {
		_, res, ctx := request("GET", "/", "")

		Convey("Respond()", func() {
			err := ctx.Respond(nil, 204)

			Convey("Should return no error", func() {
				So(err, ShouldBeNil)
			})
			Convey("Should set response status", func() {
				So(res.Code, ShouldEqual, 204)
			})
			Convey("Should set response body", func() {
				So(res.Body.String(), ShouldEqual, "")
			})
		})
	})
	Convey("Given value and code", t, func() {
		_, res, ctx := request("GET", "/", "")

		Convey("Respond()", func() {
			err := ctx.Respond(map[string]string{"foo": "bar"}, 201)

			Convey("Should return no error", func() {
				So(err, ShouldBeNil)
			})
			Convey("Should set response status", func() {
				So(res.Code, ShouldEqual, 201)
			})
			Convey("Should set response body", func() {
				So(res.Body.String(), ShouldEqual, "{\"foo\":\"bar\"}\n")
			})
		})
	})
}

func request(m string, p string, b string) (*http.Request, *httptest.ResponseRecorder, *Context) {
	params := httprouter.Params{httprouter.Param{"foo", "bar"}}

	req, _ := http.NewRequest(m, p, bytes.NewBufferString(b))

	rec := httptest.NewRecorder()

	return req, rec, &Context{req, rec, params}
}
