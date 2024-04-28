package handler

import (
	"fmt"
	"net/http"

	"github.com/AliRamdhan/trainticket/internal/model"
	"github.com/AliRamdhan/trainticket/internal/service"
	"github.com/gin-gonic/gin"
)

type TicketHandler struct {
	ticketService *service.TicketService
}

func NewTicketHandler(ps *service.TicketService) *TicketHandler {
	return &TicketHandler{ticketService: ps}
}

func (th *TicketHandler) CreateTicket(c *gin.Context) {
	var ticket model.Ticket
	if err := c.ShouldBindJSON(&ticket); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := th.ticketService.CreateTicket(&ticket); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Ticket created successfully", "product": ticket})
}

func (th *TicketHandler) GetAllTicket(c *gin.Context) {
	tickets, err := th.ticketService.GetAllTicket()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "All tickets", "Tickets": tickets})
}

func (th *TicketHandler) GetTicketById(c *gin.Context) {
	ticketIdStr := c.Param("ticketId")

	var ticketId uint
	_, err := fmt.Sscanf(ticketIdStr, "%d", &ticketId)
	// trackID, err := uuid.Parse(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ticket id format"})
		return
	}
	tickets, err := th.ticketService.GetTicketById(ticketId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Details tickets", "Tickets": tickets})
}

func (th *TicketHandler) UpdateTicket(c *gin.Context) {
	var ticket model.Ticket
	if err := c.ShouldBindJSON(&ticket); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ticketIdStr := c.Param("ticketId")
	// ticketID, err := uuid.Parse(ticketIdStr)
	var ticketId uint
	_, err := fmt.Sscanf(ticketIdStr, "%d", &ticketId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ticket ID"})
		return
	}

	if err := th.ticketService.UpdateTicket(ticketId, &ticket); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ticket updated successfully", "ticket": ticket})
}

func (th *TicketHandler) DeleteTicket(c *gin.Context) {
	ticketIdStr := c.Param("ticketId")
	// productID, err := uuid.Parse(ticketIdStr)
	var ticketId uint
	_, err := fmt.Sscanf(ticketIdStr, "%d", &ticketId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	if err := th.ticketService.DeleteTicket(ticketId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Ticket deleted successfully"})
}
