package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AliRamdhan/trainticket/api"
	"github.com/AliRamdhan/trainticket/config"
	"github.com/AliRamdhan/trainticket/internal/service"
	"github.com/AliRamdhan/trainticket/internal/service/auth_service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	if err := config.ConnectDB(); err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Connected to database")
	// Perform migrations
	// if err := config.AutoMigrate(); err != nil {
	// 	// Handle error
	// 	log.Fatalf("Error applying migration: %v", err)
	// }
	// log.Println("Migration Applied Successfully")
	// users, roles := config.SeedData()
	// log.Printf("Seeded %d users and %d roles into the database", len(users), len(roles))
	r := gin.Default()

	err := godotenv.Load(".env.dev")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	// Initialize the product service
	authenticationService := auth_service.NewAuthService()
	roleService := service.NewRoleService()
	ticketService := service.NewTicketService()
	seatService := service.NewSeatService()
	orderService := service.NewOrderTicketService()
	paymentService := service.NewPaymentService()
	userService := service.NewUserService()

	// Setup middleware
	r.Use(enableCORS())
	r.Use(jsonContentTypeMiddleware())

	// Setup routes
	api.ServiceAuth(r, authenticationService)
	api.ServiceRole(r, roleService)
	api.ServiceTicket(r, ticketService)
	api.ServiceSeat(r, seatService)
	api.ServiceOrder(r, orderService)
	api.ServicePayment(r, paymentService)
	api.UserService(r, userService)

	r.GET("/", func(c *gin.Context) {
		// message := fmt.Sprintf("Hello World %s")
		c.String(http.StatusOK, "Hello World")
	})

	// port := fmt.Sprintf(PORT)
	r.Run()
}
func enableCORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Set CORS headers
		origin := c.Request.Header.Get("Origin")
		c.Writer.Header().Set("Access-Control-Allow-Origin", origin) // Allow specific origin
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization") // Allow Content-Type and Authorization headers
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		// Check if the request is for CORS preflight
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		// Pass down the request to the next middleware (or final handler)
		c.Next()
	}
}

func jsonContentTypeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Set JSON Content-Type
		c.Header("Content-Type", "application/json")
		c.Next()
	}
}
