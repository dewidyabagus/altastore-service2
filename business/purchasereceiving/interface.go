package purchasereceiving

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
	UpdatePurchaseReceiving(item *PurchaseReceiving) error
	DeletePurchaseReceiving(item *PurchaseReceiving) error
}

type RepositoryDetail interface {
	GetPurchaseReceivingDetailByPurchaseReceivingId(id string) (*[]PurchaseReceivingDetail, error)
	GetPurchaseReceivingDetailById(id string) (*PurchaseReceivingDetail, error)
	InsertPurchaseReceivingDetail(item *PurchaseReceivingDetail) error
	UpdatePurchaseReceivingDetail(item *PurchaseReceivingDetail) error
	DeletePurchaseReceivingDetail(item *PurchaseReceivingDetail) error
}
