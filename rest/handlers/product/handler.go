package product

import (
	middleware "ecommerce/rest/middlewares"
	"ecommerce/repo"
)

type Handler struct {
	middlewares *middleware.Middlewares
	productRepo repo.ProductRepo
}

func NewHandler(
	middlewares *middleware.Middlewares,
	productRepo repo.ProductRepo,
	) (*Handler) {
	return &Handler{
		middlewares: middlewares,
		productRepo: productRepo,
	}
}