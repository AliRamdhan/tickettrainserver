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

// func (os *OrderTicketService) CreateOrder(order *model.Order, orderSeatId uint,
//
//		orderUserId uint, ticketId uint) error {
//		order.OrderNumber = uuid.New()
//		order.OrderTicketStatus = "Pending"
//		order.OrderSeatId = orderSeatId
//		order.OrderUserId = orderUserId
//		order.OrderTicketId = ticketId
//		order.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
//		return config.DB.Create(order).Error
//	}
func (os *OrderTicketService) CreateOrder(order *model.Order, orderSeatId uint, orderUserId uint, ticketId uint) error {
	seat := &model.Seat{}
	if err := config.DB.First(seat, orderSeatId).Error; err != nil {
		return err // Handle error
	}
	order.OrderNumber = uuid.New()
	order.OrderTicketStatus = "Pending"
	order.OrderSeatId = orderSeatId
	order.OrderUserId = orderUserId
	order.OrderTicketId = ticketId
	order.CreatedAt = time.Now().Format("2006-01-02 15:04:05")

	if err := config.DB.Create(&order).Error; err != nil {
		return err
	}

	if err := config.DB.Model(seat).Where("seat_id = ?", orderSeatId).Update("seat_available", false).Error; err != nil {
		return err
	}

	return nil
}

func (os *OrderTicketService) GetAllOrderByUser(userId uint) ([]model.Order, error) {
	var orderUser []model.Order
	if err := config.DB.Preload("Ticket").Preload("User").Preload("Seat").Where("order_user_id = ?", userId).Find(&orderUser).Error; err != nil {
		return nil, err
	}
	return orderUser, nil
}

func (os *OrderTicketService) GetAllOrder() ([]model.Order, error) {
	var tickets []model.Order
	if err := config.DB.Preload("Ticket").Preload("User").Preload("Seat").Find(&tickets).Error; err != nil {
		return nil, err
	}
	return tickets, nil
}
