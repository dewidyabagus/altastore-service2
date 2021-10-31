package purchasereceiving_test

import (
	"AltaStore/business"
	"AltaStore/business/purchasereceiving"
	purchaseReceivingMock "AltaStore/business/purchasereceiving/mocks"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	adminId           = "f9c8c2bf-d525-420e-86e5-4caf03cd8027"
	email             = "email"
	id                = "f9c8c2bf-d525-420e-86e5-4caf03cd8027"
	code              = "code"
	name              = "name"
	receivedby        = "receivedby"
	description       = "description"
	productid         = "f9c8c2bf-d525-420e-86e5-4caf03cd8027"
	isdelete          = false
	qty         int32 = 10
	price       int64 = 100000000
)

var (
	datereceived                      = time.Now()
	purchaseReceivingRepository       purchaseReceivingMock.Repository
	purchaseReceivingDetailRepository purchaseReceivingMock.RepositoryDetail
	purchaseReceivingService          purchasereceiving.Service

	purchaseReceivingData              purchasereceiving.PurchaseReceiving
	purchaseReceivingDetailData        purchasereceiving.PurchaseReceivingDetail
	purchaseReceivingDatas             []*purchasereceiving.PurchaseReceiving
	purchaseReceivingGetDatas          []purchasereceiving.PurchaseReceiving
	purchaseReceivingDetailDatas       []purchasereceiving.PurchaseReceivingDetail
	insertPurchaseReceivingSpec        *purchasereceiving.InsertPurchaseReceivingSpec
	insertPurchaseReceivingDetailSpec  purchasereceiving.InsertPurchaseReceivingDetailSpec
	insertPurchaseReceivingDetailSpecs []purchasereceiving.InsertPurchaseReceivingDetailSpec
	updatePurchaseReceivingSpec        *purchasereceiving.UpdatePurchaseReceivingSpec
	updatePurchaseReceivingDetailSpec  purchasereceiving.UpdatePurchaseReceivingDetailSpec
	updatePurchaseReceivingDetailSpecs []purchasereceiving.UpdatePurchaseReceivingDetailSpec
)

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func setup() {
	purchaseReceivingDetailData = purchasereceiving.PurchaseReceivingDetail{
		ID:                  id,
		PurchaseReceivingId: id,
		ProductId:           productid,
		Qty:                 qty,
		Price:               price,
	}
	purchaseReceivingDetailDatas = append(purchaseReceivingDetailDatas, purchaseReceivingDetailData)
	purchaseReceivingData = purchasereceiving.PurchaseReceiving{
		ID:           id,
		Code:         code,
		DateReceived: datereceived,
		ReceivedBy:   receivedby,
		Description:  description,
		Details:      purchaseReceivingDetailDatas,
	}
	purchaseReceivingDatas = append(purchaseReceivingDatas, &purchaseReceivingData)
	purchaseReceivingGetDatas = append(purchaseReceivingGetDatas, purchaseReceivingData)

	insertPurchaseReceivingDetailSpec = purchasereceiving.InsertPurchaseReceivingDetailSpec{
		ProductId: productid,
		Qty:       qty,
		Price:     price,
	}
	insertPurchaseReceivingDetailSpecs = append(insertPurchaseReceivingDetailSpecs, insertPurchaseReceivingDetailSpec)

	insertPurchaseReceivingSpec = &purchasereceiving.InsertPurchaseReceivingSpec{
		Code:         code,
		DateReceived: time.Now(),
		ReceivedBy:   receivedby,
		Description:  description,
		Details:      insertPurchaseReceivingDetailSpecs,
	}

	updatePurchaseReceivingDetailSpec = purchasereceiving.UpdatePurchaseReceivingDetailSpec{
		ID:        id,
		ProductId: productid,
		Qty:       qty,
		Price:     price,
		IsDelete:  isdelete,
	}
	updatePurchaseReceivingDetailSpecs = append(updatePurchaseReceivingDetailSpecs, updatePurchaseReceivingDetailSpec)

	updatePurchaseReceivingSpec = &purchasereceiving.UpdatePurchaseReceivingSpec{
		DateReceived: datereceived,
		ReceivedBy:   receivedby,
		Description:  description,
		Details:      updatePurchaseReceivingDetailSpecs,
	}
	purchaseReceivingService = purchasereceiving.NewService(&purchaseReceivingRepository, &purchaseReceivingDetailRepository)
}

