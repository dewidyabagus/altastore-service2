package request

import (
	"AltaStore/business/product"
)

type InsertProductRequest struct {
	Code              string `json:"code"`
	Name              string `json:"name"`
	Price             int64  `json:"price"`
	IsActive          bool   `json:"isactive"`
	ProductCategoryId string `json:"productcategoryid"`
	UnitName          string `json:"unitname"`
	Description       string `json:"description"`
}

func (i *InsertProductRequest) ToProductSpec() *product.InsertProductSpec {
	var spec product.InsertProductSpec

	spec.Code = i.Code
	spec.Name = i.Name
	spec.Price = i.Price
	spec.IsActive = i.IsActive
	spec.ProductCategoryId = i.ProductCategoryId
	spec.UnitName = i.UnitName
	spec.Description = i.Description

	return &spec
}
