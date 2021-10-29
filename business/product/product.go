package product

import (
	"time"

	"github.com/google/uuid"
)

// Membuat core data pada domain bisnis
type Product struct {
	ID                  string
	Code                string
	Name                string
	Price               int64
	Qty                 int32
	QtyCart             int32
	IsActive            bool
	ProductCategoryId   string
	ProductCategoryName string
	UnitName            string
	Description         string
	CreatedAt           time.Time
	CreatedBy           string
	UpdatedAt           time.Time
	UpdatedBy           string
	DeletedAt           time.Time
	DeletedBy           string
}

func NewProduct(
	code string,
	name string,
	price int64,
	isActive bool,
	productCategoryId string,
	unitName string,
	description string,
	creator string,
	createdAt time.Time,
) Product {
	return Product{
		ID:                uuid.NewString(),
		Code:              code,
		Name:              name,
		Price:             price,
		IsActive:          isActive,
		ProductCategoryId: productCategoryId,
		UnitName:          unitName,
		Description:       description,
		CreatedAt:         createdAt,
		CreatedBy:         creator,
		UpdatedAt:         createdAt,
		UpdatedBy:         creator,
	}
}

func (oldData *Product) ModifyProduct(
	name string,
	price int64,
	isActive bool,
	productCategoryId string,
	unitName string,
	description string,
	modifier string,
	updatedAt time.Time,
) Product {
	return Product{
		ID:                oldData.ID,
		Code:              oldData.Code,
		Name:              name,
		Price:             price,
		IsActive:          isActive,
		ProductCategoryId: productCategoryId,
		UnitName:          unitName,
		Description:       description,
		CreatedAt:         oldData.CreatedAt,
		CreatedBy:         oldData.CreatedBy,
		UpdatedAt:         updatedAt,
		UpdatedBy:         modifier,
	}
}

func (oldData *Product) DeleteProduct(
	deleteAt time.Time,
	deleter string) Product {

	return Product{
		ID:                oldData.ID,
		Code:              oldData.Code,
		Name:              oldData.Name,
		Price:             oldData.Price,
		IsActive:          oldData.IsActive,
		ProductCategoryId: oldData.ProductCategoryId,
		UnitName:          oldData.UnitName,
		Description:       oldData.Description,
		CreatedAt:         oldData.CreatedAt,
		CreatedBy:         oldData.CreatedBy,
		UpdatedAt:         oldData.UpdatedAt,
		UpdatedBy:         oldData.UpdatedBy,
		DeletedAt:         deleteAt,
		DeletedBy:         deleter,
	}
}
