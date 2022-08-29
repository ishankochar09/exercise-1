package product

import (
	"net/http"
	"trainig/exercise1/internal/models"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type ProductRepo struct {
}

func NewProductRepo() ProductRepo {
	return ProductRepo{}
}

func (p *ProductRepo) AddProduct(ctx *gofr.Context, product *models.Product) (int, error) {
	query := "insert into product(name, brandName,details,imageURL) values (?,?,?,?);"
	res, err := ctx.DB().ExecContext(ctx, query, product.Name, product.BrandName, product.Details, product.ImageURL)
	if err != nil {
		ctx.Logger.Errorf("Error while inserting product: %v", err)
		return 0, &errors.Response{
			StatusCode: http.StatusInternalServerError,
			Reason:     "oops!! something went wrong...! error encountered in db connection. please try after sometime",
		}
	}
	uId, _ := res.LastInsertId()
	return int(uId), nil
}

func (p *ProductRepo) GetProduct(ctx *gofr.Context, product *models.Product, id int) models.Product {
	err := ctx.DB().QueryRowContext(ctx, "select id, name, brandName, details, imageURL from product where id = ?", id).Scan(&product.ID, &product.Name, &product.BrandName, &product.Details, &product.ImageURL)
	if err != nil {
		return models.Product{}
	}
	return *product
}

func (p *ProductRepo) GetProductVariants(ctx *gofr.Context, pid int) []models.Variant {
	var variants []models.Variant
	rows, err := ctx.DB().QueryContext(ctx, "select id, product_id, variantName, variantDetails from variant where product_id = ? ", pid)
	if err != nil {
		return []models.Variant{}
	}
	defer rows.Close()
	for rows.Next() {
		var variant models.Variant
		err := rows.Scan(&variant.ID, &variant.ProductID, &variant.VariantName, &variant.VariantDetails)
		if err != nil {
			return []models.Variant{}
		}
		variants = append(variants, variant)
	}
	return variants
}