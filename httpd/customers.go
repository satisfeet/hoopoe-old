package httpd

import (
    "errors"

    "github.com/satisfeet/hoopoe/httpd/router"
    "github.com/satisfeet/hoopoe/store/customers"
)

func CustomersList(c *router.Context) {
	q := &customers.Query{}

	q.Search(c.Query().Get("search"))
	q.Filter(c.Query().Get("filter"))

    r, err := customers.FindAll(q)

    if err != nil {
        c.RespondError(err, 500)
    } else {
        c.RespondJson(r, 200)
    }
}

func CustomersShow(c *router.Context) {
	q := &customers.Query{}

	q.Id(c.Param("customer"))

    r, err := customers.FindOne(q)

    if err != nil {
        c.RespondError(err, 500)
    } else {
        c.RespondJson(r, 200)
    }
}

func CustomersCreate(c *router.Context) {
    c.RespondError(errors.New("Not implemented yet."), 405)
}

func CustomersUpdate(c *router.Context) {
    c.RespondError(errors.New("Not implemented yet."), 405)
}

func CustomersDestroy(c *router.Context) {
    c.RespondError(errors.New("Not implemented yet."), 405)
}
