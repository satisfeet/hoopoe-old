package httpd

import (
	"github.com/satisfeet/hoopoe/httpd/router"
	"github.com/satisfeet/hoopoe/store"
)

func CustomersInit(r *router.Router) {
	r.Get("/customers", CustomersList)
	r.Pos("/customers", CustomersCreate)
	r.Get("/customers/:customer", CustomersShow)
	r.Put("/customers/:customer", CustomersUpdate)
	r.Del("/customers/:customer", CustomersDestroy)
}

func CustomersList(c *router.Context) {
	r, err := store.CustomersFindAll(store.Query{
		"search": c.Query().Get("search"),
	})

	if err != nil {
		c.RespondError(err, 500)
	} else {
		c.RespondJson(r, 200)
	}
}

func CustomersShow(c *router.Context) {
	r, err := store.CustomersFindOne(store.Query{
		"id": c.Param("customer"),
	})

	if err != nil {
		c.RespondError(err, 500)
	} else {
		c.RespondJson(r, 200)
	}
}

func CustomersCreate(c *router.Context) {
	r := store.Customer{}

	if err := c.ParseJson(&r); err != nil {
		c.RespondError(err, 500)

		return
	}

	if err := store.CustomersCreate(&r); err != nil {
		c.RespondError(err, 500)

		return
	}

	c.RespondJson(&r, 200)
}

func CustomersUpdate(c *router.Context) {
	r, err := store.CustomersFindOne(store.Query{
		"id": c.Param("customer"),
	})

	if err != nil {
		c.RespondError(err, 500)

		return
	}

	if err := c.ParseJson(&r); err != nil {
		c.RespondError(err, 500)

		return
	}

	if err := store.CustomersUpdate(&r); err != nil {
		c.RespondError(err, 500)

		return
	}

	c.Respond("", 204)
}

func CustomersDestroy(c *router.Context) {
	r, err := store.CustomersFindOne(store.Query{
		"id": c.Param("customer"),
	})

	if err != nil {
		c.RespondError(err, 500)

		return
	}

	if err := store.CustomersRemove(&r); err != nil {
		c.RespondError(err, 500)

		return
	}

	c.Respond("", 204)
}
