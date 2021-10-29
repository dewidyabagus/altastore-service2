package request

import (
	"AltaStore/business/purchasereceiving"
	"time"
)

type UpdatePurchaseReceivingRequest struct {
	DateReceived time.Time                              `json:"datereceived"`
	ReceivedBy   string                                 `json:"receivedby"`
	Description  string                                 `json:"description"`
	Details      []UpdatePurchaseReceivingDetailRequest `json:"details"`
}

func (i *UpdatePurchaseReceivingRequest) ToPurchaseReceivingSpec() *purchasereceiving.UpdatePurchaseReceivingSpec {
	var spec purchasereceiving.UpdatePurchaseReceivingSpec

	spec.DateReceived = i.DateReceived
	spec.ReceivedBy = i.ReceivedBy
	spec.Description = i.Description

	var detail purchasereceiving.UpdatePurchaseReceivingDetailSpec
	for _, val := range i.Details {
		detail.ID = val.ID
		detail.ProductId = val.ProductId
		detail.Price = val.Price
		detail.Qty = val.Qty
		detail.IsDelete = val.IsDelete

		spec.Details = append(spec.Details, detail)
	}

	return &spec
}
