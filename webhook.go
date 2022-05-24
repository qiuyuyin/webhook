package main

import (
    "github.com/gin-gonic/gin"
    "webhooks/github"
)

func main() {
    engine := gin.Default()
    engine.GET("/ping", func(context *gin.Context) {
        context.JSON(200, "pong")
    })
    engine.POST("/github", github.GithubHandler)
    engine.Run("0.0.0.0:8080")
}
