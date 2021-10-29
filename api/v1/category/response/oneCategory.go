package response

import (
	"AltaStore/business/category"
	"time"
)

type OneCategory struct {
	ID        string    `json:"id"`
	Code      string    `json:"code"`
	Name      string    `json:"name"`
	UpdatedAt time.Time `json:"updated_at"`
}

func GetOneCategory(category category.Category) *OneCategory {
	var res OneCategory

	res.ID = category.ID
	res.Code = category.Code
	res.Name = category.Name
	res.UpdatedAt = category.UpdatedAt

	return &res
}
