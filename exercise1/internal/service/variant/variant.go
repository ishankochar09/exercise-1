package variant

import (
	"trainig/exercise1/internal/models"
	"trainig/exercise1/internal/service"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type variant struct {
	VariantSVC service.VariantStore
}

func NewVariantService(store service.VariantStore) *variant {
	return &variant{store}
}


func (v *variant) AddVariant(ctx *gofr.Context, variant *models.Variant) (int, error) {
	if variant.ProductID == "" {
		return 0, errors.MissingParam{Param: []string{"product_id"}}
	}

	res, err := v.VariantSVC.AddVariant(ctx, variant)

	if err != nil {
		return 0, err
	}

	return res, nil
}

func (v *variant) GetVariant(ctx *gofr.Context, productID, variantID string) (*models.Variant, error) {
	res:= v.VariantSVC.GetVariant(ctx, productID, variantID)
	return res, nil
}
