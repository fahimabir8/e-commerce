package product

import middleware "ecommerce/rest/middlewares"

type Handler struct {
	middlewares *middleware.Middlewares
}

func NewHandler(middlewares *middleware.Middlewares) (h *Handler) {
	return &Handler{
		middlewares: middlewares,
	}
}