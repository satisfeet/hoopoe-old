package httpd

import (
	"net/http"

	"github.com/satisfeet/hoopoe/httpd/router"
)

func Listen(c map[string]string) {
	r := router.New()

	r.Use(Auth)
	r.Use(Logger)

	CustomersInit(r)

	http.ListenAndServe(c["addr"], r)
}
