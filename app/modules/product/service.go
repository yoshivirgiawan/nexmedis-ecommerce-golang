package product

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"time"
)

type Service interface {
	GetProductsWithFilters(take int, skip int, search string) ([]Product, error)
	GetProductByID(ID string) (Product, error)
	CreateProduct(input CreateProductInput, fileHeader *multipart.FileHeader) (Product, error)
	UpdateProduct(inputID GetProductDetailInput, input UpdateProductInput) (Product, error)
	DeleteProduct(ID string) error
}

type service struct {
	repository Repository
}

func NewService() *service {
	repository := NewRepository()
	return &service{repository}
}

func saveUploadedFile(fileHeader *multipart.FileHeader) (string, error) {
	// Generate unique filename
	fileName := fmt.Sprintf("product-%d-%s", time.Now().UnixNano(), fileHeader.Filename)

	// Define storage path
	filePath := fmt.Sprintf("storage/public/%s", fileName)

	// Create storage directory if not exists
	err := os.MkdirAll("storage/public", os.ModePerm)
	if err != nil {
		return "", err
	}

	// Open the uploaded file
	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Create a destination file
	dst, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	// Copy the content of the uploaded file to the destination file
	_, err = io.Copy(dst, file)
	if err != nil {
		return "", err
	}

	return fileName, nil
}

func (s *service) GetProductByID(ID string) (Product, error) {
	product, err := s.repository.FindByID(ID)
	if err != nil {
		return product, err
	}

	if product.ID == "" {
		return product, errors.New("not found")
	}

	return product, nil
}

func (s *service) GetProductsWithFilters(take int, skip int, search string) ([]Product, error) {
	return s.repository.FindAllWithFilters(take, skip, search)
}

func (s *service) CreateProduct(input CreateProductInput, fileHeader *multipart.FileHeader) (Product, error) {
	product := Product{}
	product.Name = input.Name
	product.Description = input.Description
	product.Price = float64(input.Price)

	if fileHeader != nil {
		fileName, err := saveUploadedFile(fileHeader)
		if err != nil {
			return product, err
		}
		product.Image = fileName
	}

	newProduct, err := s.repository.Save(product)
	if err != nil {
		return newProduct, err
	}

	return newProduct, nil
}

func (s *service) UpdateProduct(inputID GetProductDetailInput, input UpdateProductInput) (Product, error) {
	product, err := s.repository.FindByID(inputID.ID)
	if err != nil {
		return product, err
	}
	product.Name = input.Name
	product.Description = input.Description
	product.Price = float64(input.Price)

	if input.Image != nil {
		fileName, err := saveUploadedFile(input.Image)
		if err != nil {
			return product, err
		}
		product.Image = fileName
	}

	updatedProduct, err := s.repository.Update(product)
	if err != nil {
		return updatedProduct, err
	}

	return updatedProduct, nil
}

func (s *service) DeleteProduct(ID string) error {
	product, err := s.repository.FindByID(ID)
	if err != nil {
		return err
	}

	if product.ID == "" {
		return errors.New("not found")
	}

	_, err = s.repository.Delete(ID)
	if err != nil {
		return err
	}

	return nil
}
