package main

import (
    "fmt"
    "net/http"

    "github.com/gin-gonic/gin"
)

type Notification struct {
    Message string `json:"message"`
}

func main() {
    router := gin.Default()

    router.POST("/notify", func(c *gin.Context) {
        var notif Notification
        if err := c.ShouldBindJSON(&notif); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        // Имитируем отправку уведомления
        fmt.Println("🔔 Уведомление:", notif.Message)

        c.JSON(http.StatusOK, gin.H{"status": "sent"})
    })

    router.Run(":9093")
}
