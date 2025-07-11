package order

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Order struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

type OrderService struct{}

func NewOrderService() *OrderService {
	return &OrderService{}
}

func (s *OrderService) CreateOrder(c *gin.Context) {
	var order Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order data"})
		return
	}

	// Gửi request đến Auth service để xác thực token (sử dụng HTTP client)
	token := c.GetHeader("Authorization")
	authResponse, err := verifyToken(token)
	if err != nil || authResponse["status"] != "ok" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return
	}

	// Xử lý order
	c.JSON(http.StatusOK, gin.H{"message": "Order created", "order": order})
}

func verifyToken(token string) (map[string]interface{}, error) {
	client := &http.Client{}
	// req, err := http.NewRequest("POST", "http://localhost:8081/auth/verify", nil)
	req, err := http.NewRequest("POST", "http://auth:8081/auth/verify", nil) // Nếu chạy trong docker-compose, dùng `http://auth:8081`
	req.Header.Set("Authorization", token)                                   // token đã có prefix "Bearer "
	if err != nil {
		return nil, err
	}

	// req.Header.Add("Authorization", token)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var authResponse map[string]interface{}
	json.Unmarshal(body, &authResponse)
	return authResponse, nil
}