func TestInsertPurchaseReceiving(t *testing.T) {
	// t.Run("Expect Admin Not Found", func(t *testing.T) {
	// 	adminService.On("FindAdminByID", mock.AnythingOfType("string")).Return(nil, business.ErrNotHavePermission).Once()

	// 	err := purchaseReceivingService.InsertPurchaseReceiving(insertPurchaseReceivingSpec, email)

	// 	assert.NotNil(t, err)
	// 	assert.Equal(t, err, business.ErrNotHavePermission)

	// })
	t.Run("Expect Purchase Receiving Exist", func(t *testing.T) {
		//adminService.On("FindAdminByID", mock.AnythingOfType("string")).Return(&adminData, nil).Once()
		purchaseReceivingRepository.On("GetPurchaseReceivingByCode", mock.AnythingOfType("string")).Return(&purchaseReceivingData, nil).Once()

		err := purchaseReceivingService.InsertPurchaseReceiving(insertPurchaseReceivingSpec, email)

		assert.NotNil(t, err)

		assert.Equal(t, err, business.ErrDataExists)
	})
	t.Run("Expect Insert Purchase Receiving Success", func(t *testing.T) {
		//adminService.On("FindAdminByID", mock.AnythingOfType("string")).Return(&adminData, nil).Once()
		purchaseReceivingRepository.On("GetPurchaseReceivingByCode", mock.AnythingOfType("string")).Return(nil, nil).Once()
		purchaseReceivingDetailRepository.On("InsertPurchaseReceivingDetail", mock.AnythingOfType("*purchasereceiving.PurchaseReceivingDetail")).Return(nil).Once()
		purchaseReceivingRepository.On("InsertPurchaseReceiving", mock.AnythingOfType("*purchasereceiving.PurchaseReceiving")).Return(nil).Once()

		err := purchaseReceivingService.InsertPurchaseReceiving(insertPurchaseReceivingSpec, email)

		assert.Nil(t, err)
	})
	t.Run("Expect Insert Purchase Receiving Fail", func(t *testing.T) {
		//adminService.On("FindAdminByID", mock.AnythingOfType("string")).Return(&adminData, nil).Once()
		purchaseReceivingRepository.On("GetPurchaseReceivingByCode", mock.AnythingOfType("string")).Return(nil, nil).Once()
		purchaseReceivingDetailRepository.On("InsertPurchaseReceivingDetail", mock.AnythingOfType("*purchasereceiving.PurchaseReceivingDetail")).Return(business.ErrInternalServer).Once()
		purchaseReceivingRepository.On("InsertPurchaseReceiving", mock.AnythingOfType("*purchasereceiving.PurchaseReceiving")).Return(business.ErrInternalServer).Once()

		err := purchaseReceivingService.InsertPurchaseReceiving(insertPurchaseReceivingSpec, email)

		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrInternalServer)
	})
}

