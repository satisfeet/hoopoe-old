package customers

import (
	"github.com/satisfeet/hoopoe/httpd/router"
	"github.com/satisfeet/hoopoe/store/customers"
)

func List(c *router.Context) {
	q := &customers.Query{}
	q.Search(c.Query().Get("search"))

	r, err := customers.FindAll(q)

	if err != nil {
		c.RespondError(err, 500)
	} else {
		c.RespondJson(r, 200)
	}
}

func Show(c *router.Context) {
	q := &customers.Query{}
	q.Id(c.Param("customer"))

	r, err := customers.FindOne(q)

	if err != nil {
		c.RespondError(err, 500)
	} else {
		c.RespondJson(r, 200)
	}
}

func Create(c *router.Context) {
	c.RespondError(nil, 405)
}

func Update(c *router.Context) {
	c.RespondError(nil, 405)
}

func Destroy(c *router.Context) {
	c.RespondError(nil, 405)
}
