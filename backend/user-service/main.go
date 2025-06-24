package main

import (
    "user-service/database"
    "user-service/routes"

    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "time"
)

func main() {
    database.ConnectDB()

    r := gin.Default()


    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:5173"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
        AllowCredentials: true,
        MaxAge: 12 * time.Hour,
    }))

    routes.UserRoutes(r)

    r.Run(":8001")
}