func TestUpdatePurchaseReceiving(t *testing.T) {
	// t.Run("Expect Admin Not Found", func(t *testing.T) {
	// 	adminService.On("FindAdminByID", mock.AnythingOfType("string")).Return(nil, business.ErrNotHavePermission).Once()

	// 	err := purchaseReceivingService.UpdatePurchaseReceiving(id, updatePurchaseReceivingSpec, email)

	// 	assert.NotNil(t, err)
	// 	assert.Equal(t, err, business.ErrNotHavePermission)

	// })
	t.Run("Expect Purchase Receiving Not Found", func(t *testing.T) {
		//adminService.On("FindAdminByID", mock.AnythingOfType("string")).Return(&adminData, nil).Once()
		purchaseReceivingRepository.On("GetPurchaseReceivingById", mock.AnythingOfType("string")).Return(nil, business.ErrNotFound).Once()

		err := purchaseReceivingService.UpdatePurchaseReceiving(id, updatePurchaseReceivingSpec, email)

		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrNotFound)

	})
	t.Run("Expect Update Purchase Receiving Success", func(t *testing.T) {
		//adminService.On("FindAdminByID", mock.AnythingOfType("string")).Return(&adminData, nil).Once()
		purchaseReceivingRepository.On("GetPurchaseReceivingById", mock.AnythingOfType("string")).Return(&purchaseReceivingData, nil).Once()
		purchaseReceivingDetailRepository.On("UpdatePurchaseReceivingDetail", mock.AnythingOfType("*purchasereceiving.PurchaseReceivingDetail")).Return(nil).Once()
		purchaseReceivingRepository.On("UpdatePurchaseReceiving", mock.AnythingOfType("*purchasereceiving.PurchaseReceiving")).Return(nil).Once()
		purchaseReceivingDetailRepository.On("GetPurchaseReceivingDetailById", mock.AnythingOfType("string")).Return(&purchaseReceivingDetailData, nil).Once()
		purchaseReceivingRepository.On("InsertPurchaseReceivingDetail", mock.AnythingOfType("*purchasereceiving.PurchaseReceivingDetail")).Return(nil).Once()
		purchaseReceivingRepository.On("UpdatePurchaseReceivingDetail", mock.AnythingOfType("*purchasereceiving.PurchaseReceivingDetail")).Return(nil).Once()
		purchaseReceivingRepository.On("DeletePurchaseReceivingDetail", mock.AnythingOfType("*purchasereceiving.PurchaseReceivingDetail")).Return(nil).Once()

		err := purchaseReceivingService.UpdatePurchaseReceiving(id, updatePurchaseReceivingSpec, email)

		assert.Nil(t, err)
	})
	t.Run("Expect Update Purchase Receiving Fail", func(t *testing.T) {
		//adminService.On("FindAdminByID", mock.AnythingOfType("string")).Return(&adminData, nil).Once()
		purchaseReceivingRepository.On("GetPurchaseReceivingById", mock.AnythingOfType("string")).Return(&purchaseReceivingData, nil).Once()
		purchaseReceivingDetailRepository.On("UpdatePurchaseReceivingDetail", mock.AnythingOfType("*purchasereceiving.PurchaseReceivingDetail")).Return(business.ErrInternalServer).Once()
		purchaseReceivingRepository.On("UpdatePurchaseReceiving", mock.AnythingOfType("*purchasereceiving.PurchaseReceiving")).Return(business.ErrInternalServer).Once()
		purchaseReceivingDetailRepository.On("GetPurchaseReceivingDetailById", mock.AnythingOfType("string")).Return(nil, business.ErrInternalServer).Once()
		purchaseReceivingRepository.On("InsertPurchaseReceivingDetail", mock.AnythingOfType("*purchasereceiving.PurchaseReceivingDetail")).Return(business.ErrInternalServer).Once()
		purchaseReceivingRepository.On("UpdatePurchaseReceivingDetail", mock.AnythingOfType("*purchasereceiving.PurchaseReceivingDetail")).Return(business.ErrInternalServer).Once()
		purchaseReceivingRepository.On("DeletePurchaseReceivingDetail", mock.AnythingOfType("*purchasereceiving.PurchaseReceivingDetail")).Return(business.ErrInternalServer).Once()

		err := purchaseReceivingService.UpdatePurchaseReceiving(id, updatePurchaseReceivingSpec, email)

		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrInternalServer)
	})
}

