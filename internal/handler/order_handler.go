package handler

import (
	"fmt"
	"net/http"

	"github.com/AliRamdhan/trainticket/internal/model"
	"github.com/AliRamdhan/trainticket/internal/service"
	"github.com/gin-gonic/gin"
)

type OrderTicketHandler struct {
	orderTicket *service.OrderTicketService
}

func NewOrderTicketHandler(cs *service.OrderTicketService) *OrderTicketHandler {
	return &OrderTicketHandler{orderTicket: cs}
}

// func (oh *OrderTicketHandler) CreateOrder(c *gin.Context) {
// 	var order model.Order
// 	if err := c.ShouldBindJSON(&order); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	// Extract serviceCategoryID and customerId from the request body
// 	seatId := order.OrderSeatId
// 	ticketTrainId := order.OrderTicketId
// 	userID := order.OrderUserId

// 	// Check if serviceCategoryID is missing or invalid
// 	if seatId == 0 {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing service category ID"})
// 		return
// 	}

// 	// Check if customerId is missing or invalid
// 	if ticketTrainId == 0 {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing customer ID"})
// 		return
// 	}

// 	if userID == 0 {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing customer ID"})
// 		return
// 	}

// 	if err := oh.orderTicket.CreateOrder(&order, seatId, userID, ticketTrainId); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

//		c.JSON(http.StatusCreated, gin.H{"message": "Order created successfully", "order": order})
//	}
func (oh *OrderTicketHandler) CreateOrder(c *gin.Context) {
	var order model.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	seatId := order.OrderSeatId
	ticketTrainId := order.OrderTicketId
	userID := order.OrderUserId

	if seatId == 0 || ticketTrainId == 0 || userID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing parameters"})
		return
	}

	if err := oh.orderTicket.CreateOrder(&order, seatId, userID, ticketTrainId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Order created successfully", "order": order})
}

func (oh *OrderTicketHandler) GetAllOrderByUser(c *gin.Context) {
	userIdStr := c.Param("userId")

	var userId uint
	_, err := fmt.Sscanf(userIdStr, "%d", &userId)
	// trackID, err := uuid.Parse(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order number format"})
		return
	}

	orders, err := oh.orderTicket.GetAllOrderByUser(userId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "List yout Order", "Orders": orders})
}

func (oh *OrderTicketHandler) GetAllOrder(c *gin.Context) {
	tickets, err := oh.orderTicket.GetAllOrder()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "All Orders Log", "Orders": tickets})
}
