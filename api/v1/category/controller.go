package category

import (
	"AltaStore/api/common"
	"AltaStore/api/middleware"
	"AltaStore/api/v1/category/request"
	"AltaStore/api/v1/category/response"
	"AltaStore/business/category"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	service category.Service
}

func NewController(service category.Service) *Controller {
	return &Controller{service}
}

func (c *Controller) GetAllCategory(ctx echo.Context) error {
	categories, err := c.service.GetAllCategory()
	if err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	return ctx.JSON(
		common.SuccessResponseWithData(
			response.GetAllCategory(categories).Categories,
		),
	)
}

func (c *Controller) FindCategoryById(ctx echo.Context) error {
	id := ctx.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	category, err := c.service.FindCategoryById(id)
	if err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	response := response.GetOneCategory(*category)

	return ctx.JSON(common.SuccessResponseWithData(response))
}

func (c *Controller) InsertCategory(ctx echo.Context) error {
	var err error

	insertCategory := new(request.InsertCategoryRequest)
	adminId, err := middleware.ExtractTokenUser(ctx)
	if err != nil {
		return ctx.JSON(common.UnAuthorizedResponse())
	}
	isAdmin, err := middleware.ExtractTokenRule(ctx)
	if err != nil || !isAdmin {
		return ctx.JSON(common.UnAuthorizedResponse())
	}
	if err = ctx.Bind(insertCategory); err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	if err = c.service.InsertCategory(insertCategory.ToCategorySpec(), adminId); err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	return ctx.JSON(common.SuccessResponseWithoutData())
}

func (c *Controller) UpdateCategory(ctx echo.Context) error {
	var err error

	id := ctx.Param("id")
	if _, err = uuid.Parse(id); err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	adminId, err := middleware.ExtractTokenUser(ctx)
	if err != nil {
		return ctx.JSON(common.UnAuthorizedResponse())
	}
	if _, err = uuid.Parse(adminId); err != nil {
		return ctx.JSON(common.UnAuthorizedResponse())
	}

	isAdmin, err := middleware.ExtractTokenRule(ctx)
	if err != nil || !isAdmin {
		return ctx.JSON(common.UnAuthorizedResponse())
	}

	updateCategory := new(request.UpdateCategoryRequest)
	if err = ctx.Bind(updateCategory); err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	if err = c.service.UpdateCategory(id, updateCategory.ToCategory(), adminId); err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	return ctx.JSON(common.SuccessResponseWithoutData())
}

func (c *Controller) DeleteCategory(ctx echo.Context) error {
	var err error

	id := ctx.Param("id")
	if _, err = uuid.Parse(id); err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	isAdmin, err := middleware.ExtractTokenRule(ctx)
	if err != nil || !isAdmin {
		return ctx.JSON(common.UnAuthorizedResponse())
	}

	adminId, _ := middleware.ExtractTokenUser(ctx)
	if _, err = uuid.Parse(adminId); err != nil {
		return ctx.JSON(common.UnAuthorizedResponse())
	}

	if err = c.service.DeleteCategory(id, adminId); err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	return ctx.JSON(common.SuccessResponseWithoutData())
}
