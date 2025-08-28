package main
import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)
type User struct {
    ID   string `json:"id"`
    Name string `json:"name"`
}
type Order struct {
    OrderID string `json:"order_id"`
    UserID  string `json:"user_id"`
    Item    string `json:"item"`
}
func main() {
    router := gin.Default()
    router.POST("/orders", func(c *gin.Context) {
        var order Order
        if err := c.ShouldBindJSON(&order); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        // Получение пользователя из user-service
        resp, err := http.Get(fmt.Sprintf("http://localhost:9091/users/%s", order.UserID))
        if err != nil || resp.StatusCode != 200 {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "не удалось получить пользователя"})
            return
        }
        var user User
        json.NewDecoder(resp.Body).Decode(&user)
        // Псевдо-создание заказа
        fmt.Printf("Заказ создан для %s (%s): %s\n", user.Name, user.ID, order.Item)
        // Отправка уведомления
        http.Post("http://localhost:9093/notify", "application/json", 
            jsonBody(map[string]string{
                "message": fmt.Sprintf("Пользователь %s сделал заказ на %s", user.Name, order.Item),
            }),
        )

        c.JSON(http.StatusOK, gin.H{"status": "ok"})
    })
    router.Run(":9092")
}
func jsonBody(data map[string]string) *bytes.Buffer {
    b := new(bytes.Buffer)
    json.NewEncoder(b).Encode(data)
    return b
}

