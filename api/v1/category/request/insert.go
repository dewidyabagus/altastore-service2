package request

import (
	"AltaStore/business/category"
)

type InsertCategoryRequest struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

func (i *InsertCategoryRequest) ToCategorySpec() *category.CategorySpec {
	var spec category.CategorySpec

	spec.Code = i.Code
	spec.Name = i.Name

	return &spec
}
