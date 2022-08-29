package handler

import (
	"strconv"
	"trainig/exercise1/internal/models"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type Product struct {
	service ProductService
}

func NewProductHandler(han ProductService) *Product {
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

func (p *Product) AddVariant(ctx *gofr.Context) (interface{}, error) {
	variant := &models.Variant{}
	// params := ctx.Params()
	// pid := params["pid"]
	// if pid

	param := ctx.PathParam("pid")

	err := ctx.Bind(variant)

	variant.ProductID = param

	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"id"}}
	}

	return p.service.AddVariant(ctx, variant)
}

func (p *Product) GetVariant(ctx *gofr.Context) (interface{}, error) {
	param := ctx.PathParam("pid")

	paramID := ctx.PathParam("vid")

	res, err := p.service.GetVariant(ctx, param, paramID)

	if err != nil {
		return nil, err
	}

	return res, err
}
