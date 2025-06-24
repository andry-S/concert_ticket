package main

import (
    "payment-service/config"
    "payment-service/routes"
    "github.com/gin-gonic/gin"
)

func main() {
    config.Connect()
    r := gin.Default()
    routes.SetupRoutes(r)
    r.Run(":8082")
}
