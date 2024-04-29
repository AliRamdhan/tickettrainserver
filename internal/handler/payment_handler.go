package handler

import (
	"net/http"

	"github.com/AliRamdhan/trainticket/internal/model"
	"github.com/AliRamdhan/trainticket/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PaymentHandler struct {
	paymentService *service.PaymentService
}

func NewPaymentHandler(ps *service.PaymentService) *PaymentHandler {
	return &PaymentHandler{paymentService: ps}
}

func (ph *PaymentHandler) GetAllPayments(c *gin.Context) {
	payments, err := ph.paymentService.GetAllPayments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "All payments", "Payments": payments})
}

func (ph *PaymentHandler) CreatePayment(c *gin.Context) {
	var payment model.Payment
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	paymentOrderId := payment.PaymentOrderId
	if paymentOrderId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing service order ID"})
		return
	}

	if err := ph.paymentService.CreatePayment(&payment, paymentOrderId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Payment created successfully", "payment": payment})
}

func (ph *PaymentHandler) ProcessUserPayment(c *gin.Context) {
	var request struct {
		PaymentCode  uuid.UUID `json:"paymentCode"`
		PaymentTotal uint      `json:"paymentTotal"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ph.paymentService.ProcessUserPayment(request.PaymentCode, request.PaymentTotal); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Payment processed successfully"})
}
