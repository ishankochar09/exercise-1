package product

import (
	"trainig/exercise1/internal/models"
	"trainig/exercise1/internal/service"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type product struct {
	ProductSVC service.ProductStore
}

func NewProductService(store service.ProductStore) *product {
	return &product{store}
}

func (p *product) AddProduct(ctx *gofr.Context, product *models.Product) (int, error) {
	if product.Name == "" {
		return 0, errors.MissingParam{Param: []string{"name"}}
	}

	res, err := p.ProductSVC.AddProduct(ctx, product)

	if err != nil {
		return 0, err
	}

	return res, nil
}

func (p *product) GetProduct(ctx *gofr.Context, product *models.Product, id int) (models.Products, error) {
	res := p.ProductSVC.GetProduct(ctx, product, id)

	resp := p.ProductSVC.GetProductVariants(ctx, id)

	response := models.Products{}
	response.Product = res
	response.Variant = resp
	return response, nil
}