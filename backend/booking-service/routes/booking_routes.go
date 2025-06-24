package routes

import (
  "booking-service/controller"
  "github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
  r.POST("/bookings", controller.CreateBooking)
  r.GET("/bookings", controller.GetAllBookings)
  r.GET("/bookings/:id", controller.GetBookingByID)
  r.GET("/bookings/user/:id", controller.GetBookingsByUserID)
  r.DELETE("/bookings/:id", controller.DeleteBooking)
  r.POST("/bookings/:id/pay", controller.PayBooking)

}
