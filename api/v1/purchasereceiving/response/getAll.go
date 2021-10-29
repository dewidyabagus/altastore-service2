package response

import (
	"AltaStore/business/purchasereceiving"
	"time"
)

type PurchaseReceiving struct {
	ID           string    `json:"id"`
	Code         string    `json:"code"`
	DateReceived time.Time `json:"datereceived"`
	ReceivedBy   string    `json:"receivedby"`
	Description  string    `json:"description"`
}

type PurchaseReceivings struct {
	PurchaseReceivings []PurchaseReceiving
}

func GetAll(purchasereceivings *[]purchasereceiving.PurchaseReceiving) *PurchaseReceivings {
	var allPurchases PurchaseReceivings
	var purchase PurchaseReceiving
	for _, val := range *purchasereceivings {

		purchase.ID = val.ID
		purchase.Code = val.Code
		purchase.DateReceived = val.DateReceived
		purchase.ReceivedBy = val.ReceivedBy
		purchase.Description = val.Description

		allPurchases.PurchaseReceivings = append(allPurchases.PurchaseReceivings, purchase)
	}

	if allPurchases.PurchaseReceivings == nil {
		allPurchases.PurchaseReceivings = []PurchaseReceiving{}
	}

	return &allPurchases
}