func TestDeletePurchaseReceiving(t *testing.T) {
	// t.Run("Expect Admin Not Found", func(t *testing.T) {
	// 	adminService.On("FindAdminByID", mock.AnythingOfType("string")).Return(nil, business.ErrNotHavePermission).Once()

	// 	err := purchaseReceivingService.DeletePurchaseReceiving(id, adminId)

	// 	assert.NotNil(t, err)
	// 	assert.Equal(t, err, business.ErrNotHavePermission)

	// })
	t.Run("Expect Purchase Receiving Not Found", func(t *testing.T) {
		//adminService.On("FindAdminByID", mock.AnythingOfType("string")).Return(&adminData, nil).Once()
		purchaseReceivingRepository.On("GetPurchaseReceivingById", mock.AnythingOfType("string")).Return(nil, business.ErrNotFound).Once()

		err := purchaseReceivingService.DeletePurchaseReceiving(id, adminId)

		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrNotFound)

	})
	t.Run("Expect Delete Purchase Receiving Detail Fail", func(t *testing.T) {
		//adminService.On("FindAdminByID", mock.AnythingOfType("string")).Return(&adminData, nil).Once()
		purchaseReceivingRepository.On("GetPurchaseReceivingById", mock.AnythingOfType("string")).Return(&purchaseReceivingData, nil).Once()
		purchaseReceivingDetailRepository.On("DeletePurchaseReceivingDetail2", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(business.ErrInternalServer).Once()

		err := purchaseReceivingService.DeletePurchaseReceiving(id, adminId)

		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrInternalServer)
	})
	t.Run("Expect Delete Purchase Receiving Fail", func(t *testing.T) {
		//adminService.On("FindAdminByID", mock.AnythingOfType("string")).Return(&adminData, nil).Once()
		purchaseReceivingRepository.On("GetPurchaseReceivingById", mock.AnythingOfType("string")).Return(&purchaseReceivingData, nil).Once()
		purchaseReceivingDetailRepository.On("DeletePurchaseReceivingDetail2", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil).Once()
		purchaseReceivingRepository.On("DeletePurchaseReceiving", mock.AnythingOfType("*purchasereceiving.PurchaseReceiving")).Return(business.ErrInternalServer).Once()

		err := purchaseReceivingService.DeletePurchaseReceiving(id, adminId)

		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrInternalServer)
	})
	t.Run("Expect Delete Purchase Receiving Success", func(t *testing.T) {
		//adminService.On("FindAdminByID", mock.AnythingOfType("string")).Return(&adminData, nil).Once()
		purchaseReceivingRepository.On("GetPurchaseReceivingById", mock.AnythingOfType("string")).Return(&purchaseReceivingData, nil).Once()
		purchaseReceivingDetailRepository.On("DeletePurchaseReceivingDetail2", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil).Once()
		purchaseReceivingRepository.On("DeletePurchaseReceiving", mock.AnythingOfType("*purchasereceiving.PurchaseReceiving")).Return(nil).Once()

		err := purchaseReceivingService.DeletePurchaseReceiving(id, adminId)

		assert.Nil(t, err)
	})

}

func TestGetAllPurchaseReceiving(t *testing.T) {
	// t.Run("Expect Admin Not Found", func(t *testing.T) {
	// 	adminService.On("FindAdminByID", mock.AnythingOfType("string")).Return(nil, business.ErrNotHavePermission).Once()

	// 	_, err := purchaseReceivingService.GetAllPurchaseReceiving(email)

	// 	assert.NotNil(t, err)
	// 	assert.Equal(t, err, business.ErrNotHavePermission)

	// })
	t.Run("Expect found the data Purchase Receiving", func(t *testing.T) {
		//adminService.On("FindAdminByID", mock.AnythingOfType("string")).Return(&adminData, nil).Once()
		purchaseReceivingRepository.On("GetAllPurchaseReceiving", mock.AnythingOfType("string")).Return(&purchaseReceivingGetDatas, nil).Once()

		receivings, err := purchaseReceivingService.GetAllPurchaseReceiving(email)

		assert.Nil(t, err)
		assert.NotNil(t, receivings)

		assert.Equal(t, id, (*receivings)[0].ID)
		assert.Equal(t, code, (*receivings)[0].Code)
		assert.Equal(t, datereceived, (*receivings)[0].DateReceived)
		assert.Equal(t, receivedby, (*receivings)[0].ReceivedBy)
		assert.Equal(t, description, (*receivings)[0].Description)
	})

	t.Run("Expect data nil", func(t *testing.T) {
		//adminService.On("FindAdminByID", mock.AnythingOfType("string")).Return(&adminData, nil).Once()
		purchaseReceivingRepository.On("GetAllPurchaseReceiving", mock.AnythingOfType("string")).Return(nil, nil).Once()
		categories, err := purchaseReceivingService.GetAllPurchaseReceiving(email)

		assert.Nil(t, err)
		assert.Nil(t, categories)
	})
}

