package product_test

import (
	"AltaStore/business"
	"AltaStore/business/category"
	categoryMock "AltaStore/business/category/mocks"
	"AltaStore/business/product"
	productMock "AltaStore/business/product/mocks"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	adminId = "f9c8c2bf-d525-420e-86e5-4caf03cd8027"

	id                        = "f9c8c2bf-d525-420e-86e5-4caf03cd8027"
	code                      = "code"
	name                      = "name"
	price               int64 = 100000000
	isactive                  = true
	productcategoryid         = "f9c8c2bf-d525-420e-86e5-4caf03cd8027"
	productcategoryname       = "productcategoryname"
	unitname                  = "unitname"
	description               = "description"
	qty                 int32 = 10
	qtycart             int32 = 0
)

var (
	categoryService   categoryMock.Service
	productRepository productMock.Repository
	productService    product.Service

	categoryData      category.Category
	productData       product.Product
	productDatas      []product.Product
	insertProductSpec product.InsertProductSpec
	updateProductSpec product.UpdateProductSpec
)

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func setup() {
	productData = product.Product{
		ID:                  id,
		Code:                code,
		Name:                name,
		Price:               price,
		Qty:                 qty,
		QtyCart:             qtycart,
		IsActive:            isactive,
		ProductCategoryId:   productcategoryid,
		ProductCategoryName: productcategoryname,
		UnitName:            unitname,
		Description:         description,
	}
	categoryData = category.Category{
		ID:   id,
		Code: code,
		Name: name,
	}
	insertProductSpec = product.InsertProductSpec{
		Code:              code,
		Name:              name,
		Price:             price,
		IsActive:          isactive,
		ProductCategoryId: productcategoryid,
		UnitName:          unitname,
		Description:       description,
	}

	updateProductSpec = product.UpdateProductSpec{
		Name:              name,
		Price:             price,
		IsActive:          isactive,
		ProductCategoryId: productcategoryid,
		UnitName:          unitname,
		Description:       description,
	}
	productDatas = append(productDatas, productData)

	productService = product.NewService(&categoryService, &productRepository)
}

func TestInsertProduct(t *testing.T) {
	// t.Run("Expect Admin Not Found", func(t *testing.T) {
	// 	adminService.On("FindAdminByID", mock.AnythingOfType("string")).Return(nil, business.ErrNotHavePermission).Once()

	// 	err := productService.InsertProduct(&insertProductSpec, adminId)

	// 	assert.NotNil(t, err)
	// 	assert.Equal(t, err, business.ErrNotHavePermission)

	// })
	t.Run("Expect Product Exist", func(t *testing.T) {
		//adminService.On("FindAdminByID", mock.AnythingOfType("string")).Return(&adminData, nil).Once()
		productRepository.On("FindProductByCode", mock.AnythingOfType("string")).Return(&productData, nil).Once()

		err := productService.InsertProduct(&insertProductSpec, adminId)

		assert.NotNil(t, err)

		assert.Equal(t, err, business.ErrDataExists)
	})
	t.Run("Expect category not Exist", func(t *testing.T) {
		//adminService.On("FindAdminByID", mock.AnythingOfType("string")).Return(&adminData, nil).Once()
		productRepository.On("FindProductByCode", mock.AnythingOfType("string")).Return(nil, nil).Once()
		categoryService.On("FindCategoryById", mock.AnythingOfType("string")).Return(nil, business.ErrInvalidData).Once()

		err := productService.InsertProduct(&insertProductSpec, adminId)

		assert.NotNil(t, err)

		assert.Equal(t, err, business.ErrInvalidData)

	})
	t.Run("Expect Insert Product Success", func(t *testing.T) {
		//adminService.On("FindAdminByID", mock.AnythingOfType("string")).Return(&adminData, nil).Once()
		productRepository.On("FindProductByCode", mock.AnythingOfType("string")).Return(nil, nil).Once()
		categoryService.On("FindCategoryById", mock.AnythingOfType("string")).Return(&categoryData, nil).Once()
		productRepository.On("InsertProduct", mock.AnythingOfType("product.Product")).Return(nil).Once()

		err := productService.InsertProduct(&insertProductSpec, adminId)

		assert.Nil(t, err)
	})
	t.Run("Expect Insert Product product Fail", func(t *testing.T) {
		//adminService.On("FindAdminByID", mock.AnythingOfType("string")).Return(&adminData, nil).Once()
		productRepository.On("FindProductByCode", mock.AnythingOfType("string")).Return(nil, nil).Once()
		categoryService.On("FindCategoryById", mock.AnythingOfType("string")).Return(&categoryData, nil).Once()
		productRepository.On("InsertProduct", mock.AnythingOfType("product.Product")).Return(business.ErrInternalServer).Once()

		err := productService.InsertProduct(&insertProductSpec, adminId)

		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrInternalServer)
	})
}

