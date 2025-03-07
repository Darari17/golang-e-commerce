package repository

import (
	"fmt"

	"github.com/Darari17/golang-e-commerce/model"
	"gorm.io/gorm"
)

type IOrderRepository interface {
	CreateOrderTx(tx *gorm.DB, order *model.Order) (*model.Order, error)
	DeleteOrderTx(tx *gorm.DB, orderId uint) error
	FindOrderByIdTx(tx *gorm.DB, orderId uint) (*model.Order, error)
	FindOrderStatusByIdTx(tx *gorm.DB, orderId uint) (string, error)
	FindUserOrdersTx(tx *gorm.DB, userId uint) ([]*model.Order, error)
	UpdateOrderStatusTx(tx *gorm.DB, orderId uint, status string) error
	UpdateOrderTx(tx *gorm.DB, order *model.Order) (*model.Order, error)
}

type orderRepository struct{}

func NewOrderRepository() IOrderRepository {
	return &orderRepository{}
}

// CreateOrderTx implements IOrderRepository.
func (o *orderRepository) CreateOrderTx(tx *gorm.DB, order *model.Order) (*model.Order, error) {
	err := tx.Create(order).Error
	if err != nil {
		return nil, err
	}
	return order, nil
}

// DeleteOrderTx implements IOrderRepository.
func (o *orderRepository) DeleteOrderTx(tx *gorm.DB, orderId uint) error {
	return tx.Delete(&model.Order{}, orderId).Error
}

// FindOrderByIdTx implements IOrderRepository.
func (o *orderRepository) FindOrderByIdTx(tx *gorm.DB, orderId uint) (*model.Order, error) {
	var order model.Order
	err := tx.Where("id = ?", orderId).First(&order).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

// FindOrderStatusByIdTx implements IOrderRepository.
func (o *orderRepository) FindOrderStatusByIdTx(tx *gorm.DB, orderId uint) (string, error) {
	var status string
	err := tx.Model(&model.Order{}).Select("status").Where("id = ?", orderId).First(&status).Error
	if err != nil {
		return "", err
	}
	return status, nil
}

// FindUserOrdersTx implements IOrderRepository.
func (o *orderRepository) FindUserOrdersTx(tx *gorm.DB, userId uint) ([]*model.Order, error) {
	var orders []*model.Order
	err := tx.Where("user_id = ?", userId).Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}

// UpdateOrderStatusTx implements IOrderRepository.
func (o *orderRepository) UpdateOrderStatusTx(tx *gorm.DB, orderId uint, status string) error {
	result := tx.Model(&model.Order{}).Where("id = ?", orderId).Update("status", status)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no order found with id %d", orderId)
	}
	return nil
}

// UpdateOrderTx implements IOrderRepository.
func (o *orderRepository) UpdateOrderTx(tx *gorm.DB, order *model.Order) (*model.Order, error) {
	result := tx.Model(&model.Order{}).Where("id = ?", order.ID).Updates(order)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("no order found with id %d", order.ID)
	}
	return order, nil
}
