package httpd

import (
    "errors"

    "github.com/satisfeet/hoopoe/store"
    "github.com/satisfeet/hoopoe/httpd/router"
)

func CustomersList(c *router.Context) {
    result, err := store.CustomersFind(&store.Query{
        "search": c.Query().Get("search"),
        "filter": c.Query().Get("filter"),
    })

    if err != nil {
        c.RespondError(err, 500)
    } else {
        c.RespondJson(result, 200)
    }
}

func CustomersShow(c *router.Context) {
    result, err := store.CustomersFindOne(&store.Query{
        "id": c.Param("customer"),
    })

    if err != nil {
        c.RespondError(err, 500)
    } else {
        c.RespondJson(result, 200)
    }
}

func CustomersCreate(c *router.Context) {
    c.RespondError(errors.New("Not implemented yet."), 406)
}

func CustomersUpdate(c *router.Context) {
    c.RespondError(errors.New("Not implemented yet."), 406)
}

func CustomersDestroy(c *router.Context) {
    c.RespondError(errors.New("Not implemented yet."), 406)
}
