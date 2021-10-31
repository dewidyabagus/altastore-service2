package purchasereceiving

import (
	"AltaStore/business"
	// "AltaStore/business/admin"
	"AltaStore/util/validator"
	"time"

	"github.com/google/uuid"
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
	Code string `validate:"required"`
	// DateReceived time.Time `validate:"required"`
	// ReceivedBy   string    `validate:"required"`
	Description string
	Details     []UpdatePurchaseReceivingDetailSpec `validate:"required"`
}
type UpdatePurchaseReceivingDetailSpec struct {
	// ID        string
	ProductId string `validate:"required"`
	Qty       int    `validate:"required"`
	Price     int    `validate:"required"`
	Status    int    `validate:"required,min=0,max=2"`
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

	if _, err := s.repository.GetPurchaseReceivingById(id); err != nil {
		return business.ErrNotFound
	}

	if err := s.repository.UpdatePurchaseReceiving2(&id, &item.Code, &item.Description, &modifier, time.Now()); err != nil {
		return err
	}

	for _, val := range item.Details {
		if val.Status == 0 {
			// Insert
			err = s.repositoryDetail.InsertPurchaseReceivingDetail(&PurchaseReceivingDetail{
				ID:                  uuid.NewString(),
				PurchaseReceivingId: id,
				ProductId:           val.ProductId,
				Qty:                 int32(val.Qty),
				Price:               int64(val.Price),
				CreatedBy:           modifier,
				CreatedAt:           time.Now(),
				UpdatedBy:           modifier,
				UpdatedAt:           time.Now(),
			})
		} else if val.Status == 1 {
			// Update
			err = s.repositoryDetail.UpdatePurchaseReceivingDetail(
				&id, &val.ProductId, &val.Price, &val.Qty, &modifier, time.Now(),
			)
		} else {
			// Deleted
			err = s.repositoryDetail.DeletePurchaseReceivingDetail(&id, &val.ProductId, &modifier, time.Now())
		}

		if err != nil {
			return err
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
