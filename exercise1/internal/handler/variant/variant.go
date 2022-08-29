package variant

import (
	"trainig/exercise1/internal/models"
	"trainig/exercise1/internal/handler"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type Variant struct {
	Vservice handler.VariantService
}

func NewVariantHandler(han handler.VariantService) *Variant {
	return &Variant{han}
}

func (v *Variant) AddVariant(ctx *gofr.Context) (interface{}, error) {
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

	return v.Vservice.AddVariant(ctx, variant)
}

func (v * Variant) GetVariant(ctx *gofr.Context) (interface{}, error) {
	param := ctx.PathParam("pid")

	paramID := ctx.PathParam("vid")

	res, err := v.Vservice.GetVariant(ctx, param, paramID)

	if err != nil {
		return nil, err
	}

	return res, err
}
