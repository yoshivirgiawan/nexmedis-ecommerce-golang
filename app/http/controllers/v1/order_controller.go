package v1

import (
	"ecommerce/app/modules/order"
	"ecommerce/app/modules/user"
	"ecommerce/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type orderController struct {
	orderService order.Service
}

func NewOrderController() *orderController {
	orderService := order.NewService()
	return &orderController{orderService}
}

func (h *orderController) GetOrders(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)
	orders, err := h.orderService.GetOrders(currentUser.ID)
	if err != nil {
		response := helper.APIResponse(err.Error(), http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of orders", http.StatusOK, "success", order.FormatOrders(orders))
	c.JSON(http.StatusOK, response)
}

func (h *orderController) GetOrder(c *gin.Context) {
	var input order.GetOrderInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Order detail failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser
	orderDetail, err := h.orderService.GetOrderByID(input)
	if err != nil {
		response := helper.APIResponse(err.Error(), http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Order detail", http.StatusOK, "success", order.FormatOrder(orderDetail))
	c.JSON(http.StatusOK, response)
}

func (h *orderController) Create(c *gin.Context) {
	var input order.CreateOrderInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Order creation failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser
	newOrder, err := h.orderService.CreateOrder(input)
	if err != nil {
		response := helper.APIResponse(err.Error(), http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Order successfully created", http.StatusOK, "success", order.FormatOrder(newOrder))
	c.JSON(http.StatusOK, response)
}

func (h *orderController) PayOrder(c *gin.Context) {
	var input order.GetOrderInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Order payment failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser
	paidOrder, err := h.orderService.PayOrder(input)
	if err != nil {
		response := helper.APIResponse(err.Error(), http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Order successfully paid", http.StatusOK, "success", order.FormatOrder(paidOrder))
	c.JSON(http.StatusOK, response)
}
