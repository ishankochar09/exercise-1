package variant

import (
	"net/http"
	"trainig/exercise1/internal/models"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type VariantRepo struct {
}

func NewVariantRepo() VariantRepo {
	return VariantRepo{}
}

func (v *VariantRepo) AddVariant(ctx *gofr.Context, variant *models.Variant) (int, error) {
	query := "insert into variant(product_id, variantName, variantDetails) values (?,?,?);"
	res, err := ctx.DB().ExecContext(ctx, query, variant.ProductID, variant.VariantName, variant.VariantDetails)
	if err != nil {
		ctx.Logger.Errorf("Error while inserting variant: %v", err)
		return 0, &errors.Response{
			StatusCode: http.StatusInternalServerError,
			Reason:     "oops!! something went wrong...! error encountered in db connection. please try after sometime",
		}
	}
	uId, _ := res.LastInsertId()
	return int(uId), nil
}

func (v *VariantRepo) GetVariant(ctx *gofr.Context, productID, variantID string) *models.Variant {
	var variant models.Variant
	err := ctx.DB().QueryRowContext(ctx, "select id, product_id, variantName, variantDetails from variant where product_id = ? and id = ?", productID, variantID).Scan(&variant.ID, &variant.ProductID, &variant.VariantName, &variant.VariantDetails)
	if err != nil {
		return &models.Variant{}
	}
	return &variant
}
