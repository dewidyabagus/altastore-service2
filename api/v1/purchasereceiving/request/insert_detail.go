package request

type InsertPurchaseReceivingDetailRequest struct {
	ProductId string `json:"productid"`
	Qty       int32  `json:"qty"`
	Price     int64  `json:"price"`
}
