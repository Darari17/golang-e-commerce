package service

import (
	"github.com/Darari17/golang-e-commerce/dto"
	"github.com/Darari17/golang-e-commerce/helper"
	"github.com/Darari17/golang-e-commerce/model"
	"github.com/Darari17/golang-e-commerce/repository"
	"github.com/go-playground/validator/v10"
)

type IProductService interface {
	CreateProduct(input dto.ProductRequest) (dto.ProductResponse, error)
	UpdateProduct(productId uint, input dto.ProductUpdateRequest) (dto.ProductResponse, error)
	DeleteProduct(productId uint) error
	FindProductById(productId uint) (dto.ProductResponse, error)
	FindAllProducts() ([]dto.ProductResponse, error)
}

type productService struct {
	productRepository repository.IProductRepository
}

func NewProductRepository(productRepository repository.IProductRepository) IProductService {
	return &productService{
		productRepository: productRepository,
	}
}

// CreateProduct implements IProductService.
func (p *productService) CreateProduct(input dto.ProductRequest) (dto.ProductResponse, error) {
	validate := validator.New()
	err := validate.Struct(input)
	if err != nil {
		return dto.ProductResponse{}, err
	}

	createProduct, err := p.productRepository.CreateProduct(&model.Product{
		Name:     input.Name,
		Price:    input.Price,
		Stock:    input.Stock,
		Category: input.Category,
	})
	if err != nil {
		return dto.ProductResponse{}, err
	}

	response := dto.ProductResponse{
		ID:        createProduct.ID,
		Name:      createProduct.Name,
		Price:     createProduct.Price,
		Stock:     createProduct.Stock,
		Category:  createProduct.Category,
		CreatedAt: createProduct.CreatedAt,
	}

	return response, nil
}

// DeleteProduct implements IProductService.
func (p *productService) DeleteProduct(productId uint) error {
	return p.productRepository.DeleteProduct(productId)
}

// FindAllProducts implements IProductService.
func (p *productService) FindAllProducts() ([]dto.ProductResponse, error) {
	products, err := p.productRepository.FindAllProducts()
	if err != nil {
		return nil, err
	}

	var responses []dto.ProductResponse
	for _, product := range products {
		responses = append(responses, dto.ProductResponse{
			ID:        product.ID,
			Name:      product.Name,
			Price:     product.Price,
			Stock:     product.Stock,
			Category:  product.Category,
			CreatedAt: product.CreatedAt,
		})
	}

	return responses, nil
}

// FindProductById implements IProductService.
func (p *productService) FindProductById(productId uint) (dto.ProductResponse, error) {
	product, err := p.productRepository.FindProductById(productId)
	if err != nil {
		return dto.ProductResponse{}, err
	}

	response := dto.ProductResponse{
		ID:        product.ID,
		Name:      product.Name,
		Price:     product.Price,
		Stock:     product.Stock,
		Category:  product.Category,
		CreatedAt: product.CreatedAt,
	}

	return response, nil
}

// UpdateProduct implements IProductService.
func (p *productService) UpdateProduct(productId uint, input dto.ProductUpdateRequest) (dto.ProductResponse, error) {
	validate := validator.New()
	err := validate.Struct(input)
	if err != nil {
		return dto.ProductResponse{}, err
	}

	product, err := p.productRepository.FindProductById(productId)
	if err != nil {
		return dto.ProductResponse{}, err
	}

	request := model.Product{
		ID:       product.ID,
		Name:     helper.IfNotEmpty(product.Name, input.Name),
		Price:    product.Price,
		Stock:    product.Stock,
		Category: helper.IfNotEmpty(product.Category, input.Category),
	}

	if input.Price != nil {
		product.Price = *input.Price
	}

	if input.Stock != nil {
		product.Stock = *input.Stock
	}

	updateProduct, err := p.productRepository.UpdateProduct(&request)
	if err != nil {
		return dto.ProductResponse{}, err
	}

	response := dto.ProductResponse{
		ID:        updateProduct.ID,
		Name:      updateProduct.Name,
		Price:     updateProduct.Price,
		Stock:     updateProduct.Stock,
		Category:  updateProduct.Category,
		CreatedAt: updateProduct.CreatedAt,
	}

	return response, nil
}
