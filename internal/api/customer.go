package api

import(
	"context"
	"time"
	"github.com/gofiber/fiber/v2"
	"shellrean.id/Go-RestAPI/domain"
	"shellrean.id/Go-RestAPI/dto"
)
type customerApi struct{
	customerService domain.CustomerService
}

func NewCustomer(app *fiber.App, customerService domain.CustomerService){
	ca := customerApi{
		customerService: customerService,
	}

	app.Get("/customers", ca.Index)
}

func (ca customerApi) Index(ctx *fiber.Ctx) error{
	c, cancel := context.WithTimeout(ctx.Context(), 10 * time.Second)
	defer cancel()
	res, err:=ca.customerService.Index(c)

	if err != nil{
		return ctx.Status(500).JSON(dto.CreateResponseError(err.Error()))
	}
	return ctx.JSON(dto.CreateResponseSuccess(res))
}
