package migration

import (
	"gorm.io/gorm"

	"AltaStore/modules/category"
	"AltaStore/modules/product"
	"AltaStore/modules/purchasereceiving"
	"AltaStore/modules/purchasereceivingdetail"
)

func TableMigration(db *gorm.DB) {
	db.AutoMigrate(&category.ProductCategory{},
		&product.Product{},
		&purchasereceiving.PurchaseReceiving{},
		&purchasereceivingdetail.PurchaseReceivingDetail{},
	)
}