func TestGetAllPurchaseReceivingByParameter(t *testing.T) {
	// t.Run("Expect Admin Not Found", func(t *testing.T) {
	// 	adminService.On("FindAdminByID", mock.AnythingOfType("string")).Return(nil, business.ErrNotHavePermission).Once()

	// 	_, err := purchaseReceivingService.GetAllPurchaseReceivingByParameter(code, email)

	// 	assert.NotNil(t, err)
	// 	assert.Equal(t, err, business.ErrNotHavePermission)

	// })
	t.Run("Expect found the data Purchase Receiving", func(t *testing.T) {
		//adminService.On("FindAdminByID", mock.AnythingOfType("string")).Return(&adminData, nil).Once()
		purchaseReceivingRepository.On("GetAllPurchaseReceivingByParameter", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(&purchaseReceivingGetDatas, nil).Once()

		receivings, err := purchaseReceivingService.GetAllPurchaseReceivingByParameter(code, email)

		assert.Nil(t, err)
		assert.NotNil(t, receivings)

		assert.Equal(t, id, (*receivings)[0].ID)
		assert.Equal(t, code, (*receivings)[0].Code)
		assert.Equal(t, datereceived, (*receivings)[0].DateReceived)
		assert.Equal(t, receivedby, (*receivings)[0].ReceivedBy)
		assert.Equal(t, description, (*receivings)[0].Description)
	})

	t.Run("Expect data nil", func(t *testing.T) {
		//adminService.On("FindAdminByID", mock.AnythingOfType("string")).Return(&adminData, nil).Once()
		purchaseReceivingRepository.On("GetAllPurchaseReceivingByParameter", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil, nil).Once()
		categories, err := purchaseReceivingService.GetAllPurchaseReceivingByParameter(code, email)

		assert.Nil(t, err)
		assert.Nil(t, categories)

	})
}

func TestGetPurchaseReceivingById(t *testing.T) {
	// t.Run("Expect Admin Not Found", func(t *testing.T) {
	// 	adminService.On("FindAdminByID", mock.AnythingOfType("string")).Return(nil, business.ErrNotHavePermission).Once()

	// 	purchasereceiving, err := purchaseReceivingService.GetPurchaseReceivingById(id, email)

	// 	assert.Nil(t, purchasereceiving)
	// 	assert.NotNil(t, err)
	// 	assert.Equal(t, err, business.ErrNotHavePermission)

	// })
	t.Run("Expect Purchase Receiving not found", func(t *testing.T) {
		//adminService.On("FindAdminByID", mock.AnythingOfType("string")).Return(&adminData, nil).Once()
		purchaseReceivingRepository.On("GetPurchaseReceivingById", mock.AnythingOfType("string")).Return(nil, business.ErrNotFound).Once()
		purchasereceiving, err := purchaseReceivingService.GetPurchaseReceivingById(id, email)

		assert.NotNil(t, err)
		assert.Nil(t, purchasereceiving)

		assert.Equal(t, err, business.ErrNotFound)
	})
	t.Run("Expect Purchase Receiving Detail Not Found", func(t *testing.T) {
		//adminService.On("FindAdminByID", mock.AnythingOfType("string")).Return(&adminData, nil).Once()
		purchaseReceivingRepository.On("GetPurchaseReceivingById", mock.AnythingOfType("string")).Return(&purchaseReceivingData, nil).Once()
		purchaseReceivingDetailRepository.On("GetPurchaseReceivingDetailByPurchaseReceivingId", mock.AnythingOfType("string")).Return(nil, business.ErrNotFound).Once()

		_, err := purchaseReceivingService.GetPurchaseReceivingById(id, email)

		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrNotFound)

	})
	t.Run("Expect Purchase Receiving Detail Not Found", func(t *testing.T) {
		//adminService.On("FindAdminByID", mock.AnythingOfType("string")).Return(&adminData, nil).Once()
		purchaseReceivingRepository.On("GetPurchaseReceivingById", mock.AnythingOfType("string")).Return(&purchaseReceivingData, nil).Once()
		purchaseReceivingDetailRepository.On("GetPurchaseReceivingDetailByPurchaseReceivingId", mock.AnythingOfType("string")).Return(&purchaseReceivingDetailDatas, nil).Once()

		purchasereceiving, err := purchaseReceivingService.GetPurchaseReceivingById(id, email)

		assert.Nil(t, err)
		assert.NotNil(t, purchasereceiving)

		assert.Equal(t, id, purchasereceiving.ID)
		assert.Equal(t, code, purchasereceiving.Code)
		assert.Equal(t, datereceived, purchasereceiving.DateReceived)
		assert.Equal(t, receivedby, purchasereceiving.ReceivedBy)
		assert.Equal(t, description, purchasereceiving.Description)

		assert.Equal(t, id, purchasereceiving.Details[0].ID)
		assert.Equal(t, price, purchasereceiving.Details[0].Price)
		assert.Equal(t, productid, purchasereceiving.Details[0].ProductId)
		assert.Equal(t, qty, purchasereceiving.Details[0].Qty)
	})

}

