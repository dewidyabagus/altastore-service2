package product

import (
	"AltaStore/api/common"
	"AltaStore/api/middleware"
	"AltaStore/api/v1/product/request"
	"AltaStore/api/v1/product/response"
	"AltaStore/business/product"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	service product.Service
}

func NewController(service product.Service) *Controller {
	return &Controller{service}
}

func (c *Controller) GetAllProduct(ctx echo.Context) error {
	id := ctx.QueryParam("id")
	if id != "" {
		// if _, err := uuid.Parse(id); err != nil {
		// 	return ctx.JSON(common.BadRequestResponse())
		// }
		return ctx.JSON(common.BadRequestResponse()) // Pemindahan pencarian id, khusus id menggunakan path
	}

	isActive := ctx.QueryParam("isactive")
	categoryName := ctx.QueryParam("categoryname")
	code := ctx.QueryParam("code")
	name := ctx.QueryParam("name")
	product, err := c.service.GetAllProductByParameter(id, isActive, categoryName, code, name)
	if err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	return ctx.JSON(
		common.SuccessResponseWithData(
			response.GetAll(product).Products,
		),
	)
}

func (c *Controller) FindProductById(ctx echo.Context) error {
	id := ctx.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	product, err := c.service.FindProductById(id)
	if err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	response := response.GetById(*product)

	return ctx.JSON(common.SuccessResponseWithData(response))
}

func (c *Controller) InsertProduct(ctx echo.Context) error {
	var err error

	adminId, err := middleware.ExtractTokenUser(ctx)
	if err != nil {
		return ctx.JSON(common.UnAuthorizedResponse())
	}

	isAdmin, err := middleware.ExtractTokenRule(ctx)
	if err != nil || !isAdmin {
		return ctx.JSON(common.UnAuthorizedResponse())
	}

	insertProduct := new(request.InsertProductRequest)

	if err = ctx.Bind(insertProduct); err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	if err = c.service.InsertProduct(insertProduct.ToProductSpec(), adminId); err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	return ctx.JSON(common.SuccessResponseWithoutData())
}

func (c *Controller) UpdateProduct(ctx echo.Context) error {
	var err error

	id := ctx.Param("id")
	if _, err = uuid.Parse(id); err != nil {
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

	updateProduct := new(request.UpdateProductRequest)
	if err = ctx.Bind(updateProduct); err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	if err = c.service.UpdateProduct(id, updateProduct.ToProductSpec(), adminId); err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	return ctx.JSON(common.SuccessResponseWithoutData())
}

func (c *Controller) DeleteProduct(ctx echo.Context) error {
	var err error

	id := ctx.Param("id")
	if _, err = uuid.Parse(id); err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	adminId, _ := middleware.ExtractTokenUser(ctx)
	if _, err = uuid.Parse(adminId); err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	isAdmin, err := middleware.ExtractTokenRule(ctx)
	if err != nil || !isAdmin {
		return ctx.JSON(common.UnAuthorizedResponse())
	}

	if err = c.service.DeleteProduct(id, adminId); err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	return ctx.JSON(common.SuccessResponseWithoutData())
}
