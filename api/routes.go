package api

import (
	"github.com/AliRamdhan/trainticket/internal/handler"
	"github.com/AliRamdhan/trainticket/internal/handler/auth_handler"
	"github.com/AliRamdhan/trainticket/internal/middlewares"
	"github.com/AliRamdhan/trainticket/internal/service"
	"github.com/AliRamdhan/trainticket/internal/service/auth_service"
	"github.com/gin-gonic/gin"
)

func ServiceAuth(r *gin.Engine, authUser *auth_service.AuthService) {
	authHandler := auth_handler.NewAuthHandler(authUser)
	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/register", authHandler.RegisterAuth)
		authRoutes.POST("/login", authHandler.Login)
		//authRoutes.POST("/home", authHandler.Home).Use(middlewares.Auth())
		securedUserRoutes := r.Group("/home").Use(middlewares.UserAuth())
		{
			securedUserRoutes.GET("/user", authHandler.Home)
		}
		securedAdminRoutes := r.Group("/home").Use(middlewares.AdminAuth())
		{
			securedAdminRoutes.GET("/admin", authHandler.Home)
		}
	}
}

func UserService(r *gin.Engine, userService *service.UserService) {
	handler := handler.NewUserHandler(userService)
	routes := r.Group("/user")
	{
		routes.GET("/all", handler.GetAllUser)
		routes.GET("/:userId", handler.GetUserById)
		routes.PUT("/update/:userId", handler.UpdateUser)
		routes.DELETE("delete/:userId", handler.DeleteUser)
	}
}

func ServiceRole(r *gin.Engine, roleService *service.RoleService) {
	roleHandler := handler.NewRoleHandler(roleService)
	roleRoutes := r.Group("/role")
	{
		roleRoutes.POST("/create", roleHandler.CreateRole)
		roleRoutes.GET("/all", roleHandler.GetAllRole)
		roleRoutes.PUT("/update/:roleId", roleHandler.UpdateRole)
		roleRoutes.DELETE("delete/:roleId", roleHandler.DeleteRole)
	}
}

func ServiceTicket(r *gin.Engine, ticketService *service.TicketService) {
	ticketHandler := handler.NewTicketHandler(ticketService)
	ticketRoutes := r.Group("/tickets")
	{
		ticketRoutes.POST("/create", ticketHandler.CreateTicket)
		ticketRoutes.GET("/:ticketId", ticketHandler.GetTicketById)
		ticketRoutes.GET("/details/:ticketName", ticketHandler.GetTicketByName)
		ticketRoutes.GET("/all", ticketHandler.GetAllTicket)
		ticketRoutes.PUT("/update/:ticketId", ticketHandler.UpdateTicket)
		ticketRoutes.DELETE("delete/:ticketId", ticketHandler.DeleteTicket)
	}
}

func ServiceSeat(r *gin.Engine, seatService *service.SeatService) {
	seatHandler := handler.NewSeatTrainHandler(seatService)
	seatRoutes := r.Group("/seats")
	{
		seatRoutes.GET("/all", seatHandler.GetAllSeats)
		seatRoutes.GET("/:seatName", seatHandler.GetSeatBySeatName)
		seatRoutes.POST("/create", seatHandler.CreateSeat)
		seatRoutes.GET("/all/:ticketId", seatHandler.GetAllSeatsFromTicketId)
		//GetAllSeats for admin
		seatRoutes.PUT("/update/:seatId", seatHandler.UpdateSeat)
		seatRoutes.DELETE("delete/:seatId", seatHandler.DeleteSeat)
	}
}

func ServiceOrder(r *gin.Engine, orderService *service.OrderTicketService) {
	orderHandler := handler.NewOrderTicketHandler(orderService)
	orderRoutes := r.Group("/orders")
	{
		orderRoutes.POST("/create", orderHandler.CreateOrder)
		orderRoutes.GET("/all", orderHandler.GetAllOrder)
		orderRoutes.GET("/:orderNumb", orderHandler.GetOrderByOrderNumber)
		orderRoutes.GET("/user/:userId", orderHandler.GetAllOrderByUser)
	}
}
func ServicePayment(r *gin.Engine, paymentService *service.PaymentService) {
	paymentHandler := handler.NewPaymentHandler(paymentService)
	paymentRoutes := r.Group("/payment")
	{
		paymentRoutes.POST("/create", paymentHandler.CreatePayment)
		paymentRoutes.GET("/all", paymentHandler.GetAllPayments)
		// paymentRoutes.GET("/user/:userId", paymentHandler.GetAllOrderByUser)
	}
}
