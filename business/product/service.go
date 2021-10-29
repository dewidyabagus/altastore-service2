package product

import (
	"AltaStore/business"
	"AltaStore/business/category"

	// "AltaStore/business/admin"
	"AltaStore/util/validator"
	"time"
)

type InsertProductSpec struct {
	Code              string `validate:"required"`
	Name              string `validate:"required"`
	Price             int64  `validate:"required"`
	IsActive          bool
	ProductCategoryId string `validate:"required"`
	UnitName          string `validate:"required"`
	Description       string
}

type UpdateProductSpec struct {
	Name              string `validate:"required"`
	Price             int64  `validate:"required"`
	IsActive          bool   `validate:"required"`
	ProductCategoryId string `validate:"required"`
	UnitName          string `validate:"required"`
	Description       string
}

type service struct {
	// adminService    admin.Service
	categoryService category.Service
	repository      Repository
}

// func NewService(adminService admin.Service, categoryService category.Service, repository Repository) Service {
// 	return &service{adminService, categoryService, repository}
// }

func NewService(categoryService category.Service, repository Repository) Service {
	return &service{categoryService, repository}
}

func (s *service) GetAllProduct() (*[]Product, error) {
	return s.repository.GetAllProduct()
}

func (s *service) GetAllProductByParameter(id, isActive, categoryName, code, name string) (*[]Product, error) {
	return s.repository.GetAllProductByParameter(id, isActive, categoryName, code, name)
}

func (s *service) FindProductById(id string) (*Product, error) {
	return s.repository.FindProductById(id)
}

func (s *service) FindProductByCode(code string) (*Product, error) {
	return s.repository.FindProductByCode(code)
}

func (s *service) InsertProduct(product *InsertProductSpec, creator string) error {
	err := validator.GetValidator().Struct(product)
	if err != nil {
		return business.ErrInvalidSpec
	}
	// admin, err := s.adminService.FindAdminByID(creator)
	// if err != nil {
	// 	return business.ErrNotHavePermission
	// }
	dataproduct, _ := s.repository.FindProductByCode(product.Code)
	if dataproduct != nil {
		return business.ErrDataExists
	}
	category, err := s.categoryService.FindCategoryById(product.ProductCategoryId)
	if err != nil {
		return business.ErrInvalidData
	}
	data := NewProduct(
		product.Code,
		product.Name,
		product.Price,
		product.IsActive,
		category.ID,
		product.UnitName,
		product.Description,
		//admin.ID,
		creator,
		time.Now(),
	)
	return s.repository.InsertProduct(data)
}

func (s *service) UpdateProduct(id string, updateProduct *UpdateProductSpec, modifier string) error {
	err := validator.GetValidator().Struct(updateProduct)
	if err != nil {
		return business.ErrInvalidSpec
	}
	// admin, err := s.adminService.FindAdminByID(modifier)
	// if err != nil {
	// 	return business.ErrNotHavePermission
	// }
	product, err := s.repository.FindProductById(id)
	if err != nil {
		return err
	} else if product == nil {
		return business.ErrNotFound
	} else if product.DeletedBy != "" {
		return business.ErrNotFound
	}
	category, err := s.categoryService.FindCategoryById(updateProduct.ProductCategoryId)
	if err != nil {
		return business.ErrInvalidData
	}
	modifiedproduct := product.ModifyProduct(
		updateProduct.Name,
		updateProduct.Price,
		updateProduct.IsActive,
		category.ID,
		updateProduct.UnitName,
		updateProduct.Description,
		//admin.ID,
		modifier,
		time.Now())

	return s.repository.UpdateProduct(modifiedproduct)
}

func (s *service) DeleteProduct(id string, deleter string) error {
	// admin, err := s.adminService.FindAdminByID(deleter)
	// if err != nil {
	// 	return business.ErrNotHavePermission
	// }
	product, err := s.repository.FindProductById(id)
	if err != nil {
		return err
	} else if product == nil {
		return business.ErrNotFound
	} else if product.DeletedBy != "" {
		return business.ErrNotFound
	}
	deleteProduct := product.DeleteProduct(
		time.Now(),
		//admin.ID,
		deleter,
	)
	return s.repository.DeleteProduct(deleteProduct)
}
