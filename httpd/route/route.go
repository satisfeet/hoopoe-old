package route

import "net/http"

// Defines a REST operation to route.
type Action int

const (
	Invalid Action = iota
	List
	Show
	Create
	Update
	Destroy
)

// Takes an http request and returns the correct REST action.
func Route(s string, r *http.Request) (Action, string) {
	p := r.URL.Path

	if len(p) == 1 {
		return Invalid, ""
	}

	b := []byte{}
	for i := 0; i < len(p); i++ {
		if i < len(s) {
			if p[i] != s[i] {
				return Invalid, ""
			}
		}
		if i == len(s) {
			if p[i] != '/' {
				return Invalid, ""
			}
		}
		if i > len(s) {
			if p[i] == '/' {
				return Invalid, ""
			} else {
				b = append(b, p[i])
			}
		}
	}
	id := string(b)

	if id == "" {
		switch r.Method {
		case "GET":
			return List, id
		case "POST":
			return Create, id
		}
	} else {
		switch r.Method {
		case "GET":
			return Show, id
		case "PUT":
			return Update, id
		case "DELETE":
			return Destroy, id
		}
	}

	return Invalid, id
}
