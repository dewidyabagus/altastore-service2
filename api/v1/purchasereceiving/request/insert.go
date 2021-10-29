package request

import (
	"AltaStore/business/purchasereceiving"
	"time"
)

type InsertPurchaseReceivingRequest struct {
	Code         string                                 `json:"code"`
	DateReceived time.Time                              `json:"datereceived"`
	ReceivedBy   string                                 `json:"receivedby"`
	Description  string                                 `json:"description"`
	Details      []InsertPurchaseReceivingDetailRequest `json:"details"`
}

func (i *InsertPurchaseReceivingRequest) ToPurchaseReceivingSpec() *purchasereceiving.InsertPurchaseReceivingSpec {
	var spec purchasereceiving.InsertPurchaseReceivingSpec

	spec.Code = i.Code
	spec.DateReceived = i.DateReceived
	spec.ReceivedBy = i.ReceivedBy
	spec.Description = i.Description

	var detail purchasereceiving.InsertPurchaseReceivingDetailSpec
	for _, val := range i.Details {
		detail.ProductId = val.ProductId
		detail.Price = val.Price
		detail.Qty = val.Qty

		spec.Details = append(spec.Details, detail)
	}

	return &spec
}
