package purchasereceivingdetail

import (
	"AltaStore/business/purchasereceiving"
	"time"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

type PurchaseReceivingDetail struct {
	ID                  string    `gorm:"id;type:uuid;primaryKey"`
	PurchaseReceivingId string    `gorm:"purchase_receiving_id;type:varchar(50)"`
	ProductId           string    `gorm:"product_id;type:varchar(50)"`
	Price               int64     `gorm:"price;"`
	Qty                 int32     `gorm:"qty;"`
	CreatedAt           time.Time `gorm:"created_at"`
	CreatedBy           string    `gorm:"created_by;type:varchar(50)"`
	UpdatedAt           time.Time `gorm:"updated_at"`
	UpdatedBy           string    `gorm:"updated_by;type:varchar(50)"`
	DeletedAt           time.Time `gorm:"deleted_at"`
	DeletedBy           string    `gorm:"deleted_by;type:varchar(50)"`
}

func (p *PurchaseReceivingDetail) toPurchaseReceivingDetail() *purchasereceiving.PurchaseReceivingDetail {
	return &purchasereceiving.PurchaseReceivingDetail{
		ID:        p.ID,
		ProductId: p.ProductId,
		Price:     p.Price,
		Qty:       p.Qty,
		CreatedBy: p.CreatedBy,
		CreatedAt: p.CreatedAt,
		UpdatedBy: p.UpdatedBy,
		UpdatedAt: p.UpdatedAt,
		DeletedBy: p.DeletedBy,
		DeletedAt: p.DeletedAt,
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

func (r *Repository) UpdatePurchaseReceivingDetail(p *purchasereceiving.PurchaseReceivingDetail) error {
	detail := newDataPurchaseReceivingDetail(p)
	err := r.DB.Model(&detail).Updates(PurchaseReceivingDetail{
		ProductId: detail.ProductId,
		Price:     detail.Price,
		Qty:       detail.Qty,
		UpdatedAt: detail.UpdatedAt,
		UpdatedBy: detail.UpdatedBy,
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) DeletePurchaseReceivingDetail(p *purchasereceiving.PurchaseReceivingDetail) error {
	detail := newDataPurchaseReceivingDetail(p)

	err := r.DB.Model(&detail).Updates(PurchaseReceivingDetail{
		DeletedBy: detail.DeletedBy,
		DeletedAt: detail.DeletedAt,
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetPurchaseReceivingDetailByPurchaseReceivingId(id string) (*[]purchasereceiving.PurchaseReceivingDetail, error) {
	var details []PurchaseReceivingDetail
	var tempDetail purchasereceiving.PurchaseReceivingDetail
	var detailsOuts []purchasereceiving.PurchaseReceivingDetail

	err := r.DB.Where("purchase_receiving_id = ?", id).Where("deleted_by = ''").Order("created_at asc").Order("updated_at asc").Find(&details).Error
	if err != nil {
		return nil, err
	}

	for _, value := range details {
		tempDetail = *value.toPurchaseReceivingDetail()
		detailsOuts = append(detailsOuts, tempDetail)
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
