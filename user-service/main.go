package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

type User struct {
    ID   string `json:"id"`
    Name string `json:"name"`
}

func main() {
    router := gin.Default()

    router.GET("/users/:id", func(c *gin.Context) {
        id := c.Param("id")
        user := User{ID: id, Name: "Иван Петров"}
        c.JSON(http.StatusOK, user)
    })

    router.Run(":9091")
}
