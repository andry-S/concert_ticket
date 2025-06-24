package routes

import (
    "payment-service/controller"
    "github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
    r.POST("/payments", controller.CreatePayment)
}
