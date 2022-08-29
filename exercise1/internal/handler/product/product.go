package product

import (
	"strconv"
	"trainig/exercise1/internal/handler"
	"trainig/exercise1/internal/models"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type Product struct {
	service handler.ProductService
}

func NewProductHandler(han handler.ProductService) *Product {
	return &Product{han}
}

func (p *Product) AddProduct(ctx *gofr.Context) (interface{}, error) {
	product := &models.Product{}

	err := ctx.Bind(product)

	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"id"}}
	}

	return p.service.AddProduct(ctx, product)
}

func (p *Product) GetProduct(ctx *gofr.Context) (interface{}, error) {
	product := &models.Product{}
	

	param := ctx.PathParam("pid")

	id, err := strconv.Atoi(param)

	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"id"}}
	}

	res, err := p.service.GetProduct(ctx, product, id)

	return res, err
}