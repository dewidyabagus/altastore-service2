package purchasereceiving

import "time"

type Service interface {
	GetAllPurchaseReceivingByParameter(code string, finder string) (*[]PurchaseReceiving, error)
	GetAllPurchaseReceiving(finder string) (*[]PurchaseReceiving, error)
	GetPurchaseReceivingById(id, finder string) (*PurchaseReceiving, error)
	GetPurchaseReceivingByCode(code, finder string) (*PurchaseReceiving, error)
	InsertPurchaseReceiving(item *InsertPurchaseReceivingSpec, creator string) error
	UpdatePurchaseReceiving(id string, item *UpdatePurchaseReceivingSpec, modifier string) error
	DeletePurchaseReceiving(id string, deleter string) error
}

type Repository interface {
	GetAllPurchaseReceivingByParameter(code string) (*[]PurchaseReceiving, error)
	GetAllPurchaseReceiving() (*[]PurchaseReceiving, error)
	GetPurchaseReceivingById(id string) (*PurchaseReceiving, error)
	GetPurchaseReceivingByCode(code string) (*PurchaseReceiving, error)
	InsertPurchaseReceiving(item *PurchaseReceiving) error
	// UpdatePurchaseReceiving(item *PurchaseReceiving) error
	UpdatePurchaseReceiving2(id, code, description, modifier *string, updater time.Time) error
	DeletePurchaseReceiving(item *PurchaseReceiving) error
}

type RepositoryDetail interface {
	GetPurchaseReceivingDetailByPurchaseReceivingId(id string) (*[]PurchaseReceivingDetail, error)
	GetPurchaseReceivingDetailById(id string) (*PurchaseReceivingDetail, error)
	InsertPurchaseReceivingDetail(item *PurchaseReceivingDetail) error
	// UpdatePurchaseReceivingDetail(item *PurchaseReceivingDetail) error
	UpdatePurchaseReceivingDetail(purchaseId, productId *string, price, qty *int, modifier *string, updater time.Time) error
	DeletePurchaseReceivingDetail(purchaseId, productId, remover *string, deleter time.Time) error
	// DeletePurchaseReceivingDetail(item *PurchaseReceivingDetail) error
	DeletePurchaseReceivingDetail2(id string, deleter string) error // delete detail purchase receiving by purchase id
}
