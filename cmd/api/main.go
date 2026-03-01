package main

import (
    "github.com/gin-gonic/gin"
    "github.com/luciocarvalhojr/go-api/internal/handler"
)

func main() {
    router := gin.Default()
    
    // Register album routes
    handler.RegisterRoutes(router)

    // Running on :8080 by default. 
    // In production, you would probably want to use an environment variable for the port.
    err := router.Run(":8080")
    if err != nil {
        panic("Failed to start the server: " + err.Error())
    }
}
