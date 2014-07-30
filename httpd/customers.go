package httpd

import (
	"net/http"

	"github.com/satisfeet/go-context"
	"github.com/satisfeet/go-router"
	"github.com/satisfeet/hoopoe/model"
	"github.com/satisfeet/hoopoe/store/mongo"
)

type CustomerHandler struct {
	store  *mongo.Store
	router *router.Router
}

func NewCustomerHandler(s *mongo.Store) *CustomerHandler {
	r := router.NewRouter()

	h := &CustomerHandler{
		store:  s,
		router: r,
	}

	r.HandleFunc(router.Read, "/customers", h.list)
	r.HandleFunc(router.Read, "/customers/:id", h.show)
	r.HandleFunc(router.Create, "/customers", h.create)
	r.HandleFunc(router.Update, "/customers/:id", h.update)
	r.HandleFunc(router.Destroy, "/customers/:id", h.destroy)

	return h
}

func (h *CustomerHandler) list(c *context.Context) {
	m := []model.Customer{}
	q := mongo.Query{}

	if s := c.Query("search"); len(s) != 0 {
		// append search conditions
	}

	if err := h.store.Find("customers", q, &m); err != nil {
		c.Error(err, ErrorCode(err))
	} else {
		c.Respond(m, http.StatusOK)
	}
}

func (h *CustomerHandler) show(c *context.Context) {
	m := model.Customer{}
	q := mongo.Query{}

	if err := q.Id(c.Param("id")); err != nil {
		c.Error(err, ErrorCode(err))

		return
	}

	if err := h.store.FindOne("customers", q, &m); err != nil {
		c.Error(err, ErrorCode(err))
	} else {
		c.Respond(m, http.StatusOK)
	}
}

func (h *CustomerHandler) create(c *context.Context) {
	m := model.Customer{}

	if err := c.Parse(&m); err != nil {
		c.Error(err, http.StatusBadRequest)

		return
	}

	if err := h.store.Insert("customers", &m); err != nil {
		c.Error(err, ErrorCode(err))
	} else {
		c.Respond(m, http.StatusOK)
	}
}

func (h *CustomerHandler) update(c *context.Context) {
	m := model.Customer{}
	q := mongo.Query{}

	if err := q.Id(c.Param("id")); err != nil {
		c.Error(err, ErrorCode(err))

		return
	}

	if err := c.Parse(&m); err != nil {
		c.Error(err, http.StatusBadRequest)

		return
	}

	if err := h.store.Update("customers", q, &m); err != nil {
		c.Error(err, ErrorCode(err))
	} else {
		c.Respond(nil, http.StatusNoContent)
	}
}

func (h *CustomerHandler) destroy(c *context.Context) {
	q := mongo.Query{}

	if err := q.Id(c.Param("id")); err != nil {
		c.Error(err, http.StatusBadRequest)

		return
	}
	if err := h.store.Remove("customers", q); err != nil {
		c.Error(err, ErrorCode(err))
	} else {
		c.Respond(nil, http.StatusNoContent)
	}
}

func (h *CustomerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}
