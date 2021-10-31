package response

import (
	"AltaStore/business/purchasereceiving"
	"time"
)

type PurchaseReceivingById struct {
	ID           string    `json:"id"`
	Code         string    `json:"code"`
	DateReceived time.Time `json:"datereceived"`
	ReceivedBy   string    `json:"receivedby"`
	Description  string    `json:"description"`
	Details      []Detail  `json:"details"`
}

type Detail struct {
	// ID        string `json:"id"`
	ProductId   string `json:"productid"`
	ProductCode string `json:"productcode"`
	ProductName string `json:"productname"`
	Price       int64  `json:"price"`
	Qty         int32  `json:"qty"`
}

func GetById(purchaseReceiving purchasereceiving.PurchaseReceiving) *PurchaseReceivingById {
	var purchase PurchaseReceivingById

	purchase.ID = purchaseReceiving.ID
	purchase.Code = purchaseReceiving.Code
	purchase.DateReceived = purchaseReceiving.DateReceived
	purchase.ReceivedBy = purchaseReceiving.ReceivedBy
	purchase.Description = purchaseReceiving.Description

	// var det Detail
	for _, val := range purchaseReceiving.Details {

		// det.ProductId = val.ProductId
		// det.ID = val.ID
		// det.Qty = val.Qty
		// det.Price = val.Price

		purchase.Details = append(purchase.Details, Detail{
			ProductId:   val.ProductId,
			ProductCode: val.ProductCode,
			ProductName: val.ProductName,
			Price:       val.Price,
			Qty:         val.Qty,
		})
	}

	if purchase.Details == nil {
		purchase.Details = []Detail{}
	}

	return &purchase
}
