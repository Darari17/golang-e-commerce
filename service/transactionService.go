package service

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Darari17/golang-e-commerce/dto"
	"github.com/Darari17/golang-e-commerce/model"
	"github.com/Darari17/golang-e-commerce/repository"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type ITransactionService interface {
	CreateTransaction(input dto.CreateOrder) (dto.OrderResponse, error)
	FindAllTransactionsByUserId(userId uint) ([]dto.OrderResponse, error)
	FindTransactionById(orderId uint) (dto.OrderResponse, error)
	CancelTransaction(orderId uint) error
	UpdateTransactionStatus(orderId uint, status string) error
	DeleteTransaction(orderId uint) error
}

type transactionService struct {
	orderRepository     repository.IOrderRepository
	orderItemRepository repository.IOrderItemRepository
	productRepository   repository.IProductRepository
	db                  *gorm.DB
}

func NewTransactionService(orderRepository repository.IOrderRepository, orderItemRepository repository.IOrderItemRepository, productRepository repository.IProductRepository, db *gorm.DB) ITransactionService {
	return &transactionService{
		orderRepository:     orderRepository,
		orderItemRepository: orderItemRepository,
		productRepository:   productRepository,
		db:                  db,
	}
}

// CancelTransaction implements ITransactionService.
func (t *transactionService) CancelTransaction(orderId uint) error {
	err := t.db.Transaction(func(tx *gorm.DB) error {
		order, err := t.orderRepository.FindOrderByIdTx(tx, orderId)
		if err != nil {
			return err
		}

		if !strings.EqualFold(order.Status, "pending") {
			return errors.New("order cannot be canceled")
		}

		items, err := t.orderItemRepository.FindOrderItemByOrderId(tx, order.ID)
		if err != nil {
			return err
		}

		for _, item := range items {
			err := t.productRepository.UpdateStockTx(tx, item.ProductID, item.Quantity)
			if err != nil {
				return err
			}
		}

		err = t.orderRepository.UpdateOrderStatusTx(tx, order.ID, "canceled")
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

// CreateTransaction implements ITransactionService.
func (t *transactionService) CreateTransaction(input dto.CreateOrder) (dto.OrderResponse, error) {
	validate := validator.New()
	if err := validate.Struct(input); err != nil {
		return dto.OrderResponse{}, err
	}

	var response dto.OrderResponse

	err := t.db.Transaction(func(tx *gorm.DB) error {
		order, err := t.orderRepository.CreateOrderTx(tx, &model.Order{
			UserID: input.UserID,
			Status: "pending",
		})
		if err != nil {
			return err
		}

		var totalPrice float64
		var orderItems []*model.OrderItem

		for _, item := range input.Items {
			product, err := t.productRepository.FindProductByIdTx(tx, item.ProductID)
			if err != nil {
				return err
			}

			if product.Stock < uint(item.Quantity) {
				return errors.New("insufficient stock for product")
			}

			subTotal := float64(item.Quantity) * product.Price
			totalPrice += subTotal

			orderItems = append(orderItems, &model.OrderItem{
				OrderID:   order.ID,
				ProductID: item.ProductID,
				Quantity:  item.Quantity,
				Subtotal:  subTotal,
			})

			err = t.productRepository.UpdateStockTx(tx, item.ProductID, item.Quantity)
			if err != nil {
				return err
			}
		}

		err = t.orderItemRepository.CreateBatchOrderItems(tx, orderItems)
		if err != nil {
			return err
		}

		order.TotalPrice = totalPrice

		_, err = t.orderRepository.UpdateOrderTx(tx, order)
		if err != nil {
			return err
		}

		response = dto.OrderResponse{
			ID:         order.ID,
			UserID:     order.UserID,
			TotalPrice: order.TotalPrice,
			Status:     order.Status,
			CreatedAt:  order.CreatedAt,
			Items:      input.Items,
		}

		return nil
	})

	if err != nil {
		return dto.OrderResponse{}, err
	}

	return response, nil
}

// DeleteTransaction implements ITransactionService.
func (t *transactionService) DeleteTransaction(orderId uint) error {
	err := t.db.Transaction(func(tx *gorm.DB) error {
		order, err := t.orderRepository.FindOrderByIdTx(tx, orderId)
		if err != nil {
			return fmt.Errorf("order id %d not found", orderId)
		}

		if strings.ToLower(order.Status) != "pending" {
			return fmt.Errorf("order id %d cannot be delete", orderId)
		}

		items, err := t.orderItemRepository.FindOrderItemByOrderId(tx, order.ID)
		if err != nil {
			return err
		}

		if len(items) > 0 {
			if err := t.orderItemRepository.DeleteOrderItemById(tx, order.ID); err != nil {
				return err
			}
		}

		err = t.orderRepository.DeleteOrderTx(tx, orderId)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

// FindAllTransactionsByUserId implements ITransactionService.
func (t *transactionService) FindAllTransactionsByUserId(userId uint) ([]dto.OrderResponse, error) {
	var listOfOrders []dto.OrderResponse

	err := t.db.Transaction(func(tx *gorm.DB) error {
		orders, err := t.orderRepository.FindUserOrdersTx(tx, userId)
		if err != nil {
			return err
		}

		if len(orders) == 0 {
			listOfOrders = []dto.OrderResponse{}
			return nil
		}

		listOfOrders = make([]dto.OrderResponse, 0, len(orders))

		for _, order := range orders {
			orderItems, err := t.orderItemRepository.FindOrderItemByOrderId(tx, order.ID)
			if err != nil {
				return err
			}

			listOfItems := make([]dto.OrderItemDTO, 0, len(orderItems))
			for _, item := range orderItems {
				listOfItems = append(listOfItems, dto.OrderItemDTO{
					ProductID: item.ProductID,
					Quantity:  item.Quantity,
					Subtotal:  item.Subtotal,
				})
			}

			listOfOrders = append(listOfOrders, dto.OrderResponse{
				ID:         order.ID,
				UserID:     order.UserID,
				TotalPrice: order.TotalPrice,
				Status:     order.Status,
				CreatedAt:  order.CreatedAt,
				Items:      listOfItems,
			})
		}
		return nil
	})

	if err != nil {
		return []dto.OrderResponse{}, err
	}

	return listOfOrders, nil
}

// FindTransactionById implements ITransactionService.
func (t *transactionService) FindTransactionById(orderId uint) (dto.OrderResponse, error) {
	var response dto.OrderResponse

	err := t.db.Transaction(func(tx *gorm.DB) error {
		order, err := t.orderRepository.FindOrderByIdTx(tx, orderId)
		if err != nil {
			return err
		}

		items, err := t.orderItemRepository.FindOrderItemByOrderId(tx, order.ID)
		if err != nil {
			return err
		}

		orderItems := make([]dto.OrderItemDTO, 0, len(items))
		for _, item := range items {
			orderItems = append(orderItems, dto.OrderItemDTO{
				ProductID: item.ProductID,
				Quantity:  item.Quantity,
				Subtotal:  item.Subtotal,
			})
		}

		response = dto.OrderResponse{
			ID:         order.ID,
			UserID:     order.UserID,
			TotalPrice: order.TotalPrice,
			Status:     order.Status,
			CreatedAt:  order.CreatedAt,
			Items:      orderItems,
		}

		return nil
	})

	if err != nil {
		return dto.OrderResponse{}, err
	}

	return response, nil
}

// UpdateTransactionStatus implements ITransactionService.
func (t *transactionService) UpdateTransactionStatus(orderId uint, status string) error {
	status = strings.ToLower(status)

	validStatus := map[string]bool{
		"canceled":  true,
		"pending":   true,
		"completed": true,
	}

	if !validStatus[status] {
		return errors.New("invalid order status")
	}

	return t.db.Transaction(func(tx *gorm.DB) error {
		order, err := t.orderRepository.FindOrderByIdTx(tx, orderId)
		if err != nil {
			return err
		}

		return t.orderRepository.UpdateOrderStatusTx(tx, order.ID, status)
	})
}
