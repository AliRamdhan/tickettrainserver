package service

import (
	"time"

	"github.com/AliRamdhan/trainticket/config"
	"github.com/AliRamdhan/trainticket/internal/model"
	"github.com/google/uuid"
)

type OrderTicketService struct{}

func NewOrderTicketService() *OrderTicketService {
	return &OrderTicketService{}
}

func (os *OrderTicketService) CreateOrder(order *model.Order, orderSeatId uint,
	orderUserId uint, ticketId uint) error {
	order.OrderNumber = uuid.New()
	order.OrderTicketStatus = "Pending"
	order.OrderSeatId = orderSeatId
	order.OrderUserId = orderUserId
	order.OrderTicketId = ticketId
	order.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	return config.DB.Create(order).Error
}

func (os *OrderTicketService) GetAllOrderByUser(userId uint) ([]model.Order, error) {
	var orderUser []model.Order
	if err := config.DB.Where("order_user_id = ?", userId).Find(&orderUser).Error; err != nil {
		return nil, err
	}
	return orderUser, nil
}

func (os *OrderTicketService) GetAllOrder() ([]model.Order, error) {
	var tickets []model.Order
	if err := config.DB.Find(&tickets).Error; err != nil {
		return nil, err
	}
	return tickets, nil
}
