package httpd

import (
	"net/http"
)

func Listen(c map[string]string) {
	http.Handle("/", &Customers{})

	http.ListenAndServe(c["addr"], nil)
}
