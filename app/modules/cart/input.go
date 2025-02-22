package cart

import "ecommerce/app/modules/user"

type CreateCartInput struct {
	ProductID string `json:"product_id" binding:"required"`
	Quantity  int    `json:"quantity" binding:"required"`
	User      user.User
}

type UpdateCartInput struct {
	Quantity int `json:"quantity" binding:"required"`
}

type GetCartDetailInput struct {
	ID   string `uri:"id" binding:"required"`
	User user.User
}
