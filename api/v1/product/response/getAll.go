package response

import "AltaStore/business/product"

type Product struct {
	ID                  string `json:"id"`
	Code                string `json:"code"`
	Name                string `json:"name"`
	Price               int64  `json:"price"`
	Qty                 int32  `json:"qty"`
	QtyCart             int32  `json:"qtycart"`
	IsActive            bool   `json:"isactive"`
	ProductCategoryName string `json:"productcategoryname"`
	UnitName            string `json:"unitname"`
	Description         string `json:"description"`
}

type Products struct {
	Products []Product
}

func GetAll(products *[]product.Product) *Products {
	var allProduct Products
	var prod Product
	for _, val := range *products {

		prod.ID = val.ID
		prod.Code = val.Code
		prod.Name = val.Name
		prod.Price = val.Price
		prod.Qty = val.Qty
		prod.QtyCart = val.QtyCart
		prod.IsActive = val.IsActive
		prod.ProductCategoryName = val.ProductCategoryName
		prod.UnitName = val.UnitName
		prod.Description = val.Description

		allProduct.Products = append(allProduct.Products, prod)
	}

	if allProduct.Products == nil {
		allProduct.Products = []Product{}
	}

	return &allProduct
}
