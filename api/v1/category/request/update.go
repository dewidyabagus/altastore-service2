package request

import "AltaStore/business/category"

type UpdateCategoryRequest struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

func (u *UpdateCategoryRequest) ToCategory() *category.CategorySpec {
	var spec category.CategorySpec

	spec.Code = u.Code
	spec.Name = u.Name

	return &spec
}
