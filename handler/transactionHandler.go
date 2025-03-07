package handler

import (
	"net/http"
	"strconv"

	"github.com/Darari17/golang-e-commerce/dto"
	"github.com/Darari17/golang-e-commerce/model"
	"github.com/Darari17/golang-e-commerce/service"
	"github.com/gin-gonic/gin"
)

type ITransactionHandler interface {
	CreateTransaction(c *gin.Context)
	FindAllTransactionsByUserId(c *gin.Context)
	FindTransactionById(c *gin.Context)
	CancelTransaction(c *gin.Context)
	UpdateTransactionStatus(c *gin.Context)
	DeleteTransaction(c *gin.Context)
}

type transactionHandler struct {
	txService service.ITransactionService
}

func NewTxHandler(txService service.ITransactionService) ITransactionHandler {
	return &transactionHandler{
		txService: txService,
	}
}

// CancelTransaction implements ITransactionHandler.
func (t *transactionHandler) CancelTransaction(c *gin.Context) {
	txId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction ID"})
		return
	}

	err = t.txService.CancelTransaction(uint(txId))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "Transaction canceled successfully",
	})
}

// CreateTransaction implements ITransactionHandler.
func (t *transactionHandler) CreateTransaction(c *gin.Context) {
	var payload dto.CreateOrder
	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order, err := t.txService.CreateTransaction(payload)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "Created",
		"data":   order,
	})
}

// DeleteTransaction implements ITransactionHandler.
func (t *transactionHandler) DeleteTransaction(c *gin.Context) {
	txId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction ID"})
		return
	}

	err = t.txService.DeleteTransaction(uint(txId))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "Transaction deleted successfully",
	})
}

// FindAllTransactionsByUserId implements ITransactionHandler.
func (t *transactionHandler) FindAllTransactionsByUserId(c *gin.Context) {
	userData, exists := c.Get("user")
	if !exists {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	user, ok := userData.(model.User)
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user data"})
		return
	}

	transactions, err := t.txService.FindAllTransactionsByUserId(user.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":       "OK",
		"transactions": transactions,
	})
}

// FindTransactionById implements ITransactionHandler.
func (t *transactionHandler) FindTransactionById(c *gin.Context) {
	txId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction ID"})
		return
	}

	transaction, err := t.txService.FindTransactionById(uint(txId))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":      "OK",
		"transaction": transaction,
	})
}

// UpdateTransactionStatus implements ITransactionHandler.
func (t *transactionHandler) UpdateTransactionStatus(c *gin.Context) {
	txId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction ID"})
		return
	}

	var status dto.UpdateStatusRequest
	err = c.ShouldBindJSON(&status)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = t.txService.UpdateTransactionStatus(uint(txId), status.Status)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "Transaction status updated",
	})
}
