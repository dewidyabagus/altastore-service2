package api

import (
	"AltaStore/api/middleware"
	"AltaStore/api/v1/category"
	"AltaStore/api/v1/product"
	"AltaStore/api/v1/purchasereceiving"
	"net/http"

	echo "github.com/labstack/echo/v4"
)

func RegisterPath(e *echo.Echo,
	category *category.Controller,
	productController *product.Controller,
	purchaseController *purchasereceiving.Controller,
) {
	if category == nil || productController == nil || purchaseController == nil {
		panic("Invalid parameter")
	}

	// Add logger
	e.Use(middleware.MiddlewareLogger)
	e.Use(middleware.JWTMiddleware())

	// Custome response
	e.HTTPErrorHandler = func(e error, c echo.Context) {
		type Response struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		}
		var response Response
		response.Code = http.StatusInternalServerError // defaul 500
		response.Message = "Internal Server Error"

		if he, ok := e.(*echo.HTTPError); ok {
			response.Code = he.Code
			response.Message = http.StatusText(he.Code)
		}

		c.Logger().Error(e)

		_ = c.JSON(response.Code, response)
	}

	// Routing

	cat := e.Group("v1/categories")
	cat.GET("", category.GetAllCategory)
	cat.GET("/:id", category.FindCategoryById)
	cat.POST("", category.InsertCategory)
	cat.PUT("/:id", category.UpdateCategory)
	cat.DELETE("/:id", category.DeleteCategory)

	product := e.Group("v1/products")
	product.GET("", productController.GetAllProduct)
	product.GET("/:id", productController.FindProductById)
	product.POST("", productController.InsertProduct)
	product.PUT("/:id", productController.UpdateProduct)
	product.DELETE("/:id", productController.DeleteProduct)

	purchRec := e.Group("v1/purchases")
	purchRec.POST("", purchaseController.InsertPurchaseReceiving)
	purchRec.GET("", purchaseController.GetAllPurchaseReceiving)
	purchRec.GET("/:id", purchaseController.FindPurchaseReceivingById)
	purchRec.PUT("/:id", purchaseController.UpdatePurchaseReceiving)
	purchRec.DELETE("/:id", purchaseController.DeletePurchaseReceiving)
}
