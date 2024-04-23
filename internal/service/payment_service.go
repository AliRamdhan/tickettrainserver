package service

import (
	"strconv"
	"time"

	"github.com/AliRamdhan/trainticket/config"
	"github.com/AliRamdhan/trainticket/internal/model"
	"github.com/google/uuid"
)

type PaymentService struct{}

func NewPaymentService() *PaymentService {
	return &PaymentService{}
}

func (ps *PaymentService) CreatePayment(payment *model.Payment, orderId uint) error {
	order := &model.Order{}
	if err := config.DB.Preload("Ticket").First(order, orderId).Error; err != nil {
		return err // Handle error
	}

	// Extract ticket price
	ticketPrice, err := strconv.ParseUint(order.Ticket.TicketPrice, 10, 64)
	if err != nil {
		return err // Handle error
	}
	ticketPricePPh := ticketPrice + uint64(float64(ticketPrice)*0.1)
	// Set payment details
	payment.PaymentCode = uuid.New()
	payment.PaymentOrderId = orderId
	payment.PaymentStatus = "Not Complete"
	payment.PaymentTotal = uint(ticketPricePPh)
	payment.CreatedAt = time.Now().Format("2006-01-02 15:04:05")

	// Create the payment record
	if err := config.DB.Create(payment).Error; err != nil {
		return err
	}

	if err := config.DB.Model(&order).Where("order_id = ?", orderId).Update("order_ticket_status", "active payment").Error; err != nil {
		return err
	}

	return nil
}