func TestUpdateProduct(t *testing.T) {
	// t.Run("Expect Admin Not Found", func(t *testing.T) {
	// 	adminService.On("FindAdminByID", mock.AnythingOfType("string")).Return(nil, business.ErrNotHavePermission).Once()

	// 	err := productService.UpdateProduct(id, &updateProductSpec, adminId)

	// 	assert.NotNil(t, err)
	// 	assert.Equal(t, err, business.ErrNotHavePermission)

	// })
	t.Run("Expect Product Not Found", func(t *testing.T) {
		//adminService.On("FindAdminByID", mock.AnythingOfType("string")).Return(&adminData, nil).Once()
		productRepository.On("FindProductById", mock.AnythingOfType("string")).Return(nil, business.ErrNotFound).Once()

		err := productService.UpdateProduct(id, &updateProductSpec, adminId)

		assert.NotNil(t, err, business.ErrNotFound)

	})
	t.Run("Expect category not Exist", func(t *testing.T) {
		//adminService.On("FindAdminByID", mock.AnythingOfType("string")).Return(&adminData, nil).Once()
		productRepository.On("FindProductById", mock.AnythingOfType("string")).Return(&productData, nil).Once()
		categoryService.On("FindCategoryById", mock.AnythingOfType("string")).Return(nil, business.ErrInvalidData).Once()

		err := productService.UpdateProduct(id, &updateProductSpec, adminId)

		assert.NotNil(t, err)

		assert.Equal(t, err, business.ErrInvalidData)

	})
	t.Run("Expect Update Product Success", func(t *testing.T) {
		//adminService.On("FindAdminByID", mock.AnythingOfType("string")).Return(&adminData, nil).Once()
		productRepository.On("FindProductById", mock.AnythingOfType("string")).Return(&productData, nil).Once()
		categoryService.On("FindCategoryById", mock.AnythingOfType("string")).Return(&categoryData, nil).Once()
		productRepository.On("UpdateProduct", mock.AnythingOfType("product.Product")).Return(nil).Once()

		err := productService.UpdateProduct(id, &updateProductSpec, adminId)

		assert.Nil(t, err)
	})
	t.Run("Expect Update Product Fail", func(t *testing.T) {
		//adminService.On("FindAdminByID", mock.AnythingOfType("string")).Return(&adminData, nil).Once()
		productRepository.On("FindProductById", mock.AnythingOfType("string")).Return(&productData, nil).Once()
		categoryService.On("FindCategoryById", mock.AnythingOfType("string")).Return(&categoryData, nil).Once()
		productRepository.On("UpdateProduct", mock.AnythingOfType("product.Product")).Return(business.ErrInternalServer).Once()

		err := productService.UpdateProduct(id, &updateProductSpec, adminId)

		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrInternalServer)
	})
}

func TestDeleteProduct(t *testing.T) {
	// t.Run("Expect Admin Not Found", func(t *testing.T) {
	// 	adminService.On("FindAdminByID", mock.AnythingOfType("string")).Return(nil, business.ErrNotHavePermission).Once()

	// 	err := productService.DeleteProduct(id, adminId)

	// 	assert.NotNil(t, err)
	// 	assert.Equal(t, err, business.ErrNotHavePermission)

	// })
	t.Run("Expect Product Not Found", func(t *testing.T) {
		//adminService.On("FindAdminByID", mock.AnythingOfType("string")).Return(&adminData, nil).Once()
		productRepository.On("FindProductById", mock.AnythingOfType("string")).Return(nil, business.ErrNotFound).Once()

		err := productService.DeleteProduct(id, adminId)

		assert.NotNil(t, err, business.ErrNotFound)
	})
	t.Run("Expect Delete Product Success", func(t *testing.T) {
		//adminService.On("FindAdminByID", mock.AnythingOfType("string")).Return(&adminData, nil).Once()
		productRepository.On("FindProductById", mock.AnythingOfType("string")).Return(&productData, nil).Once()
		productRepository.On("DeleteProduct", mock.AnythingOfType("product.Product")).Return(nil).Once()

		err := productService.DeleteProduct(id, adminId)

		assert.Nil(t, err)
	})
	t.Run("Expect Delete Product Fail", func(t *testing.T) {
		//adminService.On("FindAdminByID", mock.AnythingOfType("string")).Return(&adminData, nil).Once()
		productRepository.On("FindProductById", mock.AnythingOfType("string")).Return(&productData, nil).Once()
		productRepository.On("DeleteProduct", mock.AnythingOfType("product.Product")).Return(business.ErrInternalServer).Once()

		err := productService.DeleteProduct(id, adminId)

		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrInternalServer)
	})
}

