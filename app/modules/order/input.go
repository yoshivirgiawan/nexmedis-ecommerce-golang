package order

import "ecommerce/app/modules/user"

type CreateOrderInput struct {
	CartIDs []string `json:"cart_ids" binding:"required"`
	User    user.User
}

type GetOrderInput struct {
	ID   string `uri:"id" binding:"required"`
	User user.User
}
