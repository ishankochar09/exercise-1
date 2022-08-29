package service

import (
	"trainig/exercise1/internal/models"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type product struct {
	ProductSVC ProductStore
}

func NewProductService(store ProductStore) *product {
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

func (p *product) AddVariant(ctx *gofr.Context, variant *models.Variant) (int, error) {
	if variant.ProductID == "" {
		return 0, errors.MissingParam{Param: []string{"product_id"}}
	}

	res, err := p.ProductSVC.AddVariant(ctx, variant)

	if err != nil {
		return 0, err
	}

	return res, nil
}

func (p *product) GetVariant(ctx *gofr.Context, productID, variantID string) (*models.Variant, error) {
	res:= p.ProductSVC.GetVariant(ctx, productID, variantID)
	return res, nil
}
