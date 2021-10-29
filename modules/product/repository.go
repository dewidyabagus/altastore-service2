package product

import (
	"AltaStore/business/product"
	category "AltaStore/modules/category"
	"fmt"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

type Product struct {
	ID                string `gorm:"id;type:uuid;primaryKey"`
	Code              string `gorm:"code;type:varchar(50);unique"`
	Name              string `gorm:"name;type:varchar(100)"`
	Price             int64  `gorm:"price"`
	ProductCategoryId string `gorm:"productcategoryid;type:uuid"`
	ProductCategory   category.ProductCategory
	IsActive          bool      `gorm:"isactive;type:boolean"`
	UnitName          string    `gorm:"unitname;type:varchar(50)"`
	Description       string    `gorm:"description;type:varchar(255)"`
	CreatedAt         time.Time `gorm:"created_at"`
	CreatedBy         string    `gorm:"created_by;type:varchar(50)"`
	UpdatedAt         time.Time `gorm:"updated_at"`
	UpdatedBy         string    `gorm:"updated_by;type:varchar(50)"`
	DeletedAt         time.Time `gorm:"deleted_at"`
	DeletedBy         string    `gorm:"deleted_by;type:varchar(50)"`
}

type ProductQuery struct {
	ID                  string    `gorm:"id"`
	Code                string    `gorm:"code"`
	Name                string    `gorm:"name"`
	Price               int64     `gorm:"price"`
	ProductCategoryName string    `gorm:"product_category_name"`
	IsActive            bool      `gorm:"isactive"`
	UnitName            string    `gorm:"unitname"`
	Description         string    `gorm:"description"`
	Qty                 int       `gorm:"qty"`
	QtyCart             int       `gorm:"qty_cart"`
	CreatedAt           time.Time `gorm:"created_at"`
	CreatedBy           string    `gorm:"created_by"`
	UpdatedAt           time.Time `gorm:"updated_at"`
	UpdatedBy           string    `gorm:"updated_by"`
}

func (p *ProductQuery) ToProduct() product.Product {
	return product.Product{
		ID:                  p.ID,
		Code:                p.Code,
		Name:                p.Name,
		Price:               p.Price,
		IsActive:            p.IsActive,
		UnitName:            p.UnitName,
		Description:         p.Description,
		ProductCategoryName: p.ProductCategoryName,
		Qty:                 int32(p.Qty),
		QtyCart:             int32(p.QtyCart),
		CreatedAt:           p.CreatedAt,
		CreatedBy:           p.CreatedBy,
		UpdatedAt:           p.UpdatedAt,
		UpdatedBy:           p.UpdatedBy,
	}
}

func (p *Product) ToProduct() product.Product {
	return product.Product{
		ID:                p.ID,
		Code:              p.Code,
		Name:              p.Name,
		Price:             p.Price,
		IsActive:          p.IsActive,
		ProductCategoryId: p.ProductCategoryId,
		UnitName:          p.UnitName,
		Description:       p.Description,
		CreatedAt:         p.CreatedAt,
		CreatedBy:         p.CreatedBy,
		UpdatedAt:         p.UpdatedAt,
		UpdatedBy:         p.UpdatedBy,
		DeletedAt:         p.DeletedAt,
		DeletedBy:         p.DeletedBy,
	}
}

func newDataProduct(p product.Product) *Product {
	return &Product{
		ID:                p.ID,
		Code:              p.Code,
		Name:              p.Name,
		Price:             p.Price,
		IsActive:          p.IsActive,
		ProductCategoryId: p.ProductCategoryId,
		UnitName:          p.UnitName,
		Description:       p.Description,
		CreatedAt:         p.CreatedAt,
		CreatedBy:         p.CreatedBy,
		UpdatedAt:         p.CreatedAt,
		UpdatedBy:         p.CreatedBy,
		DeletedAt:         p.DeletedAt,
		DeletedBy:         p.DeletedBy,
	}
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) GetAllProduct() (*[]product.Product, error) {
	var products []ProductQuery
	// var tempProduct product.Product

	// err := r.DB.Preload("ProductCategory").Where("deleted_by = ''").Find(&products).Error
	var query = "select t1.*, stock_akhir qty, stock_cart qty_cart, t2.product_category_name from products t1" +
		" inner join product_stock(t1.id) on productid = t1.id " +
		" inner join product_categories t2 on t2.id = product_category_id " +
		" where deleted_by = ''"
	err := r.DB.Raw(query).Scan(&products).Error
	if err != nil {
		return nil, err
	}

	var productOuts []product.Product
	for _, value := range products {
		// tempProduct = value.ToProduct()
		// tempProduct.ProductCategoryName = value.ProductCategory.Name
		// tempProduct.Qty = 10
		// tempProduct.QtyCart = 0
		productOuts = append(productOuts, value.ToProduct()) //tempProduct)
	}

	return &productOuts, nil
}

func (r *Repository) GetAllProductByParameter(id, isActive, categoryName,
	code, name string) (*[]product.Product, error) {
	// var products []Product
	var products []ProductQuery
	// var tempProduct product.Product
	var productOuts []product.Product

	var query = "select t1.*, stock_akhir qty, stock_cart qty_cart, t2.name product_category_name from products t1" +
		" inner join product_stock(t1.id) on productid = t1.id " +
		" inner join product_categories t2 on t2.id = product_category_id " +
		" where t1.deleted_by = ''"

	// temp := r.DB
	// temp = temp.Preload("ProductCategory")
	if categoryName != "" {
		// temp = temp.Joins("inner join product_categories on product_categories.id = products.product_category_id")
		// temp = temp.Where("product_categories.name = ?", categoryName)
		query += fmt.Sprintf(" and t2.name = '%s'", categoryName)
	}

	if id != "" {
		// temp = temp.Where("id = ?", id)
		query += fmt.Sprintf(" and t1.id = '%s'", id)
	}
	if isActive != "" {
		res, err := strconv.ParseBool(isActive)
		if err != nil {
			return &productOuts, nil
		}
		// temp = temp.Where("is_active = ?", res)
		query += fmt.Sprintf(" and t1.is_active = %t", res)
	}
	if code != "" {
		// temp = temp.Where("code = ?", code)
		query += fmt.Sprintf(" and t1.code = '%s'", code)
	}
	if name != "" {
		// temp = temp.Where("name = ?", name)
		query += fmt.Sprintf(" and t1.name = '%s'", name)
	}

	// err := temp.Where("deleted_by = ''").Find(&products).Error
	err := r.DB.Raw(query).Scan(&products).Error
	if err != nil {
		return nil, err
	}

	// for _, value := range products {
	// 	tempProduct = value.ToProduct()

	// 	tempProduct.ProductCategoryName = value.ProductCategory.Name
	// 	tempProduct.Qty = 10
	// 	tempProduct.QtyCart = 0
	// 	productOuts = append(productOuts, tempProduct)
	// }

	for _, value := range products {
		productOuts = append(productOuts, value.ToProduct())
	}

	return &productOuts, nil
}

func (r *Repository) FindProductById(id string) (*product.Product, error) {
	var product *Product

	err := r.DB.Where("id = ?", id).Where("deleted_by = ''").First(&product).Error
	if err != nil {
		return nil, err
	}

	response := product.ToProduct()
	return &response, nil
}

func (r *Repository) FindProductByCode(code string) (*product.Product, error) {
	var product *Product

	err := r.DB.Where("code = ?", code).Where("deleted_by = ''").First(&product).Error
	if err != nil {
		return nil, err
	}

	response := product.ToProduct()
	return &response, nil
}

func (r *Repository) InsertProduct(p product.Product) error {
	product := newDataProduct(p)
	if err := r.DB.Create(product).Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdateProduct(p product.Product) error {
	product := newDataProduct(p)
	err := r.DB.Model(&product).Updates(Product{
		Name:              product.Name,
		Price:             product.Price,
		IsActive:          product.IsActive,
		ProductCategoryId: product.ProductCategoryId,
		UnitName:          product.UnitName,
		Description:       product.Description,
		UpdatedAt:         product.UpdatedAt,
		UpdatedBy:         product.UpdatedBy,
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) DeleteProduct(p product.Product) error {
	product := newDataProduct(p)

	err := r.DB.Model(&product).Updates(Product{
		DeletedBy: product.DeletedBy,
		DeletedAt: product.DeletedAt,
	}).Error
	if err != nil {
		return err
	}
	return nil
}
