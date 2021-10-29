package category

import (
	"time"

	"github.com/google/uuid"
)

// Membuat core data pada domain bisnis
type Category struct {
	ID        string
	Code      string
	Name      string
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt time.Time
	UpdatedBy string
}

// fungsi menambahkan data kategory
func NewProductCategory(code string, name string,
	creator string, createdAt time.Time,
) Category {
	return Category{
		ID:        uuid.NewString(),
		Code:      code,
		Name:      name,
		CreatedAt: createdAt,
		CreatedBy: creator,
		UpdatedAt: createdAt,
		UpdatedBy: creator,
	}
}

func ModifyProductCategory(name string, modifier string, updatedAt time.Time) Category {
	return Category{
		Name:      name,
		UpdatedAt: updatedAt,
		UpdatedBy: modifier,
	}
}
