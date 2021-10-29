package request

type UpdatePurchaseReceivingDetailRequest struct {
	ID        string `json:"id"`
	ProductId string `json:"productid"`
	Qty       int32  `json:"qty"`
	Price     int64  `json:"price"`
	IsDelete  bool   `json:"isdelete"`
}
