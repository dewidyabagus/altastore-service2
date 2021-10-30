package category_test

import (
	"AltaStore/business"
	"AltaStore/business/category"
	categoryMock "AltaStore/business/category/mocks"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	adminId = "f9c8c2bf-d525-420e-86e5-4caf03cd8027"

	id   = "f9c8c2bf-d525-420e-86e5-4caf03cd8027"
	code = "code"
	name = "name"
)

var (
	categoryRepository categoryMock.Repository
	categoryService    category.Service

	categoryData  category.Category
	categoryDatas []category.Category
	categorySpec  category.CategorySpec
)

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func setup() {
	categoryData = category.Category{
		ID:   id,
		Code: code,
		Name: name,
	}
	categorySpec = category.CategorySpec{
		code,
		name,
	}
	categoryDatas = append(categoryDatas, categoryData)

	categoryService = category.NewService(&categoryRepository)
}

func TestInsertCategory(t *testing.T) {
	// t.Run("Expect Admin Not Found", func(t *testing.T) {
	// 	adminService.On("FindAdminByID", mock.AnythingOfType("string")).Return(nil, business.ErrNotHavePermission).Once()

	// 	err := categoryService.InsertCategory(&categorySpec, adminId)

	// 	assert.NotNil(t, err)
	// 	assert.Equal(t, err, business.ErrNotHavePermission)

	// })
	t.Run("Expect Category Exist", func(t *testing.T) {
		//adminService.On("FindAdminByID", mock.AnythingOfType("string")).Return(&adminData, nil).Once()
		categoryRepository.On("FindCategoryByCode", mock.AnythingOfType("string")).Return(&categoryData, nil).Once()

		err := categoryService.InsertCategory(&categorySpec, adminId)

		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrDataExists)

	})
	t.Run("Expect Insert Product Category Success", func(t *testing.T) {
		//adminService.On("FindAdminByID", mock.AnythingOfType("string")).Return(&adminData, nil).Once()
		categoryRepository.On("FindCategoryByCode", mock.AnythingOfType("string")).Return(nil, nil).Once()
		categoryRepository.On("InsertCategory", mock.AnythingOfType("category.Category")).Return(nil).Once()

		err := categoryService.InsertCategory(&categorySpec, id)

		assert.Nil(t, err)
	})
	t.Run("Expect Insert Product Category Fail", func(t *testing.T) {
		//adminService.On("FindAdminByID", mock.AnythingOfType("string")).Return(&adminData, nil).Once()
		categoryRepository.On("FindCategoryByCode", mock.AnythingOfType("string")).Return(nil, nil).Once()
		categoryRepository.On("InsertCategory", mock.AnythingOfType("category.Category")).Return(business.ErrInternalServer).Once()

		err := categoryService.InsertCategory(&categorySpec, id)

		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrInternalServer)
	})
}

func TestUpdateCategory(t *testing.T) {
	// t.Run("Expect Admin Not Found", func(t *testing.T) {
	// 	adminService.On("FindAdminByID", mock.AnythingOfType("string")).Return(nil, business.ErrNotHavePermission).Once()

	// 	err := categoryService.UpdateCategory(id, &categorySpec, adminId)

	// 	assert.NotNil(t, err)
	// 	assert.Equal(t, err, business.ErrNotHavePermission)

	// })
	t.Run("Expect Update Product Category Success", func(t *testing.T) {
		//adminService.On("FindAdminByID", mock.AnythingOfType("string")).Return(&adminData, nil).Once()
		categoryRepository.On("UpdateCategory", mock.AnythingOfType("string"), mock.AnythingOfType("category.Category")).Return(nil).Once()

		err := categoryService.UpdateCategory(id, &categorySpec, id)

		assert.Nil(t, err)
	})
	t.Run("Expect Update Product Category Fail", func(t *testing.T) {
		//adminService.On("FindAdminByID", mock.AnythingOfType("string")).Return(&adminData, nil).Once()
		categoryRepository.On("UpdateCategory", mock.AnythingOfType("string"), mock.AnythingOfType("category.Category")).Return(business.ErrInternalServer).Once()

		err := categoryService.UpdateCategory(id, &categorySpec, id)

		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrInternalServer)
	})
}

