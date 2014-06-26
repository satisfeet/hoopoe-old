package httpd

import (
    "errors"

    "github.com/satisfeet/hoopoe/lib/store"
    "github.com/satisfeet/hoopoe/lib/httpd/mux"
)

func CustomersList(c *mux.Context) {
    result, err := store.CustomersFind()

    if err != nil {
        c.RespondError(err, 500)
    } else {
        c.RespondJson(result, 200)
    }
}

func CustomersShow(c *mux.Context) {
    result, err := store.CustomersFindOne(c.Params("customer"))

    if err != nil {
        c.RespondError(err, 500)
    } else {
        c.RespondJson(result, 200)
    }
}

func CustomersCreate(c *mux.Context) {
    c.RespondError(errors.New("Not implemented yet."), 406)
}

func CustomersUpdate(c *mux.Context) {
    c.RespondError(errors.New("Not implemented yet."), 406)
}

func CustomersDestroy(c *mux.Context) {
    c.RespondError(errors.New("Not implemented yet."), 406)
}
