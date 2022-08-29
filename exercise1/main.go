package main

import (
	"trainig/exercise1/internal/handler"
	"trainig/exercise1/internal/service"
	"trainig/exercise1/internal/store"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

func main() {

	app := gofr.New()
	app.Server.ValidateHeaders = false

	productStore := store.NewProductRepo()
	productService := service.NewProductService(&productStore)
	productHandler := handler.NewProductHandler(productService)

	app.POST("/product", productHandler.AddProduct)
	app.GET("/product/{pid}", productHandler.GetProduct)
	app.POST("/product/{pid}/variant", productHandler.AddVariant)
	app.GET("/product/{pid}/variant/{vid}", productHandler.GetVariant)

	app.Start()
}
