package handler

import (
	"trainig/exercise1/internal/models"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type ProductService interface {
	AddProduct(ctx *gofr.Context, product *models.Product) (int, error)
	GetProduct(ctx *gofr.Context, product *models.Product, id int)(models.Products, error)
}

type VariantService interface{
	AddVariant(ctx *gofr.Context, variant *models.Variant) (int, error)
	GetVariant(ctx *gofr.Context, productID, variantID string)(*models.Variant, error)
}
