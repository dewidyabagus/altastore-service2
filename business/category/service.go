package category

import (
	"AltaStore/business"
	"AltaStore/util/validator"
	"time"
)

type CategorySpec struct {
	// AdminId string
	Code string `validate:"required,max=50"`
	Name string `validate:"required,max=100"`
}

type service struct {
	// adminService admin.Service
	repository Repository
}

// func NewService(adminService admin.Service, repository Repository) Service {
func NewService(repository Repository) Service {
	// return &service{adminService, repository}
	return &service{repository}
}

func (s *service) GetAllCategory() (*[]Category, error) {
	return s.repository.GetAllCategory()
}

func (s *service) FindCategoryById(id string) (*Category, error) {
	return s.repository.FindCategoryById(id)
}

func (s *service) FindCategoryByCode(code string) (*Category, error) {
	return s.repository.FindCategoryByCode(code)
}

func (s *service) InsertCategory(category *CategorySpec, creator string) error {
	// admin, err := s.adminService.FindAdminByID(creator)
	// if err != nil {
	// 	return business.ErrNotHavePermission
	// }

	err := validator.GetValidator().Struct(category)
	if err != nil {
		return business.ErrInvalidSpec
	}

	data, _ := s.repository.FindCategoryByCode(category.Code)
	if data != nil {
		return business.ErrDataExists
	}

	dataCategory := NewProductCategory(
		category.Code, category.Name, creator, time.Now(),
	)
	return s.repository.InsertCategory(dataCategory)
}

func (s *service) UpdateCategory(id string, category *CategorySpec, modifier string) error {
	// admin, err := s.adminService.FindAdminByID(modifier)
	// if err != nil {
	// 	return business.ErrNotHavePermission
	// }

	err := validator.GetValidator().Struct(category)
	if err != nil {
		return business.ErrInvalidSpec
	}

	dataCategory := ModifyProductCategory(category.Name, modifier, time.Now())

	return s.repository.UpdateCategory(id, dataCategory)
}

func (s *service) DeleteCategory(id string, deleter string) error {
	// admin, err := s.adminService.FindAdminByID(deleter)
	// if err != nil {
	// 	return business.ErrNotHavePermission
	// }
	return s.repository.DeleteCategory(id, deleter)
}