func TestGetAllProduct(t *testing.T) {
	t.Run("Expect found the data Product", func(t *testing.T) {
		productRepository.On("GetAllProduct").Return(&productDatas, nil).Once()

		products, err := productService.GetAllProduct()

		assert.Nil(t, err)
		assert.NotNil(t, products)

		assert.Equal(t, id, (*products)[0].ID)
		assert.Equal(t, code, (*products)[0].Code)
		assert.Equal(t, name, (*products)[0].Name)
		assert.Equal(t, price, (*products)[0].Price)
		assert.Equal(t, qty, (*products)[0].Qty)
		assert.Equal(t, qtycart, (*products)[0].QtyCart)
		assert.Equal(t, isactive, (*products)[0].IsActive)
		assert.Equal(t, productcategoryid, (*products)[0].ProductCategoryId)
		assert.Equal(t, productcategoryname, (*products)[0].ProductCategoryName)
		assert.Equal(t, unitname, (*products)[0].UnitName)
		assert.Equal(t, description, (*products)[0].Description)
	})

	t.Run("Expect data nil", func(t *testing.T) {
		productRepository.On("GetAllProduct").Return(nil, nil).Once()
		products, err := productService.GetAllProduct()

		assert.Nil(t, err)
		assert.Nil(t, products)
	})
}

func TestGetAllProductByParameter(t *testing.T) {
	t.Run("Expect found the data Product", func(t *testing.T) {
		productRepository.On("GetAllProductByParameter",
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
		).Return(&productDatas, nil).Once()

		products, err := productService.GetAllProductByParameter(
			id,
			strconv.FormatBool(isactive),
			productcategoryname,
			code,
			name,
		)

		assert.Nil(t, err)
		assert.NotNil(t, products)

		assert.Equal(t, id, (*products)[0].ID)
		assert.Equal(t, code, (*products)[0].Code)
		assert.Equal(t, name, (*products)[0].Name)
		assert.Equal(t, price, (*products)[0].Price)
		assert.Equal(t, qty, (*products)[0].Qty)
		assert.Equal(t, qtycart, (*products)[0].QtyCart)
		assert.Equal(t, isactive, (*products)[0].IsActive)
		assert.Equal(t, productcategoryid, (*products)[0].ProductCategoryId)
		assert.Equal(t, productcategoryname, (*products)[0].ProductCategoryName)
		assert.Equal(t, unitname, (*products)[0].UnitName)
		assert.Equal(t, description, (*products)[0].Description)
	})

	t.Run("Expect data nil", func(t *testing.T) {
		productRepository.On("GetAllProductByParameter",
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
		).Return(nil, nil).Once()
		products, err := productService.GetAllProductByParameter(
			id,
			strconv.FormatBool(isactive),
			productcategoryname,
			code,
			name,
		)

		assert.Nil(t, err)
		assert.Nil(t, products)
	})
}

func TestFindProductById(t *testing.T) {
	t.Run("Expect found the product", func(t *testing.T) {
		productRepository.On("FindProductById", mock.AnythingOfType("string")).Return(&productData, nil).Once()

		product, err := productService.FindProductById(id)

		assert.Nil(t, err)
		assert.NotNil(t, product)

		assert.Equal(t, id, product.ID)
		assert.Equal(t, code, product.Code)
		assert.Equal(t, name, product.Name)
		assert.Equal(t, price, product.Price)
		assert.Equal(t, qty, product.Qty)
		assert.Equal(t, qtycart, product.QtyCart)
		assert.Equal(t, isactive, product.IsActive)
		assert.Equal(t, productcategoryid, product.ProductCategoryId)
		assert.Equal(t, productcategoryname, product.ProductCategoryName)
		assert.Equal(t, unitname, product.UnitName)
		assert.Equal(t, description, product.Description)
	})

	t.Run("Expect product product not found", func(t *testing.T) {
		productRepository.On("FindProductById", mock.AnythingOfType("string")).Return(nil, business.ErrNotFound).Once()
		product, err := productService.FindProductById(id)

		assert.NotNil(t, err)
		assert.Nil(t, product)

		assert.Equal(t, err, business.ErrNotFound)
	})
}

func TestFindProductByCode(t *testing.T) {
	t.Run("Expect found the product", func(t *testing.T) {
		productRepository.On("FindProductByCode", mock.AnythingOfType("string")).Return(&productData, nil).Once()

		product, err := productService.FindProductByCode(code)

		assert.Nil(t, err)
		assert.NotNil(t, product)

		assert.Equal(t, id, product.ID)
		assert.Equal(t, code, product.Code)
		assert.Equal(t, name, product.Name)
		assert.Equal(t, price, product.Price)
		assert.Equal(t, qty, product.Qty)
		assert.Equal(t, qtycart, product.QtyCart)
		assert.Equal(t, isactive, product.IsActive)
		assert.Equal(t, productcategoryid, product.ProductCategoryId)
		assert.Equal(t, productcategoryname, product.ProductCategoryName)
		assert.Equal(t, unitname, product.UnitName)
		assert.Equal(t, description, product.Description)
	})

	t.Run("Expect product product not found", func(t *testing.T) {
		productRepository.On("FindProductByCode", mock.AnythingOfType("string")).Return(nil, business.ErrNotFound).Once()
		product, err := productService.FindProductByCode(code)

		assert.NotNil(t, err)
		assert.Nil(t, product)

		assert.Equal(t, err, business.ErrNotFound)
	})
}
