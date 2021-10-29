package purchasereceiving

import (
	"time"

	"github.com/google/uuid"
)

type PurchaseReceiving struct {
	ID           string
	Code         string
	DateReceived time.Time
	ReceivedBy   string
	Description  string
	Details      []PurchaseReceivingDetail
	CreatedBy    string
	CreatedAt    time.Time
	UpdatedBy    string
	UpdatedAt    time.Time
	DeletedBy    string
	DeletedAt    time.Time
}

type PurchaseReceivingDetail struct {
	ID                  string
	PurchaseReceivingId string
	ProductId           string
	Qty                 int32
	Price               int64
	CreatedBy           string
	CreatedAt           time.Time
	UpdatedBy           string
	UpdatedAt           time.Time
	DeletedBy           string
	DeletedAt           time.Time
}

func NewPurchaseReceiving(
	newCode string,
	newDateReceived time.Time,
	newReceivedBy string,
	newDescription string,
	creator string,
	createdAt time.Time,
) PurchaseReceiving {
	return PurchaseReceiving{
		ID:           uuid.NewString(),
		Code:         newCode,
		DateReceived: newDateReceived,
		ReceivedBy:   newReceivedBy,
		Description:  newDescription,
		CreatedAt:    createdAt,
		CreatedBy:    creator,
	}
}

func NewPurchaseReceivingDetail(
	purchaseReceivingId string,
	newProductId string,
	newQty int32,
	newPrice int64,
	creator string,
	createdAt time.Time,
) PurchaseReceivingDetail {
	return PurchaseReceivingDetail{
		ID:                  uuid.NewString(),
		PurchaseReceivingId: purchaseReceivingId,
		ProductId:           newProductId,
		Qty:                 newQty,
		Price:               newPrice,
		CreatedBy:           creator,
		CreatedAt:           createdAt,
	}
}

func (oldData *PurchaseReceiving) ModifyPurchaseReceiving(
	newDateReceived time.Time,
	newReceivedBy string,
	newDescription string,
	modifier string,
	modifierAt time.Time,
) PurchaseReceiving {
	return PurchaseReceiving{
		ID:           oldData.ID,
		Code:         oldData.Code,
		DateReceived: newDateReceived,
		ReceivedBy:   newReceivedBy,
		Description:  newDescription,
		CreatedAt:    oldData.CreatedAt,
		CreatedBy:    oldData.CreatedBy,
		UpdatedAt:    modifierAt,
		UpdatedBy:    modifier,
	}
}

func (oldData *PurchaseReceivingDetail) ModifyPurchaseReceivingDetail(
	newProductId string,
	newQty int32,
	newPrice int64,
	modifier string,
	modifierAt time.Time,
) PurchaseReceivingDetail {
	return PurchaseReceivingDetail{
		ID:        oldData.ID,
		ProductId: newProductId,
		Qty:       newQty,
		Price:     newPrice,
		CreatedAt: oldData.CreatedAt,
		CreatedBy: oldData.CreatedBy,
		UpdatedAt: modifierAt,
		UpdatedBy: modifier,
	}
}

func (oldData *PurchaseReceiving) DeletePurchaseReceiving(
	deleter string,
	deletedAt time.Time,
) PurchaseReceiving {
	return PurchaseReceiving{
		ID:           oldData.ID,
		Code:         oldData.Code,
		DateReceived: oldData.DateReceived,
		ReceivedBy:   oldData.ReceivedBy,
		Description:  oldData.Description,
		CreatedAt:    oldData.CreatedAt,
		CreatedBy:    oldData.CreatedBy,
		UpdatedAt:    oldData.UpdatedAt,
		UpdatedBy:    oldData.UpdatedBy,
		DeletedAt:    deletedAt,
		DeletedBy:    deleter,
	}
}

func (oldData *PurchaseReceivingDetail) DeletePurchaseReceivingDetail(
	deleter string,
	deletedAt time.Time,
) PurchaseReceivingDetail {
	return PurchaseReceivingDetail{
		ID:        oldData.ID,
		ProductId: oldData.ProductId,
		Qty:       oldData.Qty,
		Price:     oldData.Price,
		CreatedAt: oldData.CreatedAt,
		CreatedBy: oldData.CreatedBy,
		UpdatedAt: oldData.UpdatedAt,
		UpdatedBy: oldData.UpdatedBy,
		DeletedAt: deletedAt,
		DeletedBy: deleter,
	}
}
