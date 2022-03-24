package controller

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/pgeowng/wb-l0/service"
)

type RestController struct {
	srv service.OrderService
	log *log.Logger
}

func NewRest(ctx context.Context, srv service.OrderService, log *log.Logger) *RestController {

	return &RestController{
		srv: srv,
		log: log,
	}
}

func (ctl *RestController) GetOrder(c *fiber.Ctx) error {
	id := c.Params("id")
	ctx := c.UserContext()

	result, err := ctl.srv.GetOrder(ctx, id)
	if err != nil {
		c.SendStatus(404)
		return c.JSON(map[string]interface{}{
			"error": err.Error(),
		})
	}

	// c.Context().SetContentType("application/json")
	return c.Send(result)
}

func (ctl *RestController) GetIds(c *fiber.Ctx) error {
	ctx := c.UserContext()
	result, err := ctl.srv.GetIds(ctx)
	if err != nil {
		ctl.log.Print("err", err)
		c.SendStatus(500)
		return c.JSON(map[string]interface{}{
			"error": err.Error(),
		})
	}

	ctl.log.Print(result)

	return c.JSON(result)
}
