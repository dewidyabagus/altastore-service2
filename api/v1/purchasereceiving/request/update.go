package request

import (
	"AltaStore/business/purchasereceiving"
)

type DetailPurchaseReceiving struct {
	ProductId string `json:"productid"`
	Price     int    `json:"price"`
	Qty       int    `json:"qty"`
	Status    int    `json:"status"`
}

type UpdatePurchaseReceivingRequest struct {
	Code        string                    `json:"code"`
	Description string                    `json:"description"`
	Details     []DetailPurchaseReceiving `json:"details"`
}

func (i *UpdatePurchaseReceivingRequest) ToPurchaseReceivingSpec() *purchasereceiving.UpdatePurchaseReceivingSpec {
	var spec purchasereceiving.UpdatePurchaseReceivingSpec

	// spec.DateReceived = i.DateReceived
	// spec.ReceivedBy = i.ReceivedBy
	spec.Code = i.Code
	spec.Description = i.Description

	// var detail purchasereceiving.UpdatePurchaseReceivingDetailSpec
	for _, val := range i.Details {
		// detail.ID = val.ID
		// detail.ProductId = val.ProductId
		// detail.Price = val.Price
		// detail.Qty = val.Qty
		// detail.IsDelete = val.IsDelete

		spec.Details = append(spec.Details, purchasereceiving.UpdatePurchaseReceivingDetailSpec{
			ProductId: val.ProductId,
			Qty:       val.Qty,
			Price:     val.Price,
			Status:    val.Status,
		})
	}

	return &spec
}
