package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"math/big"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()
	r.Use(Cors())

	r.GET("/api/permit-typed-data", getPermitTypedData)
	r.POST("/api/submit-permit", submitPermit)

	r.Run(":8080")
}

func getPermitTypedData(c *gin.Context) {
	// 这里模拟从 URL 参数中获取 `owner`，实际中可以做更复杂的身份验证和处理
	owner := c.DefaultQuery("owner", "0xYourWalletAddress")

	// 数据结构构造
	deadline := big.NewInt(time.Now().Add(time.Hour).Unix())
	nonce := big.NewInt(1) // 假设 nonce 为 1，实际应通过调用合约来获取)

	message := map[string]interface{}{
		"owner":    owner,
		"spender":  "0x82F35a173204BfBa3899bfd9cCDfbF17BEaeC8B2",
		"nonce":    nonce.String(),
		"value":    "1000000",
		"deadline": deadline.String(),
	}

	domain := map[string]interface{}{
		"name":              "USD Coin",
		"version":           "2",
		"chainId":           1,
		"verifyingContract": "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
	}

	// EIP712 Domain 信息
	typedData := map[string]interface{}{
		"types": map[string]interface{}{
			"EIP712Domain": []map[string]string{
				{"name": "name", "type": "string"},
				{"name": "version", "type": "string"},
				{"name": "chainId", "type": "uint256"},
				{"name": "verifyingContract", "type": "address"},
			},
			"Permit": []map[string]string{
				{"name": "owner", "type": "address"},
				{"name": "spender", "type": "address"},
				{"name": "value", "type": "uint256"},
				{"name": "nonce", "type": "uint256"},
				{"name": "deadline", "type": "uint256"},
			},
		},
		"primaryType": "Permit",
		"domain":      domain,
		"message":     message,
	}
	c.JSON(http.StatusOK, typedData)
}

func submitPermit(c *gin.Context) {
	var request struct {
		Signature string                 `json:"signature"`
		TypedData map[string]interface{} `json:"typedData"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 模拟处理签名结果，可以在这里加入调用合约的逻辑
	log.Println("Received signature:", request.Signature)

	// 返回一个成功消息
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
