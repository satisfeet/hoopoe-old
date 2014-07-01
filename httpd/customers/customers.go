package customers

import (
	"github.com/satisfeet/hoopoe/httpd/router"
	"github.com/satisfeet/hoopoe/store/customers"
)

func List(context *router.Context) {
	q := customers.Query{
		"search": context.Query().Get("search"),
	}

	r, err := customers.FindAll(q)

	if err != nil {
		context.RespondError(err, 500)
	} else {
		context.RespondJson(r, 200)
	}
}

func Show(context *router.Context) {
	q := customers.Query{
		"id": context.Param("customer"),
	}

	r, err := customers.FindOne(q)

	if err != nil {
		context.RespondError(err, 500)
	} else {
		context.RespondJson(r, 200)
	}
}

func Create(context *router.Context) {
	r := customers.Customer{}

	err := context.ParseJson(&r)

	if err != nil {
		context.RespondError(err, 500)

		return
	}

	err = customers.Create(&r)

	if err != nil {
		context.RespondError(err, 500)

		return
	}

	context.RespondJson(&r, 200)
}

func Update(context *router.Context) {
	q := customers.Query{
		"id": context.Param("customer"),
	}

	r, err := customers.FindOne(q)

	if err != nil {
		context.RespondError(err, 500)

		return
	}

	err = context.ParseJson(&r)

	if err != nil {
		context.RespondError(err, 500)

		return
	}

	err = customers.Update(&r)

	if err != nil {
		context.RespondError(err, 500)

		return
	}

	context.Respond("", 204)
}

func Destroy(context *router.Context) {
	q := customers.Query{
		"id": context.Param("customer"),
	}

	r, err := customers.FindOne(q)

	if err != nil {
		context.RespondError(err, 500)

		return
	}

	err = customers.Remove(&r)

	if err != nil {
		context.RespondError(err, 500)

		return
	}

	context.Respond("", 204)
}
