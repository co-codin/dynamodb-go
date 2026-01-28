package product

import (
	"dynamo-golang/internal/controllers/product"
	"dynamo-golang/internal/handlers"
)

type Handler struct {
	handlers.Interface
	Controller product.Interface
}