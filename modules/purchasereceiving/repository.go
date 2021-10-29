package purchasereceiving

import (
	"AltaStore/business/purchasereceiving"
	"time"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

type PurchaseReceiving struct {
	ID           string    `gorm:"id;type:uuid;primaryKey"`
	Code         string    `gorm:"code;type:varchar(50);unique"`
	DateReceived time.Time `gorm:"datereceived;type:timestamp"`
	ReceivedBy   string    `gorm:"receivedby;type:varchar(50)"`
	Description  string    `gorm:"description;type:varchar(100)"`
	CreatedAt    time.Time `gorm:"created_at"`
	CreatedBy    string    `gorm:"created_by;type:varchar(50)"`
	UpdatedAt    time.Time `gorm:"updated_at"`
	UpdatedBy    string    `gorm:"updated_by;type:varchar(50)"`
	DeletedAt    time.Time `gorm:"deleted_at"`
	DeletedBy    string    `gorm:"deleted_by;type:varchar(50)"`
}

func (p *PurchaseReceiving) toPurchaseReceiving() *purchasereceiving.PurchaseReceiving {
	return &purchasereceiving.PurchaseReceiving{
		ID:           p.ID,
		Code:         p.Code,
		DateReceived: p.DateReceived,
		ReceivedBy:   p.ReceivedBy,
		Description:  p.Description,
		CreatedAt:    p.CreatedAt,
		CreatedBy:    p.CreatedBy,
		UpdatedBy:    p.UpdatedBy,
		UpdatedAt:    p.UpdatedAt,
		DeletedBy:    p.DeletedBy,
		DeletedAt:    p.DeletedAt,
	}
}

func newPurchaseReceiving(
	p *purchasereceiving.PurchaseReceiving,
) *PurchaseReceiving {
	return &PurchaseReceiving{
		ID:           p.ID,
		Code:         p.Code,
		DateReceived: p.DateReceived,
		ReceivedBy:   p.ReceivedBy,
		Description:  p.Description,
		CreatedAt:    p.CreatedAt,
		CreatedBy:    p.CreatedBy,
		UpdatedAt:    p.UpdatedAt,
		UpdatedBy:    p.UpdatedBy,
		DeletedAt:    p.DeletedAt,
		DeletedBy:    p.DeletedBy,
	}
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) InsertPurchaseReceiving(p *purchasereceiving.PurchaseReceiving) error {
	purchase := newPurchaseReceiving(p)
	if err := r.DB.Create(purchase).Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdatePurchaseReceiving(p *purchasereceiving.PurchaseReceiving) error {
	purchase := newPurchaseReceiving(p)
	err := r.DB.Model(&purchase).Updates(PurchaseReceiving{
		DateReceived: purchase.DateReceived,
		ReceivedBy:   purchase.ReceivedBy,
		Description:  purchase.Description,
		UpdatedAt:    purchase.UpdatedAt,
		UpdatedBy:    purchase.UpdatedBy,
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) DeletePurchaseReceiving(p *purchasereceiving.PurchaseReceiving) error {
	purchase := newPurchaseReceiving(p)

	err := r.DB.Model(&purchase).Updates(PurchaseReceiving{
		DeletedBy: purchase.DeletedBy,
		DeletedAt: purchase.DeletedAt,
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetAllPurchaseReceivingByParameter(code string) (*[]purchasereceiving.PurchaseReceiving, error) {
	var purchases []PurchaseReceiving
	var tempPurchase purchasereceiving.PurchaseReceiving
	purchasetOuts := []purchasereceiving.PurchaseReceiving{}

	temp := r.DB

	if code != "" {
		temp = temp.Where("code = ?", code)
	}
	err := temp.Where("deleted_by = ''").Find(&purchases).Error
	if err != nil {
		return &purchasetOuts, nil
	}

	for _, value := range purchases {
		tempPurchase = *value.toPurchaseReceiving()
		purchasetOuts = append(purchasetOuts, tempPurchase)
	}

	return &purchasetOuts, nil
}
func (r *Repository) GetAllPurchaseReceiving() (*[]purchasereceiving.PurchaseReceiving, error) {
	var purchases []PurchaseReceiving
	var tempPurchase purchasereceiving.PurchaseReceiving
	purchasetOuts := []purchasereceiving.PurchaseReceiving{}

	temp := r.DB
	err := temp.Where("deleted_by = ''").Find(&purchases).Error
	if err != nil {
		return &purchasetOuts, nil
	}

	for _, value := range purchases {
		tempPurchase = *value.toPurchaseReceiving()
		purchasetOuts = append(purchasetOuts, tempPurchase)
	}

	return &purchasetOuts, nil
}

func (r *Repository) GetPurchaseReceivingById(id string) (*purchasereceiving.PurchaseReceiving, error) {
	var purchase PurchaseReceiving

	err := r.DB.Where("id = ?", id).Where("deleted_by = ''").First(&purchase).Error
	if err != nil {
		return nil, err
	}

	return purchase.toPurchaseReceiving(), nil
}

func (r *Repository) GetPurchaseReceivingByCode(code string) (*purchasereceiving.PurchaseReceiving, error) {
	var purchase PurchaseReceiving

	err := r.DB.Where("code = ?", code).Where("deleted_by = ''").First(&purchase).Error
	if err != nil {
		return nil, err
	}

	return purchase.toPurchaseReceiving(), nil
}
