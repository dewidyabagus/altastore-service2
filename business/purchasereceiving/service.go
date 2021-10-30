package purchasereceiving

import (
	"AltaStore/business"
	// "AltaStore/business/admin"
	"AltaStore/util/validator"
	"time"
)

type InsertPurchaseReceivingSpec struct {
	Code         string    `validate:"required"`
	DateReceived time.Time `validate:"required"`
	ReceivedBy   string    `validate:"required"`
	Description  string
	Details      []InsertPurchaseReceivingDetailSpec `validate:"required"`
}

type InsertPurchaseReceivingDetailSpec struct {
	ProductId string `validate:"required"`
	Qty       int32  `validate:"required"`
	Price     int64  `validate:"required"`
}

type UpdatePurchaseReceivingSpec struct {
	DateReceived time.Time `validate:"required"`
	ReceivedBy   string    `validate:"required"`
	Description  string
	Details      []UpdatePurchaseReceivingDetailSpec `validate:"required"`
}
type UpdatePurchaseReceivingDetailSpec struct {
	ID        string
	ProductId string `validate:"required"`
	Qty       int32  `validate:"required"`
	Price     int64  `validate:"required"`
	IsDelete  bool   `validate:"required"`
}

type service struct {
	// adminService     admin.Service
	repository       Repository
	repositoryDetail RepositoryDetail
}

// func NewService(
// 	adminService admin.Service,
// 	repository Repository,
// 	repositoryDetail RepositoryDetail,
// ) Service {
// 	return &service{
// 		adminService, repository, repositoryDetail,
// 	}
// }

func NewService(
	repository Repository,
	repositoryDetail RepositoryDetail,
) Service {
	return &service{repository, repositoryDetail}
}

// GetAllPurchaseReceiving(finder string) (*PurchaseReceiving, error)
// GetAllPurchaseReceivingById(id, finder string) (*PurchaseReceiving, error)

func (s *service) GetAllPurchaseReceivingByParameter(code string, finder string) (*[]PurchaseReceiving, error) {
	// _, err := s.adminService.FindAdminByID(finder)
	// if err != nil {
	// 	empty := []PurchaseReceiving{}
	// 	return &empty, business.ErrNotHavePermission
	// }
	return s.repository.GetAllPurchaseReceivingByParameter(code)
}

func (s *service) GetAllPurchaseReceiving(finder string) (*[]PurchaseReceiving, error) {
	// _, err := s.adminService.FindAdminByID(finder)
	// if err != nil {
	// 	empty := []PurchaseReceiving{}
	// 	return &empty, business.ErrNotHavePermission
	// }
	return s.repository.GetAllPurchaseReceiving()
}

func (s *service) GetPurchaseReceivingById(id, finder string) (*PurchaseReceiving, error) {
	// _, err := s.adminService.FindAdminByID(finder)
	// if err != nil {
	// 	return nil, business.ErrNotHavePermission
	// }
	purchase, err := s.repository.GetPurchaseReceivingById(id)
	if err != nil {
		return nil, err
	}

	details, err := s.repositoryDetail.GetPurchaseReceivingDetailByPurchaseReceivingId(purchase.ID)
	if err != nil {
		return nil, err
	}
	purchase.Details = append(purchase.Details, *details...)
	return purchase, nil
}

func (s *service) GetPurchaseReceivingByCode(code, finder string) (*PurchaseReceiving, error) {
	// _, err := s.adminService.FindAdminByID(finder)
	// if err != nil {
	// 	return nil, business.ErrNotHavePermission
	// }
	return s.repository.GetPurchaseReceivingByCode(code)
}

func (s *service) InsertPurchaseReceiving(item *InsertPurchaseReceivingSpec, creator string) error {
	err := validator.GetValidator().Struct(item)
	if err != nil {
		return business.ErrInvalidSpec
	}

	// admin, err := s.adminService.FindAdminByID(creator)
	// if err != nil {
	// 	return business.ErrNotHavePermission
	// }

	data, _ := s.repository.GetPurchaseReceivingByCode(item.Code)
	if data != nil {
		return business.ErrDataExists
	}

	newItem := NewPurchaseReceiving(
		item.Code,
		item.DateReceived,
		item.ReceivedBy,
		item.Description,
		//admin.ID,
		creator,
		time.Now(),
	)
	err = s.repository.InsertPurchaseReceiving(&newItem)
	if err != nil {
		return err
	}

	for _, val := range item.Details {
		newDetail := NewPurchaseReceivingDetail(
			newItem.ID,
			val.ProductId,
			val.Qty,
			val.Price,
			newItem.CreatedBy,
			newItem.CreatedAt,
		)
		err = s.repositoryDetail.InsertPurchaseReceivingDetail(&newDetail)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *service) UpdatePurchaseReceiving(id string, item *UpdatePurchaseReceivingSpec, modifier string) error {
	err := validator.GetValidator().Struct(item)
	if err != nil {
		return business.ErrInvalidSpec
	}

	// admin, err := s.adminService.FindAdminByID(modifier)
	// if err != nil {
	// 	return business.ErrNotHavePermission
	// }

	purchase, err := s.repository.GetPurchaseReceivingById(id)
	if err != nil {
		return business.ErrNotFound
	}

	updateData := purchase.ModifyPurchaseReceiving(
		item.DateReceived,
		item.ReceivedBy,
		item.Description,
		//admin.ID,
		modifier,
		time.Now(),
	)

	err = s.repository.UpdatePurchaseReceiving(&updateData)
	if err != nil {
		return err
	}

	for _, val := range item.Details {
		if val.ID == "" {
			if !val.IsDelete {
				newDetail := NewPurchaseReceivingDetail(
					purchase.ID,
					val.ProductId,
					val.Qty,
					val.Price,
					purchase.UpdatedBy,
					purchase.UpdatedAt,
				)
				err = s.repositoryDetail.InsertPurchaseReceivingDetail(&newDetail)
				if err != nil {
					return err
				}
			}
		} else {
			detail, err := s.repositoryDetail.GetPurchaseReceivingDetailById(val.ID)
			if err != nil {
				return business.ErrNotFound
			}
			if !val.IsDelete {
				updateData := detail.ModifyPurchaseReceivingDetail(
					val.ProductId,
					val.Qty,
					val.Price,
					purchase.UpdatedBy,
					purchase.UpdatedAt,
				)
				err = s.repositoryDetail.UpdatePurchaseReceivingDetail(&updateData)
				if err != nil {
					return err
				}
			} else {
				deleteData := detail.DeletePurchaseReceivingDetail(
					purchase.UpdatedBy,
					purchase.UpdatedAt,
				)
				err = s.repositoryDetail.DeletePurchaseReceivingDetail(&deleteData)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (s *service) DeletePurchaseReceiving(id string, deleter string) error {
	// admin, err := s.adminService.FindAdminByID(deleter)
	// if err != nil {
	// 	return business.ErrNotHavePermission
	// }

	purchase, err := s.repository.GetPurchaseReceivingById(id)
	if err != nil {
		return business.ErrNotFound
	}
	deleteData := purchase.DeletePurchaseReceiving(
		//admin.ID,
		deleter,
		time.Now(),
	)

	// return s.repository.DeletePurchaseReceiving(&deleteData)

	err = s.repositoryDetail.DeletePurchaseReceivingDetail2(id, deleter)
	if err != nil {
		return err
	}

	return s.repository.DeletePurchaseReceiving(&deleteData)
}
