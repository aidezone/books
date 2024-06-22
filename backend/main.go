package main

import (
    "log"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "books-manage-system/config"
    "books-manage-system/routes"
    "books-manage-system/services"
    _ "books-manage-system/docs"
    swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Library Management API
// @version 1.0
// @description This is a sample server for a library management system.

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
    if err := godotenv.Load(); err != nil {
        log.Fatalf("Error loading .env file")
    }

    r := gin.Default()

    config.ConnectDatabase()
    routes.RegisterRoutes(r)

    services.InitializeAdmin()

    // Swagger
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    s := &http.Server{
        Addr:           ":8080",
        Handler:        r,
        ReadTimeout:    10 * time.Second,
        WriteTimeout:   10 * time.Second,
        MaxHeaderBytes: 1 << 20,
    }

    s.ListenAndServe()
}
