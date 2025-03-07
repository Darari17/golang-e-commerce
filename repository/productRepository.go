package repository

import (
	"errors"

	"github.com/Darari17/golang-e-commerce/model"
	"gorm.io/gorm"
)

type IProductRepository interface {
	CreateProduct(product *model.Product) (*model.Product, error)
	UpdateProduct(product *model.Product) (*model.Product, error)
	UpdateStockTx(tx *gorm.DB, productId uint, quantity int) error
	FindProductById(productId uint) (*model.Product, error)
	FindProductByIdTx(tx *gorm.DB, productId uint) (*model.Product, error)
	FindAllProducts() ([]*model.Product, error)
	DeleteProduct(productId uint) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) IProductRepository {
	return &productRepository{
		db: db,
	}
}

// CreateProduct implements IProductRepository.
func (p *productRepository) CreateProduct(product *model.Product) (*model.Product, error) {
	err := p.db.Create(product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

// DeleteProduct implements IProductRepository.
func (p *productRepository) DeleteProduct(productId uint) error {
	return p.db.Where("id = ?", productId).Delete(&model.Product{}).Error
}

// FindAllProducts implements IProductRepository.
func (p *productRepository) FindAllProducts() ([]*model.Product, error) {
	var products []*model.Product
	err := p.db.Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

// FindProductById implements IProductRepository.
func (p *productRepository) FindProductById(productId uint) (*model.Product, error) {
	var product model.Product
	err := p.db.Where("id = ?", productId).First(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

// UpdateProduct implements IProductRepository.
func (p *productRepository) UpdateProduct(product *model.Product) (*model.Product, error) {
	err := p.db.Save(product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

// FindProductByIdTx implements IProductRepository.
func (p *productRepository) FindProductByIdTx(tx *gorm.DB, productId uint) (*model.Product, error) {
	if tx == nil {
		tx = p.db
	}

	var product model.Product
	err := tx.Where("id = ?", productId).First(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

// UpdateStockTx implements IProductRepository.
func (p *productRepository) UpdateStockTx(tx *gorm.DB, productId uint, quantity int) error {
	if tx == nil {
		tx = p.db
	}

	result := tx.Model(&model.Product{}).
		Where("id = ? AND stock >= ?", productId, quantity).
		UpdateColumn("stock", gorm.Expr("stock - ?", quantity))

	if result.RowsAffected == 0 {
		return errors.New("insufficient stock or product not found")
	}

	if result.Error != nil {
		return result.Error
	}

	return nil
}
