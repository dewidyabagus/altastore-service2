package purchasereceiving

import (
	"AltaStore/api/common"
	"AltaStore/api/middleware"
	"AltaStore/api/v1/purchasereceiving/request"
	"AltaStore/api/v1/purchasereceiving/response"
	"AltaStore/business/purchasereceiving"

	"github.com/google/uuid"
	echo "github.com/labstack/echo/v4"
)

type Controller struct {
	service purchasereceiving.Service
}

func NewController(service purchasereceiving.Service) *Controller {
	return &Controller{service}
}

func (c *Controller) InsertPurchaseReceiving(ctx echo.Context) error {
	var err error
	adminId, err := middleware.ExtractTokenUser(ctx)
	if err != nil {
		return ctx.JSON(common.UnAuthorizedResponse())
	}
	isAdmin, err := middleware.ExtractTokenRule(ctx)
	if err != nil || !isAdmin {
		return ctx.JSON(common.UnAuthorizedResponse())
	}
	if _, err := uuid.Parse(adminId); err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	insertData := new(request.InsertPurchaseReceivingRequest)

	if err = ctx.Bind(insertData); err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	if err = c.service.InsertPurchaseReceiving(insertData.ToPurchaseReceivingSpec(), adminId); err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	return ctx.JSON(common.SuccessResponseWithoutData())
}

func (c *Controller) UpdatePurchaseReceiving(ctx echo.Context) error {
	var err error
	id := ctx.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	adminId, err := middleware.ExtractTokenUser(ctx)
	if err != nil {
		return ctx.JSON(common.UnAuthorizedResponse())
	}
	isAdmin, err := middleware.ExtractTokenRule(ctx)
	if err != nil || !isAdmin {
		return ctx.JSON(common.UnAuthorizedResponse())
	}
	if _, err := uuid.Parse(adminId); err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	updateData := new(request.UpdatePurchaseReceivingRequest)
	if err = ctx.Bind(updateData); err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	if err = c.service.UpdatePurchaseReceiving(id, updateData.ToPurchaseReceivingSpec(), adminId); err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	return ctx.JSON(common.SuccessResponseWithoutData())
}

func (c *Controller) DeletePurchaseReceiving(ctx echo.Context) error {
	var err error

	id := ctx.Param("id")
	adminId, err := middleware.ExtractTokenUser(ctx)
	if err != nil {
		return ctx.JSON(common.UnAuthorizedResponse())
	}
	isAdmin, err := middleware.ExtractTokenRule(ctx)
	if err != nil || !isAdmin {
		return ctx.JSON(common.UnAuthorizedResponse())
	}
	if _, err = uuid.Parse(id); err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}
	if _, err = uuid.Parse(adminId); err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	if err = c.service.DeletePurchaseReceiving(id, adminId); err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	return ctx.JSON(common.SuccessResponseWithoutData())
}

func (c *Controller) GetAllPurchaseReceiving(ctx echo.Context) error {
	id := ctx.QueryParam("id")
	if id != "" {
		if _, err := uuid.Parse(id); err != nil {
			return ctx.JSON(common.BadRequestResponse())
		}
	}
	code := ctx.QueryParam("code")

	adminId, err := middleware.ExtractTokenUser(ctx)
	if err != nil {
		return ctx.JSON(common.UnAuthorizedResponse())
	}
	isAdmin, err := middleware.ExtractTokenRule(ctx)
	if err != nil || !isAdmin {
		return ctx.JSON(common.UnAuthorizedResponse())
	}
	if _, err := uuid.Parse(adminId); err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	data, err := c.service.GetAllPurchaseReceivingByParameter(code, adminId)
	if err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	return ctx.JSON(
		common.SuccessResponseWithData(
			response.GetAll(data).PurchaseReceivings,
		),
	)
}

func (c *Controller) FindPurchaseReceivingById(ctx echo.Context) error {
	id := ctx.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}
	adminId, err := middleware.ExtractTokenUser(ctx)
	if err != nil {
		return ctx.JSON(common.UnAuthorizedResponse())
	}
	isAdmin, err := middleware.ExtractTokenRule(ctx)
	if err != nil || !isAdmin {
		return ctx.JSON(common.UnAuthorizedResponse())
	}
	if _, err := uuid.Parse(adminId); err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}
	data, err := c.service.GetPurchaseReceivingById(id, adminId)
	if err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	response := response.GetById(*data)

	return ctx.JSON(common.SuccessResponseWithData(response))
}
