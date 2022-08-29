package main

import (
	"trainig/exercise1/internal/handler/product"
	"trainig/exercise1/internal/handler/variant"

	productSVC "trainig/exercise1/internal/service/product"
	variantSVC "trainig/exercise1/internal/service/variant"

	productSTR "trainig/exercise1/internal/store/product"
	variantSTR "trainig/exercise1/internal/store/variant"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

func main() {

	app := gofr.New()
	app.Server.ValidateHeaders = false

	productStore := productSTR.NewProductRepo()
	variantStore := variantSTR.NewVariantRepo()

	productService := productSVC.NewProductService(&productStore)
	variantService := variantSVC.NewVariantService(&variantStore)

	productHandler := product.NewProductHandler(productService)
	variantHandler := variant.NewVariantHandler(variantService)

	app.POST("/product", productHandler.AddProduct)
	app.GET("/product/{pid}", productHandler.GetProduct)

	app.POST("/product/{pid}/variant", variantHandler.AddVariant)
	app.GET("/product/{pid}/variant/{vid}", variantHandler.GetVariant)

	app.Start()
}
