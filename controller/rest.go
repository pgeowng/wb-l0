package controller

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/pgeowng/wb-l0/model"
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
	fmt.Println(id)

	result, err := ctl.srv.GetOrder(ctx, id)
	if err != nil {
		c.SendStatus(404)
		return c.JSON(map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(result)
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

	return c.JSON(result)
}

func (ctl *RestController) IndexPage(c *fiber.Ctx) error {
	idx := c.Params("idx")
	leftBound := 0
	limit := 20
	pageIndex := 0

	if idx64, err := strconv.ParseUint(idx, 0, 32); err == nil {
		leftBound = int(idx64) * limit
		pageIndex = int(idx64)
	}

	ctx := c.UserContext()
	result, err := ctl.srv.GetIds(ctx)
	if err != nil {
		c.SendStatus(500)
		return c.JSON(map[string]interface{}{
			"error": err.Error(),
		})
	}

	size := len(result)
	pageCount := (size + limit - 1) / limit

	if leftBound > size {
		leftBound = size
	}

	rightBound := leftBound + limit
	if rightBound > size {
		rightBound = size
	}

	idList := result[leftBound:rightBound]
	orderList := make([]*model.Order, 0, len(idList))

	for _, id := range idList {
		order, err := ctl.srv.GetOrder(ctx, id)
		if err != nil {
			c.SendStatus(500)
			return c.JSON(map[string]interface{}{
				"error": err.Error(),
			})
		}

		orderList = append(orderList, order)
	}

	return c.Render("index", fiber.Map{
		"totalAmount": size,
		"pageIndex":   pageIndex,
		"pageCount":   pageCount,
		"orderList":   orderList,
		"offset":      leftBound,
	})
}