func TestGetPurchaseReceivingByCode(t *testing.T) {
	// t.Run("Expect Admin Not Found", func(t *testing.T) {
	// 	adminService.On("FindAdminByID", mock.AnythingOfType("string")).Return(nil, business.ErrNotHavePermission).Once()

	// 	purchasereceiving, err := purchaseReceivingService.GetPurchaseReceivingByCode(code, email)

	// 	assert.Nil(t, purchasereceiving)
	// 	assert.NotNil(t, err)
	// 	assert.Equal(t, err, business.ErrNotHavePermission)
	// })
	t.Run("Expect found the Purchase Receiving", func(t *testing.T) {
		//adminService.On("FindAdminByID", mock.AnythingOfType("string")).Return(&adminData, nil).Once()
		purchaseReceivingRepository.On("GetPurchaseReceivingByCode", mock.AnythingOfType("string")).Return(&purchaseReceivingData, nil).Once()

		purchasereceiving, err := purchaseReceivingService.GetPurchaseReceivingByCode(code, email)

		assert.Nil(t, err)
		assert.NotNil(t, purchasereceiving)

		assert.Equal(t, id, purchasereceiving.ID)
		assert.Equal(t, code, purchasereceiving.Code)
		assert.Equal(t, datereceived, purchasereceiving.DateReceived)
		assert.Equal(t, receivedby, purchasereceiving.ReceivedBy)
		assert.Equal(t, description, purchasereceiving.Description)
		assert.Equal(t, id, (purchasereceiving.Details)[0].ID)
		assert.Equal(t, productid, (purchasereceiving.Details)[0].ProductId)
		assert.Equal(t, qty, (purchasereceiving.Details)[0].Qty)
		assert.Equal(t, price, (purchasereceiving.Details)[0].Price)
	})

	t.Run("Expect Purchase Receiving not found", func(t *testing.T) {
		//adminService.On("FindAdminByID", mock.AnythingOfType("string")).Return(&adminData, nil).Once()
		purchaseReceivingRepository.On("GetPurchaseReceivingByCode", mock.AnythingOfType("string")).Return(nil, business.ErrNotFound).Once()
		purchasereceiving, err := purchaseReceivingService.GetPurchaseReceivingByCode(code, email)

		assert.NotNil(t, err)
		assert.Nil(t, purchasereceiving)

		assert.Equal(t, err, business.ErrNotFound)
	})
}