func TestDeleteCategory(t *testing.T) {
	// t.Run("Expect Admin Not Found", func(t *testing.T) {
	// 	adminService.On("FindAdminByID", mock.AnythingOfType("string")).Return(nil, business.ErrNotHavePermission).Once()

	// 	err := categoryService.DeleteCategory(id, adminId)

	// 	assert.NotNil(t, err)
	// 	assert.Equal(t, err, business.ErrNotHavePermission)

	// })
	t.Run("Expect Delete Product Category Success", func(t *testing.T) {
		//adminService.On("FindAdminByID", mock.AnythingOfType("string")).Return(&adminData, nil).Once()
		categoryRepository.On("DeleteCategory", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil).Once()

		err := categoryService.DeleteCategory(id, adminId)

		assert.Nil(t, err)
	})
	t.Run("Expect Delete Product Category Fail", func(t *testing.T) {
		//adminService.On("FindAdminByID", mock.AnythingOfType("string")).Return(&adminData, nil).Once()
		categoryRepository.On("DeleteCategory", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(business.ErrInternalServer).Once()

		err := categoryService.DeleteCategory(id, adminId)

		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrInternalServer)
	})
}

func TestGetAllCategory(t *testing.T) {
	t.Run("Expect found the data ProductCategory", func(t *testing.T) {
		categoryRepository.On("GetAllCategory").Return(&categoryDatas, nil).Once()

		categories, err := categoryService.GetAllCategory()

		assert.Nil(t, err)
		assert.NotNil(t, categories)

		assert.Equal(t, id, (*categories)[0].ID)
		assert.Equal(t, code, (*categories)[0].Code)
		assert.Equal(t, name, (*categories)[0].Name)
	})

	t.Run("Expect data nil", func(t *testing.T) {
		categoryRepository.On("GetAllCategory").Return(nil, nil).Once()
		categories, err := categoryService.GetAllCategory()

		assert.Nil(t, err)
		assert.Nil(t, categories)
	})

}

func TestFindCategoryById(t *testing.T) {
	t.Run("Expect found the product category", func(t *testing.T) {
		categoryRepository.On("FindCategoryById", mock.AnythingOfType("string")).Return(&categoryData, nil).Once()

		category, err := categoryService.FindCategoryById(id)

		assert.Nil(t, err)
		assert.NotNil(t, category)

		assert.Equal(t, id, category.ID)
		assert.Equal(t, code, category.Code)
		assert.Equal(t, name, category.Name)

	})

	t.Run("Expect product category not found", func(t *testing.T) {
		categoryRepository.On("FindCategoryById", mock.AnythingOfType("string")).Return(nil, business.ErrNotFound).Once()
		category, err := categoryService.FindCategoryById(id)

		assert.NotNil(t, err)
		assert.Nil(t, category)

		assert.Equal(t, err, business.ErrNotFound)
	})
}

func TestFindCategoryByCode(t *testing.T) {
	t.Run("Expect found the product category", func(t *testing.T) {
		categoryRepository.On("FindCategoryByCode", mock.AnythingOfType("string")).Return(&categoryData, nil).Once()

		category, err := categoryService.FindCategoryByCode(code)

		assert.Nil(t, err)
		assert.NotNil(t, category)

		assert.Equal(t, id, category.ID)
		assert.Equal(t, code, category.Code)
		assert.Equal(t, name, category.Name)

	})

	t.Run("Expect product category not found", func(t *testing.T) {
		categoryRepository.On("FindCategoryByCode", mock.AnythingOfType("string")).Return(nil, business.ErrNotFound).Once()
		category, err := categoryService.FindCategoryByCode(code)

		assert.NotNil(t, err)
		assert.Nil(t, category)

		assert.Equal(t, err, business.ErrNotFound)
	})
}
