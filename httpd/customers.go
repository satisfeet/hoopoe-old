package httpd

import (
	"github.com/satisfeet/hoopoe/httpd/router"
	. "github.com/satisfeet/hoopoe/store/customers"
)

func customersList(c *router.Context) {
	r, err := FindAll(Query{"search": c.Query().Get("search")})

	if err != nil {
		c.RespondError(err, 500)
	} else {
		c.RespondJson(r, 200)
	}
}

func customersShow(c *router.Context) {
	r, err := FindOne(Query{"id": c.Param("customer")})

	if err != nil {
		c.RespondError(err, 500)
	} else {
		c.RespondJson(r, 200)
	}
}

func customersCreate(c *router.Context) {
	r := Customer{}

	err := c.ParseJson(&r)

	if err != nil {
		c.RespondError(err, 500)

		return
	}

	err = Create(&r)

	if err != nil {
		c.RespondError(err, 500)

		return
	}

	c.RespondJson(&r, 200)
}

func customersUpdate(c *router.Context) {
	r, err := FindOne(Query{"id": c.Param("customer")})

	if err != nil {
		c.RespondError(err, 500)

		return
	}

	err = c.ParseJson(&r)

	if err != nil {
		c.RespondError(err, 500)

		return
	}

	err = Update(&r)

	if err != nil {
		c.RespondError(err, 500)

		return
	}

	c.Respond("", 204)
}

func customersDestroy(c *router.Context) {
	r, err := FindOne(Query{"id": c.Param("customer")})

	if err != nil {
		c.RespondError(err, 500)

		return
	}

	err = Remove(&r)

	if err != nil {
		c.RespondError(err, 500)

		return
	}

	c.Respond("", 204)
}
