package webhooks

import "github.com/gin-gonic/gin"

func main() {
    engine := gin.Default()
    engine.GET("/", func(context *gin.Context) {
        context.JSON(200, "pong")
    })
    engine.Run("0.0.0.0:8080")
}
