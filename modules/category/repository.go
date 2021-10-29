package category

import (
	"AltaStore/business/category"
	"time"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

type ProductCategory struct {
	ID        string    `gorm:"id;type:uuid;primaryKey"`
	Code      string    `gorm:"code;type:varchar(50);unique"`
	Name      string    `gorm:"name;type:varchar(100)"`
	CreatedAt time.Time `gorm:"created_at"`
	CreatedBy string    `gorm:"created_by;type:varchar(50)"`
	UpdatedAt time.Time `gorm:"updated_at"`
	UpdatedBy string    `gorm:"updated_by;type:varchar(50)"`
	DeletedAt time.Time `gorm:"deleted_at"`
	DeletedBy string    `gorm:"deleted_by;type:varchar(50)"`
}

func (p *ProductCategory) toCategory() category.Category {
	return category.Category{
		ID:        p.ID,
		Code:      p.Code,
		Name:      p.Name,
		CreatedAt: p.CreatedAt,
		CreatedBy: p.CreatedBy,
		UpdatedAt: p.UpdatedAt,
		UpdatedBy: p.UpdatedBy,
	}
}

func newDataCategory(cat category.Category) *ProductCategory {
	return &ProductCategory{
		ID:        cat.ID,
		Code:      cat.Code,
		Name:      cat.Name,
		CreatedAt: cat.CreatedAt,
		CreatedBy: cat.CreatedBy,
		UpdatedAt: cat.CreatedAt,
		UpdatedBy: cat.CreatedBy,
	}
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) GetAllCategory() (*[]category.Category, error) {
	var categories []ProductCategory

	err := r.DB.Where(" deleted_by = '' ").Find(&categories).Error
	if err != nil {
		return nil, err
	}

	var result []category.Category
	for _, value := range categories {
		result = append(result, value.toCategory())
	}

	return &result, nil
}

func (r *Repository) FindCategoryById(id string) (*category.Category, error) {
	var category ProductCategory

	err := r.DB.First(&category, "id = ? and deleted_by = ''", id).Error
	if err != nil {
		return nil, err
	}

	response := category.toCategory()
	return &response, nil
}

func (r *Repository) FindCategoryByCode(code string) (*category.Category, error) {
	var category ProductCategory

	err := r.DB.First(&category, "code = ? and deleted_by = ''", code).Error
	if err != nil {
		return nil, err
	}

	response := category.toCategory()
	return &response, nil
}

func (r *Repository) InsertCategory(cat category.Category) error {
	dataCategory := newDataCategory(cat)
	if err := r.DB.Create(dataCategory).Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdateCategory(id string, cat category.Category) error {
	var category ProductCategory
	var err error

	err = r.DB.First(&category, "id = ?", id).Error
	if err != nil {
		return err
	}

	err = r.DB.Model(&category).Updates(
		ProductCategory{Name: cat.Name, UpdatedAt: cat.UpdatedAt, UpdatedBy: cat.UpdatedBy},
	).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) DeleteCategory(id string, userid string) error {
	var category ProductCategory
	var err error

	if err = r.DB.First(&category, "id = ? and deleted_by = ''", id).Error; err != nil {
		return err
	}

	return r.DB.Model(&category).Updates(map[string]interface{}{
		"deleted_by": userid,
		"deleted_at": time.Now(),
	}).Error
}
