package v1

import (
	"ecommerce/app/modules/cart"
	"ecommerce/app/modules/user"
	"ecommerce/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type cartController struct {
	cartService cart.Service
}

func NewCartController() *cartController {
	cartService := cart.NewService()
	return &cartController{cartService}
}

func (h *cartController) GetCarts(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)
	carts, err := h.cartService.GetCarts(currentUser.ID)
	if err != nil {
		response := helper.APIResponse(err.Error(), http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of carts", http.StatusOK, "success", cart.FormatCarts(carts))
	c.JSON(http.StatusOK, response)
}

func (h *cartController) GetCart(c *gin.Context) {
	var input cart.GetCartDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Cart detail failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser
	cartDetail, err := h.cartService.GetCartByID(input)
	if err != nil {
		response := helper.APIResponse(err.Error(), http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Cart detail", http.StatusOK, "success", cart.FormatCart(cartDetail))
	c.JSON(http.StatusOK, response)
}

func (h *cartController) AddToCart(c *gin.Context) {
	var input cart.CreateCartInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Cart creation failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser
	newCart, err := h.cartService.AddToCart(input)
	if err != nil {
		response := helper.APIResponse(err.Error(), http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Cart successfully created", http.StatusOK, "success", cart.FormatCart(newCart))
	c.JSON(http.StatusOK, response)
}

func (h *cartController) Update(c *gin.Context) {
	var inputDetail cart.GetCartDetailInput

	err := c.ShouldBindUri(&inputDetail)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Cart update failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	inputDetail.User = c.MustGet("currentUser").(user.User)
	var inputData cart.UpdateCartInput

	err = c.ShouldBind(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Cart update failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updatedCart, err := h.cartService.UpdateCart(inputDetail, inputData)
	if err != nil {
		response := helper.APIResponse(err.Error(), http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Cart successfully updated", http.StatusOK, "success", cart.FormatCart(updatedCart))
	c.JSON(http.StatusOK, response)
}

func (h *cartController) Delete(c *gin.Context) {
	var input cart.GetCartDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Cart delete failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	input.User = c.MustGet("currentUser").(user.User)
	err = h.cartService.DeleteCart(input)
	if err != nil {
		response := helper.APIResponse(err.Error(), http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Cart successfully deleted", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}
