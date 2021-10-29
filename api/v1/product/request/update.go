package request

import "AltaStore/business/product"

type UpdateProductRequest struct {
	Name              string `json:"name"`
	Price             int64  `json:"price"`
	IsActive          bool   `json:"isactive"`
	ProductCategoryId string `json:"productcategoryid"`
	UnitName          string `json:"unitname"`
	Description       string `json:"description"`
}

func (u *UpdateProductRequest) ToProductSpec() *product.UpdateProductSpec {
	var spec product.UpdateProductSpec

	spec.Name = u.Name
	spec.Price = u.Price
	spec.IsActive = u.IsActive
	spec.ProductCategoryId = u.ProductCategoryId
	spec.UnitName = u.UnitName
	spec.Description = u.Description

	return &spec
}
