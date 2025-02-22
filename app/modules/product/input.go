package product

import "mime/multipart"

type CreateProductInput struct {
	Name        string `form:"name" binding:"required"`
	Description string `form:"description" binding:"required"`
	Price       int    `form:"price" binding:"required"`
}

type UpdateProductInput struct {
	Name        string                `form:"name" binding:"required"`
	Description string                `form:"description" binding:"required"`
	Price       int                   `form:"price" binding:"required"`
	Image       *multipart.FileHeader `form:"image"`
}

type GetProductDetailInput struct {
	ID string `uri:"id" binding:"required"`
}
