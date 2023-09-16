package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    // Create a new Gin router
    r := gin.Default()

    // Define a route
    r.GET("/hello", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "Hello, Gin!",
        })
    })

    // Run the server on port 8080
    r.Run(":9090")
}
