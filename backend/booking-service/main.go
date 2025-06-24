package main

import (
  "booking-service/config"
  "booking-service/routes"
  "github.com/gin-contrib/cors"
  "github.com/gin-gonic/gin"
  "time"
)

func main() {
  config.Connect()

  r := gin.Default()


  r.Use(cors.New(cors.Config{
    AllowOrigins:     []string{"http://localhost:5173"}, 
    AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
    AllowCredentials: true,
    MaxAge:           12 * time.Hour,
  }))

  routes.SetupRoutes(r)

  r.Run(":8081")
}
