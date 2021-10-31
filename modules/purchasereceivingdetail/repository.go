package purchasereceivingdetail

import (
	"AltaStore/business"
	"AltaStore/business/purchasereceiving"
	"errors"
	"time"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

type PurchaseReceivingDetail struct {
	ID                  string    `gorm:"id;type:uuid;primaryKey"`
	PurchaseReceivingId string    `gorm:"purchase_receiving_id;type:uuid"`
	ProductId           string    `gorm:"product_id;type:uuid"`
	Price               int64     `gorm:"price;"`
	Qty                 int32     `gorm:"qty;"`
	CreatedAt           time.Time `gorm:"created_at"`
	CreatedBy           string    `gorm:"created_by;type:uuid"`
	UpdatedAt           time.Time `gorm:"updated_at"`
	UpdatedBy           string    `gorm:"updated_by;type:varchar(50)"`
	DeletedAt           time.Time `gorm:"deleted_at"`
	DeletedBy           string    `gorm:"deleted_by;type:varchar(50)"`
}

type PurchaseReceivingDetailQuery struct {
	ID                  string    `gorm:"id"`
	PurchaseReceivingId string    `gorm:"purchase_receiving_id"`
	ProductId           string    `gorm:"product_id"`
	ProductCode         string    `gorm:"product_code"`
	ProductName         string    `gorm:"product_name"`
	Price               int64     `gorm:"price;"`
	Qty                 int32     `gorm:"qty;"`
	CreatedAt           time.Time `gorm:"created_at"`
	CreatedBy           string    `gorm:"created_by"`
	UpdatedAt           time.Time `gorm:"updated_at"`
	UpdatedBy           string    `gorm:"updated_by"`
}

// func (p *PurchaseReceivingDetail) toPurchaseReceivingDetail() *purchasereceiving.PurchaseReceivingDetail {
// 	return &purchasereceiving.PurchaseReceivingDetail{
// 		ID:        p.ID,
// 		ProductId: p.ProductId,
// 		Price:     p.Price,
// 		Qty:       p.Qty,
// 		CreatedBy: p.CreatedBy,
// 		CreatedAt: p.CreatedAt,
// 		UpdatedBy: p.UpdatedBy,
// 		UpdatedAt: p.UpdatedAt,
// 		DeletedBy: p.DeletedBy,
// 		DeletedAt: p.DeletedAt,
// 	}
// }

func (p *PurchaseReceivingDetailQuery) toPurchaseReceivingDetail() *purchasereceiving.PurchaseReceivingDetail {
	return &purchasereceiving.PurchaseReceivingDetail{
		ID:                  p.ID,
		PurchaseReceivingId: p.PurchaseReceivingId,
		ProductId:           p.ProductId,
		ProductCode:         p.ProductCode,
		ProductName:         p.ProductName,
		Qty:                 p.Qty,
		Price:               p.Price,
		CreatedBy:           p.CreatedBy,
		CreatedAt:           p.CreatedAt,
		UpdatedBy:           p.UpdatedBy,
		UpdatedAt:           p.UpdatedAt,
	}
}

func newDataPurchaseReceivingDetail(
	p *purchasereceiving.PurchaseReceivingDetail,
) *PurchaseReceivingDetail {
	return &PurchaseReceivingDetail{
		ID:                  p.ID,
		PurchaseReceivingId: p.PurchaseReceivingId,
		ProductId:           p.ProductId,
		Price:               p.Price,
		Qty:                 p.Qty,
		CreatedBy:           p.CreatedBy,
		CreatedAt:           p.CreatedAt,
		UpdatedBy:           p.UpdatedBy,
		UpdatedAt:           p.UpdatedAt,
		DeletedBy:           p.DeletedBy,
		DeletedAt:           p.DeletedAt,
	}
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) InsertPurchaseReceivingDetail(p *purchasereceiving.PurchaseReceivingDetail) error {
	detail := newDataPurchaseReceivingDetail(p)
	if err := r.DB.Create(detail).Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdatePurchaseReceivingDetail(purchaseId, productId *string, price, qty *int, modifier *string, updater time.Time) error {
	var purchase = new(PurchaseReceivingDetail)

	if err := r.DB.First(purchase, "purchase_receiving_id = ? and product_id = ?", purchaseId, productId).Error; err != nil {
		return err
	}

	return r.DB.Model(purchase).Updates(map[string]interface{}{
		"price":      price,
		"qty":        qty,
		"updated_by": modifier,
		"updated_at": updater,
	}).Error
}

func (r *Repository) DeletePurchaseReceivingDetail(purchaseId, productId, remover *string, deleter time.Time) error {
	var purchase = new(PurchaseReceivingDetail)

	if err := r.DB.First(purchase, "purchase_receiving_id = ? and product_id = ?", purchaseId, productId).Error; err != nil {
		return err
	}

	return r.DB.Model(purchase).Updates(map[string]interface{}{
		"deleted_by": remover,
		"deleted_at": deleter,
	}).Error
}

// Aksi akan dijalankan ketika menghapus ringkasan pembelian
func (r *Repository) DeletePurchaseReceivingDetail2(purchaseid string, deleter string) error {
	purchaseReceivingDetail := new([]PurchaseReceivingDetail)

	err := r.DB.Where("purchase_receiving_id = ?", purchaseid).Find(purchaseReceivingDetail).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return business.ErrNotFound
		}
		return err
	}

	return r.DB.Model(purchaseReceivingDetail).Updates(map[string]interface{}{
		"deleted_at": time.Now(),
		"deleted_by": deleter,
	}).Error
}

func (r *Repository) GetPurchaseReceivingDetailByPurchaseReceivingId(id string) (*[]purchasereceiving.PurchaseReceivingDetail, error) {
	var details []PurchaseReceivingDetailQuery

	var query = "select t2.code product_code, t2.name product_name, t1.* " +
		"from purchase_receiving_details t1 " +
		"inner join products t2 on t2.id = t1.product_id " +
		"where t1.purchase_receiving_id = '" + id + "' and t1.deleted_by = ''" +
		"order by t1.id;"
	if err := r.DB.Raw(query).Scan(&details).Error; err != nil {
		return nil, err
	}

	var detailsOuts []purchasereceiving.PurchaseReceivingDetail

	for _, value := range details {
		detailsOuts = append(detailsOuts, *value.toPurchaseReceivingDetail())
	}

	return &detailsOuts, nil
}

func (r *Repository) GetPurchaseReceivingDetailById(id string) (*purchasereceiving.PurchaseReceivingDetail, error) {
	var detail purchasereceiving.PurchaseReceivingDetail

	err := r.DB.Where("deleted_by = ''").Where("id = ?", id).First(&detail).Error
	if err != nil {
		return nil, err
	}
	return &detail, nil
}
