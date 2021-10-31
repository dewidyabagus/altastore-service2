package response

import (
	"AltaStore/business/product"
)

type ProductById struct {
	ID                  string `json:"id"`
	Code                string `json:"code"`
	Name                string `json:"name"`
	Price               int64  `json:"price"`
	Qty                 int32  `json:"qty"`
	QtyCart             int32  `json:"qtycart"`
	IsActive            bool   `json:"isactive"`
	ProductCategoryName string `json:"ProductCategoryname"`
	UnitName            string `json:"unitname"`
	Description         string `json:"description"`
}

func GetById(product product.Product) *ProductById {
	var prod ProductById

	prod.ID = product.ID
	prod.Code = product.Code
	prod.Name = product.Name
	prod.Price = product.Price
	prod.Qty = product.Qty
	prod.QtyCart = product.QtyCart
	prod.IsActive = product.IsActive
	prod.ProductCategoryName = product.ProductCategoryName
	prod.UnitName = product.UnitName
	prod.Description = product.Description

	return &prod
}
