package main

import (
    "github.com/gin-gonic/gin"
    "github.com/luciocarvalhojr/go-api/internal/handler"
    _ "github.com/luciocarvalhojr/go-api/docs"
    swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Go Album API
// @version         1.0
// @description     A sample album management API in Go.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /

func main() {
    router := gin.Default()
    
    // Register album routes
    handler.RegisterRoutes(router)

    // Swagger route
    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    // Running on :8080 by default. 
    // In production, you would probably want to use an environment variable for the port.
    err := router.Run(":8080")
    if err != nil {
        panic("Failed to start the server: " + err.Error())
    }
}
