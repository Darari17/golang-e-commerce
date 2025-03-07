package repository

import (
	"github.com/Darari17/golang-e-commerce/model"
	"gorm.io/gorm"
)

type IOrderItemRepository interface {
	CreateSingleOrderItem(tx *gorm.DB, item *model.OrderItem) (*model.OrderItem, error)
	CreateBatchOrderItems(tx *gorm.DB, items []*model.OrderItem) error
	FindOrderItemByOrderId(tx *gorm.DB, orderId uint) ([]*model.OrderItem, error)
	UpdateOrderItem(tx *gorm.DB, item *model.OrderItem) (*model.OrderItem, error)
	DeleteOrderItemById(tx *gorm.DB, id uint) error
}

type orderItemRepository struct{}

func NewOrderItemRepository() IOrderItemRepository {
	return &orderItemRepository{}
}

// CreateBatchOrderItems implements IOrderItemRepository.
func (o *orderItemRepository) CreateBatchOrderItems(tx *gorm.DB, items []*model.OrderItem) error {
	if len(items) == 0 {
		return nil
	}
	return tx.Create(items).Error
}

// CreateSingleOrderItem implements IOrderItemRepository.
func (o *orderItemRepository) CreateSingleOrderItem(tx *gorm.DB, item *model.OrderItem) (*model.OrderItem, error) {
	err := tx.Create(item).Error
	if err != nil {
		return nil, err
	}
	return item, nil
}

// DeleteOrderItem implements IOrderItemRepository.
func (o *orderItemRepository) DeleteOrderItemById(tx *gorm.DB, id uint) error {
	return tx.Where("id = ?", id).Delete(&model.OrderItem{}).Error
}

// FindOrderItemByOrderId implements IOrderItemRepository.
func (o *orderItemRepository) FindOrderItemByOrderId(tx *gorm.DB, orderId uint) ([]*model.OrderItem, error) {
	var orderItems []*model.OrderItem
	err := tx.Find(&orderItems, "order_id = ?", orderId).Error
	if err != nil {
		return nil, nil
	}
	return orderItems, nil
}

// UpdateOrderItem implements IOrderItemRepository.
func (o *orderItemRepository) UpdateOrderItem(tx *gorm.DB, item *model.OrderItem) (*model.OrderItem, error) {
	err := tx.Save(item).Error
	if err != nil {
		return nil, err
	}
	return item, nil
}
